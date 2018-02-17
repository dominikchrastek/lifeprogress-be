package meta

import (
	"lifeprogress/api/wealth/currency"
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnitsQuery SQL
const UnitsQuery = `
	SELECT * FROM weight_unit
`

// AssetTypeQuery SQL
const wsTypeQuery = `
	SELECT * FROM ws_type
`

// Get route
func (r *Routes) Get(c *gin.Context) {
	var units []models.Unit
	var ws_types []models.WSType

	if err := r.Db.Select(&units, UnitsQuery); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := r.Db.Select(&ws_types, wsTypeQuery); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	currencies, err := currency.GetCurrencies(r.Db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"units":      units,
		"ws_types":   ws_types,
		"currencies": currencies,
	})
}
