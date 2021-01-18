package tvmaze

import (
	"log"

	"github.com/jgolang/config"
	searcher "github.com/jhuygens/searcher-engine"
)

var tvmazeSearcher = Searcher{}

func init() {
	err := searcher.RegisterSearcher(config.GetString("searchers.tvmaze"), tvmazeSearcher)
	if err != nil {
		log.Fatal(err)
		return
	}
}
