package streamserver

import (
	"os"
	"time"
	"net/http"
	"github.com/julienschmidt/httprouter"
)


func StreamHandler(w http.ResponseWriter, r *http.Request,p httprouter.Params)  {
	var videoId string = p.ByName("vid-id")
	vl := VIDEO_DIR + videoId
	video,err := os.Open(vl)
	defer video.Close()
	if(err!= nil){
		SendErrorResponse(w,http.StatusInternalServerError,"Video Not Exist")
		return
	}
	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)
}

func UploadHandle(w http.ResponseWriter, r *http.Request,p httprouter.Params)  {
	
}