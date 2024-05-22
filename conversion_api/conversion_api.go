package conversion_api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"remotedevteam-task/models"
)

type SuccessResponse struct {
	Data map[string]float64 `json:"data"`
}

func Fetch(r models.Request) (float64, error) {
	baseURL := "https://api.freecurrencyapi.com/v1/latest"
	params := url.Values{}
	params.Add("apikey", "fca_live_1pgI4XExLrjB9q7DjEhIUnob1btztjnG2H8JFYgC")
	params.Add("currencies", r.TargetCurrency)
	params.Add("base_currency", r.BaseCurrency)

	resp, err := http.Get(baseURL + "?" + params.Encode())
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	if resp.StatusCode == http.StatusOK {
		var response SuccessResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return 0, err
		}
		value, ok := response.Data[r.TargetCurrency]
		if !ok {
			return 0, nil
		}
		return value * r.Amount, nil
	} else {
		return 0, fmt.Errorf("Such currencies are not present ")
	}

}
