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
	wsID := c.Param("ws_id")
	var data = map[string]interface{}{"user_id": userID, "ws_id": wsID}

	// insert connect weight with user
	if _, err := r.Db.NamedExec(connectorQuery, data); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
