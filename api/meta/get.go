package meta

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UnitsQuery SQL
const UnitsQuery = `
	SELECT * FROM weight_unit
`

// AssetTypeQuery SQL
const AssetTypeQuery = `
	SELECT * FROM asset_type
`

// Get route
func (r *Routes) Get(c *gin.Context) {
	var units []models.Unit
	var assets []models.Asset
	err := r.Db.Select(&units, UnitsQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	err = r.Db.Select(&assets, AssetTypeQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"units":  units,
		"assets": assets,
	})
}
