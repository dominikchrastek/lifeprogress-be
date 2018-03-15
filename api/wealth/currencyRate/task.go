package currencyRate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// https://api.fixer.io/
type Response struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func GetRate(dateString string, from string, to string) (float64, error) {
	var date string
	if t, err := time.Parse(time.RFC3339, dateString); err != nil {
		date = "latest"
	} else {
		date = t.Format("2006-01-02")
	}
	url := "https://api.fixer.io/" + date + "?base=" + strings.ToUpper(from) + "&symbols=" + strings.ToUpper(to)

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	content, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	textBytes := []byte(string(content))
	rates := Response{}

	jsonErr := json.Unmarshal(textBytes, &rates)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return 0, err
	}

	fmt.Println(rates)
	res, err = http.Get(url)
	if err != nil {
		return 0, err
	}

	return rates.Rates[to], nil
}

func (r *Routes) Task(c *gin.Context) {
	date := c.Query("date")
	kek, err := GetRate(date, "CZK", "USD")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(kek)
}
