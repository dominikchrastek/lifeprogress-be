package wsource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteConnectorQuery = `
	DELETE FROM ws_user_connector WHERE user_id = $1 and ws_id = $2
`

func (r *Routes) Delete(c *gin.Context) {
	userID := c.Param("id")
	wsID := c.Param("ws_id")

	if _, err := r.Db.Exec(deleteConnectorQuery, userID, wsID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
