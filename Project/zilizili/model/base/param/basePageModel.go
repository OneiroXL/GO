package baseParam

type BaseParamModel struct {
	PageIndex int `form:"PageIndex" json:"PageIndex" valid:"Required"`
	PageSize int `form:"PageSize" json:"PageSize" valid:"Required"`
	Search string
}

