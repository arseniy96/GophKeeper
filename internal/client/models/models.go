package models

type AuthModel struct {
	Login    string
	Password string
}

type AuthToken string

type UserDataList struct {
	ID       int64
	Name     string
	DataType string
	Version  int64
}

type UserDataModel struct {
	ID int64
}

type UserData struct {
	ID       int64
	Name     string
	DataType string
	Version  int64
	Data     []byte
}
