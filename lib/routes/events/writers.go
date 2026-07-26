package events

import "net/http"

var Writers = make(map[string]http.ResponseWriter, 0)
