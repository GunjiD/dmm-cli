package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/d-gunji/api/dmm-cli/actress"
)

//ActressSearch の結果を代入
type ActressSearch struct {
	Request struct {
		Parameters struct {
			AffiliateID string `json:"affiliate_id"`
			APIID       string `json:"api_id"`
			Keyword     string `json:"keyword"`
			Output      string `json:"output"`
		} `json:"parameters"`
	} `json:"request"`
	Result struct {
		Status        string `json:"status"`
		ResultCount   int    `json:"result_count"`
		TotalCount    string `json:"total_count"`
		FirstPosition int    `json:"first_position"`
		Actress       []struct {
			ID          string      `json:"id"`
			Name        string      `json:"name"`
			Ruby        string      `json:"ruby"`
			Bust        interface{} `json:"bust"`
			Waist       interface{} `json:"waist"`
			Hip         interface{} `json:"hip"`
			Height      interface{} `json:"height"`
			Birthday    interface{} `json:"birthday"`
			BloodType   interface{} `json:"blood_type"`
			Hobby       interface{} `json:"hobby"`
			Prefectures interface{} `json:"prefectures"`
			ListURL     struct {
				Digital string `json:"digital"`
				Monthly string `json:"monthly"`
				Ppm     string `json:"ppm"`
				Mono    string `json:"mono"`
				Rental  string `json:"rental"`
			} `json:"listURL"`
			ImageURL struct {
				Small string `json:"small"`
				Large string `json:"large"`
			} `json:"imageURL,omitempty"`
			Cup string `json:"cup,omitempty"`
		} `json:"actress"`
	} `json:"result"`
}

func main() {

	data := new(actress.ActressSearch)
	GetActress(data)
	fmt.Println(data.Result)
}

//GetActress は引数の data にGetRequestの結果を返します
func GetActress(data *actress.ActressSearch) {
	req, _ := http.NewRequest("GET", "https://api.dmm.com/affiliate/v3/ActressSearch", nil)

	query := req.URL.Query()
	query.Add("api_id", "")
	query.Add("affiliate_id", "")
	query.Add("keyword", "小倉")
	query.Add("output", "json")

	req.URL.RawQuery = query.Encode() //reqのクエリパラメータを設定。リファレンス参照

	client := new(http.Client)
	resp, _ := client.Get(req.URL.String())

	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)

	defer resp.Body.Close()
	json.Unmarshal(jsonBytes, data)

}
