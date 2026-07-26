package search

import "main/lib/schema"

type Props struct {
	Query        string
	Items        []schema.Result
	PagesCounter int64
	CurrentPage  int64
}
