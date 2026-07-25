package index

import (
	"log"
	"main/lib/core/receive"
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"main/lib/core/views"
	"main/lib/core/views/renders"
	"main/lib/schema"
	"main/lib/services"
	"net/http"
)

func Post(
	queries *schema.Queries,
	infoLog *log.Logger,
	errorLog *log.Logger,
	render renders.Render,
) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		var form Form
		_ = receive.Form(request, &form)
		_ = services.Index(
			request.Context(),
			queries,
			infoLog,
			errorLog,
			form.Address,
			form.Depth,
			map[string]string{},
		)
		infoLog.Printf("done indexing %s\n", form.Address)
		send.View(writer, request, render, views.View{
			Name: "Index",
		})
	}
}
