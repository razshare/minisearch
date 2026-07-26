package main

import (
	"embed"
	"errors"
	"log"
	"os"

	"main/lib/core/routes"
	"main/lib/core/servers"
	"main/lib/core/ssr"
	"main/lib/databases"
	"main/lib/routes/events"
	"main/lib/routes/fallback"
	"main/lib/routes/index"
	"main/lib/routes/search"
	"main/lib/schema"
)

//go:generate frizzante clean
//go:generate frizzante configure
//go:embed app/dist
var efs embed.FS
var errorLog = log.New(os.Stderr, "[error]: ", log.Ldate|log.Ltime)
var infoLog = log.New(os.Stdout, "[info]: ", log.Ldate|log.Ltime)
var database, databaseError = databases.Connect()
var queries = schema.New(database)
var render = ssr.New(ssr.Options{
	Efs:      efs,
	ErrorLog: errorLog,
	InfoLog:  infoLog,
	Limit:    1,
})
var appRoutes = []routes.Route{
	{
		Pattern: "GET /",
		Handler: fallback.Get(efs),
	},
	{
		Pattern: "GET /search",
		Handler: search.Get(queries, render, infoLog),
	},
	{
		Pattern: "GET /index",
		Handler: index.Get(render),
	},
	{
		Pattern: "POST /index",
		Handler: index.Post(queries, infoLog, errorLog, render),
	},
	{
		Pattern: "GET /events/index-progress",
		Handler: events.GetIndexProgress(infoLog),
	},
}
var startError = servers.Start(servers.StartOptions{
	ErrorLog: errorLog,
	InfoLog:  infoLog,
	Routes:   appRoutes,
	Address:  "127.0.0.1:38123",
})

func main() {
	if err := errors.Join(databaseError, startError); err != nil {
		log.Fatal(err)
	}
}
