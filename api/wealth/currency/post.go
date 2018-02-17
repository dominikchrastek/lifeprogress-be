package currency

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const postQuery = `
	INSERT INTO currency (name)
	VALUES (:name)
	RETURNING id
`

// Post create currency record
func (r *Routes) Post(c *gin.Context) {

	var data models.Currency
	// reponse to JSON
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// prepare currency query
	stmt, err := r.Db.PrepareNamed(postQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// id of created currency
	var id string
	if err := stmt.QueryRow(&data).Scan(&id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
