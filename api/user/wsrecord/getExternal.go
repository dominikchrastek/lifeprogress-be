package wsrecord

import (
	"lifeprogress/api/wealth/wsource"
	"lifeprogress/data"
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getAutomatedWSources = `
	SELECT
	 id,
	 name,
	 automated,
	 ws_type
	FROM user_ws WHERE automated = TRUE and user_id = $1
`

// GetExternal will get data from sources, that are already automated, atm just fio bank
func (r *Routes) GetExternal(c *gin.Context) {
	userID := c.Param("id")

	var wsources []models.WSource
	if err := r.Db.Select(&wsources, getAutomatedWSources, userID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	wsourcesC, err := wsource.GetWSourcesC(r.Db, wsources)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	for _, wsource := range wsourcesC {
		// todo store token somehwere and hash it or something
		value, _ := data.Fio("")
		record := models.WSRecordPost{
			WSRecordCommon: models.WSRecordCommon{Value: value},
			WSourceID:      wsource.ID,
			CurrencyID:     wsource.Currencies[0].ID,
			UserID:         userID,
		}
		_, err := CreateWSRecord(r.Db, record)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	c.Status(http.StatusNoContent)
}
