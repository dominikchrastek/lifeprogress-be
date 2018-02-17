package wsource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const connectorQuery = `
	INSERT INTO ws_user_connector (user_id, ws_id)
	VALUES (:user_id, :ws_id)
	RETURNING id
`

// Post create wsource record
func (r *Routes) Post(c *gin.Context) {
	userID := c.Param("id")
	var data = map[string]interface{}{"user_id": userID, "ws_id": ""}

	// reponse to JSON
	if err := c.BindJSON(data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// prepare wsource query
	stmt, err := r.Db.PrepareNamed(connectorQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var id string
	if err := stmt.QueryRow(&data).Scan(&id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
