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

	if err := c.ShouldBindJSON(currency); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
