package main

import (
	"main/lib/core/types"
	"main/lib/routes/index"
	"main/lib/routes/search"
)

func main() {
	_ = types.Clear()
	_ = types.Generate[search.Props]()
	_ = types.Generate[search.Form]()
	_ = types.Generate[index.Form]()
	_ = types.Generate[index.Progress]()
}
