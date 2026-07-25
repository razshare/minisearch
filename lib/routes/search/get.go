package search

import (
	"main/lib/core/receive"
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"main/lib/core/views"
	"main/lib/core/views/renders"
	"main/lib/schema"
	"net/http"
)

func Get(
	queries *schema.Queries,
	render renders.Render,
) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		var form Form
		var items []schema.Result
		_ = receive.Form(request, &form)
		if form.Query != "" {
			items, _ = queries.FindResultsByDescription(request.Context(), "%"+form.Query+"%")
		}
		_ = send.View(writer, request, render, views.View{
			Name: "Search",
			Props: Props{
				Query: form.Query,
				Items: items,
			},
		})
	}
}
