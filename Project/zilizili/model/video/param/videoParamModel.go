package param

import (
	
)

type VideoParamModel struct{
	ID uint `form:"ID"`
	Title string `form:"Title" json:"Title" valid:"Required"`
	Info  string `form:"Info" json:"Info" valid:"Required"`
	URL   string `form:"URL" json:"URL" valid:"Required"`
}
