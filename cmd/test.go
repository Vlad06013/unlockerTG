package main

import (
	"fmt"
	"github.com/Vlad06013/unlockerTG.git/repository/entities/CarModel"
)

func main() {
	//// URL API
	//url := "http://127.0.0.1/api/telegram-bot/v1/car-marks?limit=2&page=3"
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Println("Ошибка при отправке запроса:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Получен неожиданный статус:", resp.Status)
	//	return
	//}
	//
	//decoder := json.NewDecoder(resp.Body)
	//
	//var result ApiResponse
	//err = decoder.Decode(&result)
	//if err != nil {
	//	fmt.Println("Ошибка при парсинге JSON:", err)
	//	return
	//}
	//
	//// Теперь у вас есть массив items внутри result.Data
	//for _, item := range result.Data {
	//	fmt.Printf("ID: %d, Name: %s\n", item.ID, item.Name)
	//}

	s := CarModel.GetByMarkId(1, 1, 2)
	fmt.Println("APIRESP", s)

}
