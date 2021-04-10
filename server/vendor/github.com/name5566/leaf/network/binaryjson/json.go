package binaryjson

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	chanrpc2 "github.com/name5566/leaf/chanrpc"
	log2 "github.com/name5566/leaf/log"
	"reflect"
	"time"
)

type Processor struct {
	msgInfo map[string]*MsgInfo
	routeMap map[string]string
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc2.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}

type MsgHandler func([]interface{})

type MsgRaw struct {
	msgID      string
	msgRawData json.RawMessage
}

func NewProcessor() *Processor {
	p := new(Processor)
	p.msgInfo = make(map[string]*MsgInfo)
	p.routeMap = make(map[string]string)
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(msg interface{},routeKey string) string {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log2.Fatal("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	if msgID == "" {
		log2.Fatal("unnamed json message")
	}
	if _, ok := p.msgInfo[msgID]; ok {
		log2.Fatal("message %v is already registered", msgID)
	}
	if _, ok := p.routeMap[routeKey]; ok {
		log2.Fatal("message %v is already registered route", routeKey)
	}
	i := new(MsgInfo)
	i.msgType = msgType
	p.msgInfo[msgID] = i
	p.routeMap[routeKey] = msgID
	return msgID
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRouter(msg interface{}, msgRouter *chanrpc2.Server) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log2.Fatal("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	i, ok := p.msgInfo[msgID]
	if !ok {
		log2.Fatal("message %v not registered", msgID)
	}

	i.msgRouter = msgRouter
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetHandler(msg interface{}, msgHandler MsgHandler) {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log2.Fatal("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	i, ok := p.msgInfo[msgID]
	if !ok {
		log2.Fatal("message %v not registered", msgID)
	}

	i.msgHandler = msgHandler
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRawHandler(msgID string, msgRawHandler MsgHandler) {
	i, ok := p.msgInfo[msgID]
	if !ok {
		log2.Fatal("message %v not registered", msgID)
	}

	i.msgRawHandler = msgRawHandler
}

// goroutine safe
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	// raw
	if msgRaw, ok := msg.(MsgRaw); ok {
		i, ok := p.msgInfo[msgRaw.msgID]
		if !ok {
			return fmt.Errorf("message %v not registered", msgRaw.msgID)
		}
		if i.msgRawHandler != nil {
			i.msgRawHandler([]interface{}{msgRaw.msgID, msgRaw.msgRawData, userData})
		}
		return nil
	}

	// json
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return errors.New("json message pointer required")
	}
	msgID := msgType.Elem().Name()
	i, ok := p.msgInfo[msgID]
	if !ok {
		return fmt.Errorf("message %v not registered", msgID)
	}
	if i.msgHandler != nil {
		i.msgHandler([]interface{}{msg, userData})
	}
	if i.msgRouter != nil {
		i.msgRouter.Go(msgType, msg, userData)
	}
	return nil
}

// goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	log2.Debug("v%",string(data))
	data=bytes.Trim(data, "\x00")
	data=data[4:len(data)]
	var m map[string]interface{}
	err := json.Unmarshal(data, &m)
	//log.Debug("v%",m)
	if err != nil {
		return nil, err
	}
	route:=m["route"].(string)
	msgID, ok := p.routeMap[route]
	if !ok {
		return nil, fmt.Errorf("message %v not registered route", route)
	}
	i, ok := p.msgInfo[msgID]
	if !ok {
		return nil, fmt.Errorf("message %v not registered", msgID)
	}

	// msg
	if i.msgRawHandler != nil {
		return MsgRaw{msgID, data}, nil
	} else {
		msg := reflect.New(i.msgType.Elem()).Interface()
		return msg, json.Unmarshal(data, msg)
	}
	/*for msgID, data := range m {
		i, ok := p.msgInfo[msgID]
		if !ok {
			return nil, fmt.Errorf("message %v not registered", msgID)
		}

		// msg
		if i.msgRawHandler != nil {
			return MsgRaw{msgID, data}, nil
		} else {
			msg := reflect.New(i.msgType.Elem()).Interface()
			return msg, json.Unmarshal(data, msg)
		}
	}*/

	panic("bug")
}

// goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	msgMap:=msg.(map[string]interface{})
	// data
	m := map[string]interface{}{"r": "mp","type":4}
	msgBody:=map[string]interface{}{"t": msgMap["t"],"sc":float64(time.Now().UnixNano())/float64(1e9),"d":msgMap}
	m["d"]=msgBody
	data, err := json.Marshal(m)
	jsonLen := len(string(data))
	jsonLenByte:= intToBytes(int32(jsonLen))
	allData:= BytesCombine(jsonLenByte,data);
	return [][]byte{allData}, err
}

func intToBytes(n int32) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, n)
	return buf.Bytes()
}
func BytesCombine(pBytes ...[]byte) []byte {
	length := len(pBytes)
	s := make([][]byte, length)
	for index := 0; index < length; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

