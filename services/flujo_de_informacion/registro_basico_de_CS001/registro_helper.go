package registro

type Response_WithInt struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      int    `json:"data"`
}

type Response_WithString struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_WithPhoneAndCountry struct {
	Error     bool            `json:"error"`
	DataError string          `json:"dataError"`
	Data      PhoneAndCountry `json:"data"`
}

type PhoneAndCountry struct {
	Phone   int `json:"phone"`
	Country int `json:"country"`
}
