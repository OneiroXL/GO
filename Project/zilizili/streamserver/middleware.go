package streamserver


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


type MiddleWareHandle struct  {
	r *httprouter.Router
	l *Limiter
}

/*
构造函数
*/
func NewMiddleWareHandle(r *httprouter.Router,n int) *MiddleWareHandle {
	middleWareHandle := &MiddleWareHandle{
		r:r,
		l:NewLimiter(n),
	}
	return middleWareHandle
}

func (this *MiddleWareHandle) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if(!this.l.Borrow()){
		SendErrorResponse(w,http.StatusTooManyRequests,"too many request")
		return
	}
	this.r.ServeHTTP(w,r)
	this.l.Return()
}