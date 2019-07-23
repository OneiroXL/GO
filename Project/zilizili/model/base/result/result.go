package result

import (
	"zilizili/common/enum/resultEnum"
)

type Result struct {
	Status resultEnum.Status
	Message interface{}
}

type ResultData struct{
	Result
	Data interface{}
}

type ResultPageData struct{
	ResultData
	Count int
}

func NewResult(status resultEnum.Status) *Result {
	var result *Result
	switch status {
		case resultEnum.Success:{
			result = &Result{
				Status : resultEnum.Success,
				Message:"成功",
			}
			return result
		}
		case resultEnum.Failed:{
			result = &Result{
				Status : resultEnum.Failed,
				Message:"失败",
			}
			return result
		}
		case resultEnum.Error:{
			result = &Result{
				Status : resultEnum.Error,
				Message:"错误",
			}
			return result
		}
		default:{
			return result
		}
	}
}

func NewResultData(status resultEnum.Status) *ResultData {
	resultData := &ResultData{}
	switch status {
		case resultEnum.Success:{
			resultData.Status = resultEnum.Success
			resultData.Message = "成功"

			return resultData
		}
		case resultEnum.Failed:{
			resultData.Status = resultEnum.Failed
			resultData.Message = "失败"

			return resultData
		}
		case resultEnum.Error:{
			resultData.Status = resultEnum.Error
			resultData.Message = "错误"
			return resultData
		}
		default:{
			return resultData
		}
	}
}

func NewResultPageData(status resultEnum.Status) *ResultPageData {
	resultPageData := &ResultPageData{}
	switch status {
		case resultEnum.Success:{
			resultPageData.Status = resultEnum.Success
			resultPageData.Message = "成功"
			
			return resultPageData
		}
		case resultEnum.Failed:{
			resultPageData.Status = resultEnum.Failed
			resultPageData.Message = "失败"

			return resultPageData
		}
		case resultEnum.Error:{
			resultPageData.Status = resultEnum.Error
			resultPageData.Message = "错误"
			return resultPageData
		}
		default:{
			return resultPageData
		}
	}
}

func NewSuccessResult() *Result {
	return NewResult(resultEnum.Success)
}

func NewFailedResult() *Result {
	return NewResult(resultEnum.Failed)
}

func NewErrorResult() *Result {
	return NewResult(resultEnum.Error)
}


func NewSuccessResultData() *ResultData {
	return NewResultData(resultEnum.Success)
}

func NewFailedResultData() *ResultData {
	return NewResultData(resultEnum.Failed)
}

func NewErrorResultData() *ResultData {
	return NewResultData(resultEnum.Error)
}

func NewSuccessResultDataPage() *ResultPageData {
	return NewResultPageData(resultEnum.Success)
}

func NewFailedResultDataPage() *ResultPageData {
	return NewResultPageData(resultEnum.Failed)
}

func NewErrorResultDataPage() *ResultPageData {
	return NewResultPageData(resultEnum.Error)
}