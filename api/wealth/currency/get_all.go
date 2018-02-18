package currency

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const getCurencies = `
	SELECT * FROM currency
`

// GetCurrencies will get all currencies from database
func GetCurrencies(db *sqlx.DB) ([]models.Currency, error) {
	var currencies []models.Currency
	if err := db.Select(&currencies, getCurencies); err != nil {
		return nil, err
	}
	return currencies, nil
}

// GetAll all currencies
func (r *Routes) GetAll(c *gin.Context) {
	currencies, err := GetCurrencies(r.Db)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": currencies,
	})
}
