package wsource

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getWSource = `
	SELECT
	 id,
	 name,
	 ws_type
	FROM ws_with_type WHERE id = $1
`

// Get wsource
func (r *Routes) Get(c *gin.Context) {
	var wsource models.WSource
	wsourceID := c.Param("id")

	if err := r.Db.Get(&wsource, getWSource, wsourceID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsource,
	})
}
