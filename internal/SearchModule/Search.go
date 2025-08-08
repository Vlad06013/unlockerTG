package SearchModule

//import (
//	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarMark"
//	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
//	"github.com/jinzhu/gorm"
//)
//
//var SearchString string
//var DbConnection *gorm.DB
//
//type SearchResult struct {
//	Type     string `json:"type"`
//	Entities []Entities
//}
//
//type Entities struct {
//	ID   uint64 `json:"id"`
//	Name string `json:"name"`
//}
//
//var searchResultSlice []SearchResult
//
//func Find() []SearchResult {
//	searchResultSlice = []SearchResult{}
//
//	findInCarMarks()
//	findInCarModels()
//	return searchResultSlice
//}
//
//func findInCarMarks() {
//	s := CarMark.Storage{DB: DbConnection}
//	var entities []Entities
//
//	carMarks := s.FindByName(SearchString)
//	for i := 0; i < len(carMarks); i++ {
//		entities = append(entities, Entities{
//			ID:   carMarks[i].ID,
//			Name: carMarks[i].Name,
//		})
//	}
//	searchResultSlice = append(searchResultSlice, SearchResult{
//		Type:     "carMark",
//		Entities: entities,
//	})
//}
//
//func findInCarModels() {
//	s := CarModel.Storage{DB: DbConnection}
//	var entities []Entities
//
//	carModels := s.FindByName(SearchString)
//	for i := 0; i < len(carModels); i++ {
//		entities = append(entities, Entities{
//			ID:   carModels[i].ID,
//			Name: carModels[i].Name,
//		})
//	}
//	searchResultSlice = append(searchResultSlice, SearchResult{
//		Type:     "carModel",
//		Entities: entities,
//	})
//}
