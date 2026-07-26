package services

import (
	"context"
	"log"
	"main/lib/schema"
	"strings"
	"sync"

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
	mutex *sync.Mutex,
	onProgress func(current int, maximum int),
) (err error) {
	var current int
	var maximum int
	if depth == 0 {
		return
	}
	if tracked == nil {
		tracked = make(map[string]string, 0)
	}
	addresses := make([]string, 0)
	collector := colly.NewCollector()
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		address := element.Attr("href")
		if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
			return
		}
		maximum += 1
		addresses = append(addresses, address)
	})
	collector.OnHTML("title", func(element *colly.HTMLElement) {
		maximum += 1
		current += 1
		mutex.Lock()
		defer mutex.Unlock()
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
	var lock sync.Mutex
	var group sync.WaitGroup
	for _, address := range addresses {
		group.Go(func() {
			Index(ctx, queries, infoLog, errorLog, address, depth-1, tracked, mutex, nil)
			if onProgress != nil {
				group.Go(func() {
					lock.Lock()
					current++
					onProgress(current, maximum)
					lock.Unlock()
				})
			}
		})
	}
	group.Wait()
	return
}
