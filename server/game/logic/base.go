package logic

type Base struct {
}

func (this *Base) Handle(args []interface{}) {

}
func (this *Base) ErrorMsg(Route string, ReqId int, params map[string]interface{}, msg string, code int) map[string]interface{} {
	res := make(map[string]interface{})
	res["msg"] = msg
	res["s"] = code
	res["t"] = Route
	res["mp"] = params
	res["reqId"] = ReqId
	return res
}
func (this *Base) Success(Route string, ReqId int, params map[string]interface{}) map[string]interface{} {
	params["s"] = 0
	params["t"] = Route
	params["reqId"] = ReqId
	return params
}
