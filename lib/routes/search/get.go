package search

import (
	"log"
	"main/lib/core/receive"
	"main/lib/core/routes"
	"main/lib/core/scopes"
	"main/lib/core/send"
	"main/lib/core/views"
	"main/lib/core/views/renders"
	"main/lib/schema"
	"math"
	"net/http"
	"strings"
)

func Get(
	queries *schema.Queries,
	render renders.Render,
	infoLog *log.Logger,
) routes.Handler {
	return func(
		scope scopes.Scope,
		request *http.Request,
		writer http.ResponseWriter,
	) {
		var form Form
		var items []schema.Result
		_ = receive.Form(request, &form)
		if form.Page <= 0 {
			form.Page = 1
		}
		var pagesCounter int64 = 1
		var currentPage int64 = form.Page
		if form.Query != "" {
			chunks := strings.Split(form.Query, " ")
			query := "%" + strings.Join(chunks, "%") + "%"
			infoLog.Printf("querying for %s\n", query)
			count, _ := queries.CountResultsByDescription(
				request.Context(),
				query,
			)
			pagesCounter = int64(math.Ceil(float64(count) / float64(10)))
			if currentPage > pagesCounter {
				currentPage = pagesCounter
			}
			items, _ = queries.FindResultsByDescription(
				request.Context(),
				schema.FindResultsByDescriptionParams{
					Description: query,
					Offset:      currentPage,
					Count:       10,
				},
			)
		}
		_ = send.View(writer, request, render, views.View{
			Name: "Search",
			Props: Props{
				Query:        form.Query,
				Items:        items,
				PagesCounter: pagesCounter,
				CurrentPage:  currentPage,
			},
		})
	}
}
