package datacollectors

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// CollectData to get stock data
func CollectData() {
	resp, err := http.Get("https://www.google.com/async/finance_wholepage_price_updates?ei=qfdnW6GSCuOT0PEP2cmqkAg&yv=3&async=mids:%2Fg%2F1dv0gb5j%7C%2Fg%2F1dv1hvg6%7C%2Fg%2F1dty2r0m%7C%2Fg%2F1hbvps5v1%7C%2Fg%2F1dvcb0kb%7C%2Fg%2F1hbvy4_z_%7C%2Fm%2F07zm276%7C%2Fm%2F046k_p%7C%2Fm%2F04t5sp,currencies:,_fmt:jspb")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body[:])
	bodyStr = strings.Replace(bodyStr, ")]}'", "", 1)
	convertJSON(bodyStr)
}

// https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
func convertJSON(parsedJSON string) float64 {
	var result map[string]interface{}
	json.Unmarshal([]byte(parsedJSON), &result)
	PriceUpdate := result["PriceUpdate"].([]interface{})
	zerothEle := PriceUpdate[0].([]interface{})
	zerothEle = zerothEle[0].([]interface{})
	zerothEle = zerothEle[0].([]interface{})
	zerothEle = zerothEle[17].([]interface{})
	hdfc := zerothEle[4].(float64)
	return hdfc
}
