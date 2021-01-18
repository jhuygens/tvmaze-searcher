package tvmaze

import (
	"github.com/jgolang/config"
	"github.com/jgolang/log"
	searcher "github.com/jhuygens/searcher-engine"
)

var tvmazeSearcher = Searcher{}

func init() {
	name := config.GetString("searchers.tvmaze")
	err := searcher.RegisterSearcher(name, tvmazeSearcher)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("Searcher %v has been register", name)
}
