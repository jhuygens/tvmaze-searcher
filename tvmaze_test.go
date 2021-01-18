package tvmaze

import (
	"reflect"
	"testing"

	searcher "github.com/jhuygens/searcher-engine"
)

func Test_searchTVMAZEServiceItems(t *testing.T) {
	query := make(map[string]string)
	query["limit"] = "20"
	query["q"] = "casa"
	// query["entity"] = "movie"
	// query["country"] = "GT"
	type args struct {
		queryParams map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    []searcher.Item
		wantErr bool
	}{
		{
			name: "hola",
			args: args{
				queryParams: query,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := searchTVMAZEServiceItems(tt.args.queryParams)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchTVMAZEServiceItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchTVMAZEServiceItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearcher_Search(t *testing.T) {
	type args struct {
		filter searcher.Filter
	}
	tests := []struct {
		name    string
		s       Searcher
		args    args
		want    []searcher.Item
		wantErr bool
	}{
		{
			name: "prueba inicial",
			s:    Searcher{},
			args: args{
				filter: searcher.Filter{
					Name:   []searcher.FieldValue{{Value: "hola"}, {Value: "como"}},
					Artist: []searcher.FieldValue{{Value: "nirvana"}},
					Types:  []string{"show", "track"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Searcher{}
			got, err := s.Search(tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Searcher.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Searcher.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
