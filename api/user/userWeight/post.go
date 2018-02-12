package userWeight

import (
	"fmt"
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const postQuery = `
	INSERT INTO weight (value, unit, timestamp)
	VALUES (:value, :unit, current_timestamp)
	RETURNING id
`
const postQueryWithTimestamp = `
	INSERT INTO weight (value, unit, timestamp)
	VALUES (:value, :unit, :timestamp)
	RETURNING id
`
const connectorQuery = `
	INSERT INTO user_weight_connector (user_id, weight_id)
	VALUES (:user_id, :weight_id)
	RETURNING id
`
const getPostWeight = `
	SELECT
		id,
		value,
		unit,
		timestamp
	FROM user_weight WHERE user_id = $1 and id = $2
`

// Post create weight record
func (r *Routes) Post(c *gin.Context) {

	var data models.Weight
	userID := c.Param("id")

	// reponse to JSON
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// decide if timestamp was sent or not
	var query string
	if data.Timestamp == "" {
		query = postQuery
	} else {
		fmt.Print(data.Timestamp)
		query = postQueryWithTimestamp
	}

	// prepare weight query
	stmt, err := r.Db.PrepareNamed(query)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// id of created weight
	var id string
	err = stmt.QueryRow(&data).Scan(&id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// prepare connector query
	cStmt, err := r.Db.PrepareNamed(connectorQuery)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// unsert connect weight with user
	cData := models.WeightConnectUser{ID: 1, UserID: userID, WeightID: id}
	if _, err := cStmt.Exec(&cData); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// select data for response
	responseData := models.Weight{}
	err = r.Db.Get(&responseData, getPostWeight, userID, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": responseData,
	})
}
