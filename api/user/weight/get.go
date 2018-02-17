package weight

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getWeights = `
	SELECT
	 id,
	 value,
	 unit,
	 timestamp
	FROM user_weight WHERE user_id = $1
	ORDER BY timestamp
`

// Get route
func (r *Routes) Get(c *gin.Context) {
	var weights []models.Weight
	id := c.Param("id")
	err := r.Db.Select(&weights, getWeights, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": weights,
	})
}
