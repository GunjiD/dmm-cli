package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/d-gunji/api/dmm-cli/actress"
)

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
