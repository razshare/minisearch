package fallback

import (
	"embed"
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"net/http"
)

func Get(efs embed.FS) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		if found, _ := send.RequestedFile(writer, request, efs, "/"); !found {
			send.ToLocation(writer, "/search")
		}
	}
}
