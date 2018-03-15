package wsrecord

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const createWSRecord = `
	INSERT INTO ws_record (name, user_id, ws_id, currency_id, value, timestamp)
	VALUES (:name, :user_id, :ws_id, :currency_id, :value, current_timestamp)
	RETURNING id
`

// Post create wsource record
func (r *Routes) Post(c *gin.Context) {
	userID := c.Param("id")
	var data models.WSRecordPost
	data.UserID = userID
	// reponse to JSON
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// prepare wsRecord query
	stmt, err := r.Db.PrepareNamed(createWSRecord)
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
