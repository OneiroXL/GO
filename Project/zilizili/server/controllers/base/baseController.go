package baseControllers
 
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strings"
	"log"
	"zilizili/model/base/result"
)
 
type BaseController struct {
	beego.Controller
}

/*
结果
*/
func (this *BaseController) Response(resultType string,obj interface{}) {
	resultTypeKey := strings.ToLower(resultType)
	this.Data[resultTypeKey] = &obj
	switch(resultType){
		case "JSON":{
			this.ServeJSON()
			return
		}
		case "XML":{
			this.ServeXML()
			return
		}
		case "JSONP":{
			this.ServeJSONP()
			return
		}
	}
}

/*
返回JSON数据结构
*/
func (this *BaseController) ResponseJSON(obj interface{}) {
	this.Response("JSON",obj)
}

/*
返回JSONP数据结构
*/
func (this *BaseController) ResponseJSONP(obj interface{}) {
	this.Response("JSONP",obj)
}

/*
返回XML数据结构
*/
func (this *BaseController) ResponseXML(obj interface{}) {
	this.Response("XML",obj)
}

/*
绑定数据
*/
func (this *BaseController) BindParseForm(obj interface{} ) (bool) {
	if err := this.ParseForm(obj); err != nil {
		log.Println(err)
		return false
	}
	return true;
}

/*
验证数据
*/
func (this *BaseController) ValidParam(obj interface{} ) (bool,[]result.VerificationMessage) {
	valid := validation.Validation{}
	valid.Valid(obj)
	if valid.HasErrors() {
		var verificationMessageList []result.VerificationMessage
        for _, err := range valid.Errors {
			verificationMessage := result.VerificationMessage{}
			verificationMessage.Field = err.Field
			verificationMessage.Value = err.Value
			verificationMessage.Message = err.Message

			verificationMessageList = append(verificationMessageList,verificationMessage)
		}
		return !valid.HasErrors(),verificationMessageList
	}
	return !valid.HasErrors(),nil
}

