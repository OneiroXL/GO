package account

type UserModel struct {
	Id      int64  `json:"Id"`
	Type    string `json:"Type"`
	Subtype string `json:"Subtype"`
	State   string `json:"State"`
}
