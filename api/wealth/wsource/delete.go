package wsource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteQuery = `
	DELETE FROM wsource WHERE id = $1
`

const deleteConnectorQuery = `
	DELETE FROM ws_currency_connector WHERE ws_id = $1
`

func (r *Routes) Delete(c *gin.Context) {
	id := c.Param("id")

	tx, err := r.Db.Beginx()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if _, err := tx.Exec(deleteConnectorQuery, id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if _, err := tx.Exec(deleteQuery, id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = tx.Commit(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
