package index

import (
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"main/lib/core/views"
	"main/lib/core/views/renders"
	"net/http"
)

func Get(
	render renders.Render,
) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		send.View(writer, request, render, views.View{Name: "Index"})
	}
}
