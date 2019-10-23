package streamserver

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)




func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video/:vid-id",StreamHandler)
	router.POST("/upload/:vid-id",UploadHandle)
	return router
}


func Start()  {
	router := RegisterHandlers();
	mh := NewMiddleWareHandle(router,2)
	http.ListenAndServe(":9000",mh)
}