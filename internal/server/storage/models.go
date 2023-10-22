package storage

type User struct {
	ID       int64
	Login    string
	Password string
	Token    string
}

type ShortRecord struct {
	ID        int64
	Name      string
	DataType  string
	Version   int64
	CreatedAt string
}

type Record struct {
	ID        int64
	Name      string
	Data      []byte
	DataType  string
	Version   int64
	CreatedAt string
}
