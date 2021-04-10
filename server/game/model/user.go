package model

type UserModel struct {
	Id            int32
	OpenId        string
	UserSpid      string
	UserInitTime  int32
	UserMoney     float32
	UserId        int64
	UnionId       string
	UserPlatform  int32
	UserYouKe     int8
	IsCertificate int8
	DeviceId      string
	DeviceType    string
}
