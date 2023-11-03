package application

const (
	passwordData = iota + 1
	cardData
	fileData
	textData
	passwordDataType = "password"
	cardDataType     = "card"
	fileDataType     = "file"
	textDataType     = "text"
	loginInput       = "login: "
	passwordInput    = "password: "
	siteInput        = "site: "
	cardNumberInput  = "card number: "
	cardHolderInput  = "card holder: "
	cardExpDateInput = "expires date (mm/yy): "
)

const (
	getUserDataList = iota + 1
	getUserData
	saveUserData
	editUserData
)
