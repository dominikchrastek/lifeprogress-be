package wsource

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const createWSource = `
	INSERT INTO wsource (name, ws_type)
	VALUES (:name, :ws_type)
	RETURNING id
`

// Post create wsource
func (r *Routes) Post(c *gin.Context) {
	var wsource models.WSourceC

	// reponse to JSON
	if err := c.BindJSON(&wsource); err != nil {
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

	defer stmt.Close()

	// id of created wsource
	var wsourceID string
	// execute wsource insert query
	if err := stmt.QueryRow(&wsource).Scan(&wsourceID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	// prepare connector query (bulk)
	cstmt, err := tx.Prepare(pq.CopyIn("ws_currency_connector", "currency_id", "ws_id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		tx.Rollback()
		return
	}

	defer cstmt.Close()

	// exec bulk stuff
	for _, currency := range wsource.Currencies {
		_, err = cstmt.Exec(currency.ID, wsourceID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}

	// execute connector insert (bulk)
	if _, err := cstmt.Exec(); err != nil {
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
		"data": wsourceID,
	})
}
