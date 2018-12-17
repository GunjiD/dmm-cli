package main

import (
	"net/http/httputil"
	//	"net/url"
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	query := req.URL.Query()
	query.Add("key", "")

	req.URL.RawQuery = query.Encode() //reqのクエリパラメータを設定。リファレンス参照
	fmt.Println(req.URL.String())
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
	//関数を抜ける際にresponceをclose
	defer resp.Body.Close()

}