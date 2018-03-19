package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type fio struct {
	AccountStatement struct {
		Info struct {
			ClosingBalance float32 `json:"closingBalance"`
			Currency       string  `json:"currency"`
			DateEnd        string  `json:"dateEnd"`
			DateStart      string  `json:"dateStart"`
		} `json:"info"`
	} `json:"accountStatement"`
}

// Fio will get amount of money from Fio bank API
func Fio(token string) (float32, error) {
	today := time.Now()
	todayFormated := today.Format("2006-01-02")

	res, err := http.Get("https://www.fio.cz/ib_api/rest/periods/" + token + "/" + todayFormated + "/" + todayFormated + "/transactions.json")
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	textBytes := []byte(string(content))
	fio := fio{}

	if err = json.Unmarshal(textBytes, &fio); err != nil {
		return 0, err
	}

	return fio.AccountStatement.Info.ClosingBalance, nil
}
