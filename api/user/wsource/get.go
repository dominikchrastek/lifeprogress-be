package wsource

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getSources = `
	SELECT
	 id,
	 name,
	 automated,
	 ws_type,
	FROM user_ws WHERE user_id = $1
`

// Get route
func (r *Routes) Get(c *gin.Context) {
	var wsource models.WSource
	wsourceID := c.Param("id")

	if err := r.Db.Get(&wsource, getSources, wsourceID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsource,
	})
}
