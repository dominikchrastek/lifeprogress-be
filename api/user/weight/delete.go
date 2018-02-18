package weight

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const deleteQuery = `
	DELETE FROM weight WHERE id = $1
`

const deleteConnectorQuery = `
	DELETE FROM user_weight_connector WHERE weight_id = $1
`

// Delete weight
func (r *Routes) Delete(c *gin.Context) {

	// create transaction
	tx, err := r.Db.Beginx()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// delete from 'user_weight_connector'
	if _, err := tx.Exec(deleteConnectorQuery, c.Param("weight-id")); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// delete from 'weight'
	if _, err := tx.Exec(deleteQuery, c.Param("weight-id")); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
