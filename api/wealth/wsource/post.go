package wsource

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const createWSource = `
	INSERT INTO wsource (name, ws_type)
	VALUES (:name, :ws_type)
	RETURNING id
`

const connectWSC = `
	INSERT INTO ws_currency_connector (currency_id, ws_id)
	VALUES (:currency_id, :ws_id)
`

// Post create wsource
func (r *Routes) Post(c *gin.Context) {
	var data models.WSource

	// reponse to JSON
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// create transaction
	tx, err := r.Db.Beginx()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// prepare wsource query
	stmt, err := tx.PrepareNamed(createWSource)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// id of created wsource
	var id string
	// execute wsource insert query
	if err := stmt.QueryRow(&data).Scan(&id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// connector query data
	cData := map[string]interface{}{"currency_id": data.CurrencyID, "ws_id": id}
	// prepare connector query
	cstmt, err := tx.PrepareNamed(connectWSC)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}
	// execute connector query
	if _, err := cstmt.Exec(cData); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// commit transaction
	if err = tx.Commit(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// send data
	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
