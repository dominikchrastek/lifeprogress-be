package user

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getAllQuery = `
	SELECT * FROM user_profile
`

// GetAll route
func (r *Routes) GetAll(c *gin.Context) {
	var users []models.User
	if err := r.Db.Select(&users, getAllQuery); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
