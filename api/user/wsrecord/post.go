package wsrecord

import (
	"lifeprogress/models"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

const createWSRecord = `
	INSERT INTO ws_record (user_id, ws_id, currency_id, value, timestamp)
	VALUES (:user_id, :ws_id, :currency_id, :value, current_timestamp)
	RETURNING id
`

// CreateWSRecord will insert WSRecord into the db
func CreateWSRecord(db *sqlx.DB, wsrecord models.WSRecordPost) (string, error) {
	stmt, err := db.PrepareNamed(createWSRecord)
	if err != nil {
		return "", err
	}
	var id string
	if err := stmt.QueryRow(&wsrecord).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

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
	id, err := CreateWSRecord(r.Db, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": id,
	})
}
