package wsource

import (
	"lifeprogress/api/wealth/wsource"
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const getAllSources = `
	SELECT
	 id,
	 name,
	 automated,
	 ws_type
	FROM user_ws WHERE user_id = $1
`

func GetUserWSources(db *sqlx.DB, userID string) ([]models.WSourceC, error) {
	var wsources []models.WSource

	if err := db.Select(&wsources, getAllSources, userID); err != nil {
		return nil, err
	}

	wsourcesc, err := wsource.GetWSourcesC(db, wsources)
	if err != nil {
		return nil, err
	}
	return wsourcesc, nil
}

// GetAll route
func (r *Routes) GetAll(c *gin.Context) {
	userID := c.Param("id")

	wsWithCurrencies, err := GetUserWSources(r.Db, userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": wsWithCurrencies,
	})
}
