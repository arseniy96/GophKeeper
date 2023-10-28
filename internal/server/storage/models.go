package storage

type User struct {
	Login    string
	Password string
	Token    string
	ID       int64
}

type ShortRecord struct {
	Name      string
	DataType  string
	CreatedAt string
	ID        int64
	Version   int64
}

type Record struct {
	Name      string
	DataType  string
	CreatedAt string
	Data      []byte
	Version   int64
	ID        int64
}
