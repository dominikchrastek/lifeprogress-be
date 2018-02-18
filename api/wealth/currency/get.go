package currency

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getCurency = `
	SELECT * FROM currency WHERE id = $1
`

// Get currency
func (r *Routes) Get(c *gin.Context) {
	var currency models.Currency
	var currencyID = c.Param("id")
	if err := r.Db.Get(&currency, getCurency, currencyID); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": currency,
	})
}
