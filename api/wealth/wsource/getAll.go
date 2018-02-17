package wsource

import (
	"fmt"
	"lifeprogress/models"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

const getWSources = `
	SELECT
	 id,
	 name,
	 ws_type
	FROM ws_with_type
`

const getCurrencies = `
	SELECT
	 id,
	 name
	FROM ws_currency WHERE  ws_id = $1
`

// GetWSourcesC get WSources with currencies
func GetWSourcesC(db *sqlx.DB, wsources []models.WSource) ([]models.WSourceC, error) {
	var wsWithCurrencies []models.WSourceC
	for _, source := range wsources {
		var currency []models.Currency
		fmt.Println(source)
		if err := db.Select(&currency, getCurrencies, source.ID); err != nil {
			return nil, err
		}
		wsWithCurrencies = append(wsWithCurrencies, models.WSourceC{models.WSourceCommon{source.ID, source.Name, source.Type}, currency})
	}

	return wsWithCurrencies, nil
}

// GetAll wsources
func (r *Routes) GetAll(c *gin.Context) {
	var wsources []models.WSource
	if err := r.Db.Select(&wsources, getWSources); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	wsWithCurrencies, err := GetWSourcesC(r.Db, wsources)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsWithCurrencies,
	})
}
