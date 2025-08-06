package SearchModule

import (
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarMark"
	"github.com/jinzhu/gorm"
)

var SearchString string
var DbConnection *gorm.DB

type SearchResult struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

var searchResultSlice []SearchResult

func Find() []SearchResult {
	findInCarMarks()
	return searchResultSlice
}

func findInCarMarks() {
	s := CarMark.Storage{DB: DbConnection}
	carMarks := s.FindByName(SearchString)
	for i := 0; i < len(carMarks); i++ {
		searchResultSlice = append(searchResultSlice, SearchResult{
			Type:  "carMark",
			Value: carMarks[i],
		})
	}
}
