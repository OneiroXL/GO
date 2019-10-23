package streamserver

import (
	"io"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter,sc int,mesage string)  {
	w.WriteHeader(sc)
	io.WriteString(w,mesage)
}