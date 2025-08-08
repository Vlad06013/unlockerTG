package ApiClientBackend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiResponse struct {
	Data []map[string]interface{} `json:"data"`
}

var url = "http://127.0.0.1/api/telegram-bot/v1/"

func send(uri string) (response *http.Response, err error) {
	resp, err := http.Get(url + uri)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Получен неожиданный статус:", resp.Status)
		defer resp.Body.Close()

		return
	}
	return resp, err
}

func decode(resp *http.Response) (response ApiResponse, err error) {
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var result ApiResponse
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	return result, err
}

func Get(uri string) ApiResponse {
	resp, _ := send(uri)
	result, _ := decode(resp)

	return result
}
