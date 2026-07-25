package services

import (
	"context"
	"log"
	"main/lib/schema"
	"strings"

	"github.com/gocolly/colly/v2"
	uuid "github.com/nu7hatch/gouuid"
)

func Index(
	ctx context.Context,
	queries *schema.Queries,
	infoLog *log.Logger,
	errorLog *log.Logger,
	address string,
	depth int,
	tracked map[string]string,
) (err error) {
	if depth == 0 {
		return
	}
	if tracked == nil {
		tracked = make(map[string]string, 0)
	}
	collector := colly.NewCollector()
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		address := element.Attr("href")
		if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
			return
		}
		Index(ctx, queries, infoLog, errorLog, address, depth-1, tracked)
	})
	collector.OnHTML("title", func(element *colly.HTMLElement) {
		if _, exists := tracked[element.Text]; exists {
			return
		}
		id, _ := uuid.NewV4()
		if err := queries.AddResult(ctx, schema.AddResultParams{
			ID:          id.String(),
			Address:     address,
			Description: element.Text,
		}); err != nil {
			errorLog.Printf("error while indexing %s: %v", address, err)
			return
		}
		tracked[element.Text] = address
		infoLog.Printf("indexing %s\n", address)
	})
	collector.Visit(address)
	return
}
