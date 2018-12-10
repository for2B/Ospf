package middleware

import (
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"sync"
	"strconv"
	"OutsourcingPlatform/ospf/constant"
)

var recordOnlineCount = 0
var lock = new(sync.Mutex)

func Online(next httpapi.APIHandler) httpapi.APIHandler {

	return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if recordOnlineCount==constant.MAX_ONLINE{
		return nil,httpapi.NewErr(3,"Online users count is more than "+strconv.Itoa(constant.MAX_ONLINE)+"! Please try again later.",nil)
	}
		lock.Lock()
		recordOnlineCount++
		lock.Unlock()
		response, err := next(w, r)
		lock.Lock()
		recordOnlineCount--
		//fmt.Println("users online:" + strconv.Itoa(recordOnlineCount))
		lock.Unlock()
		return response, err
	}
}
