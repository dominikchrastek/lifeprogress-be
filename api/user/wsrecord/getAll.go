package wsrecord

import (
	"lifeprogress/models"
	"net/http"

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

// GetAll route
func (r *Routes) GetAll(c *gin.Context) {
	var wsources []models.WSRecordGet
	userID := c.Param("id")
	dates := pq.Array([]string{"2018-03-16 13:53:50.171906", "2018-03-15 13:55:50.171906"})

	if err := r.Db.Select(&wsources, getAllSources, userID, dates); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsources,
	})
}
