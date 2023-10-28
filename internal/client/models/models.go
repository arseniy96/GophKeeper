package models

type AuthModel struct {
	Login    string
	Password string
}

type AuthToken string

type UserDataList struct {
	Name     string
	DataType string
	ID       int64
	Version  int64
}

type UserDataModel struct {
	ID int64
}

type UserData struct {
	Name     string
	DataType string
	Data     []byte
	ID       int64
	Version  int64
}
