package server

func BroadCastByUidS(uids []int64, data map[string]interface{}, t int) {
	if len(uids) == 0 {
		return
	}
	for _, uid := range uids {
		user := GetUser(uid)
		if user == nil {
			continue
		}
		data["t"] = t
		user.Agent.WriteMsg(data)
	}
}
