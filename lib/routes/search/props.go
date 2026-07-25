package search

import "main/lib/schema"

type Props struct {
	Query string
	Items []schema.Result
}
