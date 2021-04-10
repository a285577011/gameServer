package model

import (
	"database/sql"
	"github.com/name5566/leaf/log"
	"server/db/mysql"
	"server/game/common"
	"strings"
)

type Orm struct {
	db *sql.DB
}

var (
	CenterOrm *Orm
	ServerOrm *Orm
)

func init() {
	NewCenterOrm()
	NewServerOrm()
}
func NewCenterOrm() *Orm {
	if CenterOrm == nil {
		CenterOrm = &Orm{mysql.MysqlCenterDb}
	}
	return CenterOrm
}
func NewServerOrm() *Orm {
	if ServerOrm == nil {
		ServerOrm = &Orm{mysql.MysqlServerDb}
	}
	return ServerOrm
}

//根据SQL查询
func (this *Orm) query(sql string, isOne bool) ([]map[string]string, error) {
	if isOne && !strings.Contains(sql, "limit") {
		sql = sql + " limit 1"
	}
	rows, err := this.db.Query(sql)
	if err != nil {
		log.Error("sqlErr", err)
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Error("sqlErr", err)
	}

	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	data := make([]map[string]string, 0) //创建一个新的list

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		obj := this.parseQuery(columns, values)
		data = append(data, obj)
	}
	return data, err
}

//对一条查询结果进行封装
func (this *Orm) parseQuery(columns []string, values []interface{}) map[string]string {
	data := make(map[string]string)
	for index, val := range columns {
		str := ""

		if values[index] != nil {
			str = common.Bytes2String(values[index].([]byte))
		}
		data[val] = str
	}

	return data
}

//根据SQL查询多条记录
func (this *Orm) GetAll(sql string) ([]map[string]string, error) {
	return this.query(sql, false)
}

//根据SQL查询一条记录,如果找到不数据，data会返回nil
func (this *Orm) FindOne(sql string) (map[string]string, error) {
	datas, err := this.query(sql, true)

	var data map[string]string
	if len(datas) > 0 {
		data = datas[0]
	}
	return data, err

}
func (this *Orm) Insert(sql string) int64 {
	datas, err := this.db.Exec(sql)
	if err != nil {
		lastInsertID, _ := datas.LastInsertId()
		return lastInsertID
	}
	return 0
}
func (this *Orm) Update(sql string) int64 {
	datas, err := this.db.Exec(sql)
	if err != nil {
		//影响行数
		rowsEffect, _ := datas.RowsAffected()
		return rowsEffect
	}
	return 0
}
