package currency

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteQuery = `
	DELETE FROM currency WHERE id = $1
`

// Delete currency
func (r *Routes) Delete(c *gin.Context) {

	if _, err := r.Db.Exec(deleteQuery, c.Param("id")); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
