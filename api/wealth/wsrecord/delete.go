package wsrecord

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteWSRecord = `
	DELETE FROM ws_record WHERE id = $1
`

// Delete wsrecord
func (r *Routes) Delete(c *gin.Context) {

	if _, err := r.Db.Exec(deleteWSRecord, c.Param("id")); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
