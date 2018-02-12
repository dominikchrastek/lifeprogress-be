package weight

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getQuery = `
	SELECT * FROM weight_with_unit ORDER BY timestamp
`

// GetAll weights
func (r *Routes) GetAll(c *gin.Context) {
	data := &[]models.Weight{}
	err := r.Db.Select(data, getQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
