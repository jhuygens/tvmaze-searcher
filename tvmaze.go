package tvmaze

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jgolang/config"
	"github.com/jgolang/consumer/rest"
	"github.com/jgolang/log"
	searcher "github.com/jhuygens/searcher-engine"
)

// Limit returns limit
var Limit = "20"

// Searcher searcher interface implement
type Searcher struct{}

// Search doc ...
func (s Searcher) Search(filter searcher.Filter) ([]searcher.Item, error) {
	if len(filter.Types) == 0 {
		filter.Types = append(filter.Types, "show")
	}
	var items []searcher.Item
	for _, resource := range filter.Types {
		if resource == "show" {
			for _, name := range filter.Name {
				queryParams := make(map[string]string)
				queryParams["limit"] = "20"
				queryParams["q"] = name.Value
				results, err := searchTVMAZEServiceItems(queryParams)
				if err != nil {
					log.Error(err)
				}
				items = append(items, results...)
			}
		}
	}
	return items, nil
}

func searchTVMAZEServiceItems(queryParams map[string]string) ([]searcher.Item, error) {
	resquest := rest.RequestInfo{
		Method:      http.MethodGet,
		Endpoint:    config.GetString("integrations.tvmaze.endpoint"),
		Timeout:     time.Duration(config.GetInt("integrations.tvmaze.timeout")) * time.Second,
		QueryParams: queryParams,
	}
	var responseData ResponseData
	response, err := rest.ConsumeRestService(resquest, &responseData)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code %v", response.StatusCode)
	}
	var items []searcher.Item
	for _, result := range responseData {
		show := result.Show
		items = append(
			items,
			searcher.Item{
				Type:    "show",
				Library: config.GetString("searchers.tvmaze"),
				Name:    show.Name,
				Artwork: show.Image.Medium,
				Info: searcher.Info{
					PreviewURL:  show.URL,
					Title:       show.Name,
					Collection:  show.Type,
					Artist:      show.Network.Name,
					Languages:   []string{show.Language},
					RatingAvg:   show.Rating.Average,
					Genres:      show.Genres,
					Description: show.Summary,
					MoreInfo:    show.OfficialSite,
					ReleaseDate: show.Premiered,
					Country:     show.Network.Country.Code,
					Price:       0.0,
					RentalPrice: 0.0,
					Currency:    "USD",
					URL:         show.URL,
				},
			},
		)
	}
	return items, nil
}
