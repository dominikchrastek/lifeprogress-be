package currency

import (
	"net/http"

	"lifeprogress/models"

	"github.com/gin-gonic/gin"
)

const putQuery = `
	UPDATE currency SET
		name = :name
	WHERE id = :id
`

// Put update currency
func (r *Routes) Put(c *gin.Context) {
	currency := &models.Currency{}

	if err := c.BindJSON(currency); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	currency.ID = c.Param("id")
	_, err := r.Db.NamedExec(putQuery, currency)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
