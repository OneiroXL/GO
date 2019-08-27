package base

type InteractiveResponse struct {
	Status int
	Type int
	Message string
	Data string
}

type InteractiveRequest struct{
	Type int
	Data string
}