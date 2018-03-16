package wsrecord

import (
	"lifeprogress/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const getAllSources = `
	SELECT
	 id,
	 name,
	 value,
	 timestamp,
	 currency_name
	FROM wsr_currency WHERE user_id = $1 and DATE(timestamp) = ANY($2)
`

// getDateInterval will return []time.Time and these times are between from - to (included these values)
func getDateInterval(dateFrom string, dateTo string, array []time.Time) ([]time.Time, error) {

	from, err := time.Parse(time.RFC3339, dateFrom)
	if err != nil {
		return array, err
	}
	// parse toDate to time type
	to, err := time.Parse(time.RFC3339, dateTo)
	if err != nil {
		return array, err
	}

	for from.Before(to) {
		array = append(array, from)
		from = from.Add(time.Hour * 24)
		if from.Equal(to) {
			array = append(array, to)
		}
	}

	return array, nil
}

// GetAll route
// Params:
// 	id: userId
//	from: start of interval
//	to: end of interval
// records in this interval will be returned
func (r *Routes) GetAll(c *gin.Context) {
	var wsources []models.WSRecordGet
	userID := c.Param("id")
	dateFrom := c.Query("from")
	dateTo := c.Query("to")

	var arrayTimes []time.Time
	if dateFrom != "" && dateTo != "" {
		newArrayTimes, err := getDateInterval(dateFrom, dateTo, arrayTimes)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		arrayTimes = newArrayTimes
	}

	dates := pq.Array(arrayTimes)

	if err := r.Db.Select(&wsources, getAllSources, userID, dates); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsources,
	})
}
