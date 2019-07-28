package param

type UserLoginModel struct{
	LoginType 		int `from:"LoginType" json:"LoginType"  valid:"Required"`
	LoginParam		string `from:"LoginParam" json:"LoginParam"  valid:"Required"`
	Password		string `from:"Password" json:"Password"  valid:"Required"`
	MessageCode		string `from:"MessageCode" json:"MessageCode"`
}