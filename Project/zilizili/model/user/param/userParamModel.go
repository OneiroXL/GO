package param

type UserParamModel struct{
	ID			uint `from:"ID" json:"Title"`
	UserName	string `from:"UserName" json:"UserName"  valid:"Required"`
	Password	string `from:"Password" json:"Password"  valid:"Required"`
	Nickname	string `from:"Nickname" json:"Nickname"  valid:"Required"`
	Status		string `from:"Status" json:"Status"`
	Avatar		string `from:"Avatar" json:"Avatar"`
}