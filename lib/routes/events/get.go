package events

import (
	"log"
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"net/http"

	uuid "github.com/nu7hatch/gouuid"
)

func GetIndexProgress(infoLog *log.Logger) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		id, _ := uuid.NewV4()
		go func() {
			Mutex.Lock()
			defer Mutex.Unlock()
			_ = send.SseUpgrade(&writer)
			Writers[id.String()] = writer
		}()
		<-request.Context().Done()
		Mutex.Lock()
		defer Mutex.Unlock()
		delete(Writers, id.String())
	}
}
