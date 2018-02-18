package weight

import (
	"net/http"

	"lifeprogress/models"

	"github.com/gin-gonic/gin"
)

const putQuery = `
	UPDATE weight SET
		value = :value,
		unit = :unit
	WHERE id = :id
`

// Put update weight
func (r *Routes) Put(c *gin.Context) {
	weight := &models.Weight{}

	if err := c.BindJSON(weight); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	weight.ID = c.Param("id")
	_, err := r.Db.NamedExec(putQuery, weight)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
