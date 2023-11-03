package application

type PasswordData struct {
	Site     string `json:"site"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type CardData struct {
	Number     string `json:"number"`
	ExpDate    string `json:"exp_date"`
	CardHolder string `json:"card_holder"`
}

type FileData struct {
	Path string `json:"path"`
	Data []byte `json:"data"`
}

type TextData struct {
	Text string `json:"text"`
}
