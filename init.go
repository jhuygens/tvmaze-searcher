package tvmaze

import (
	"github.com/jgolang/config"
	"github.com/jgolang/log"
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
