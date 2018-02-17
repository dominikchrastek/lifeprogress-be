package wsource

import (
	"lifeprogress/api/wealth/wsource"
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getAllSources = `
	SELECT
	 id,
	 name,
	 ws_type
	FROM user_ws WHERE user_id = $1
`

// GetAll route
func (r *Routes) GetAll(c *gin.Context) {
	var wsources []models.WSource
	userID := c.Param("id")

	if err := r.Db.Select(&wsources, getAllSources, userID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	wsWithCurrencies, err := wsource.GetWSourcesC(r.Db, wsources)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsWithCurrencies,
	})
}
