package game

import (
	"encoding/json"
	"io/ioutil"
)

var gameJsonMap map[string]map[string]interface{}

func init() {
	gameJsonMap = make(map[string]map[string]interface{})
}
func GetConf(confName string) map[string]interface{} {
	jsonData, ok := gameJsonMap[confName]
	if ok {
		return jsonData
	}
	jsonByte, err := ioutil.ReadFile("conf/game/" + confName + ".json")
	if err != nil {
		return nil
	}
	jsonConfData := make(map[string]interface{})
	json.Unmarshal(jsonByte, &jsonConfData)
	gameJsonMap[confName] = jsonConfData
	return jsonConfData
}
