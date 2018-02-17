package user

import (
	"lifeprogress/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getQuery = `
	SELECT * FROM user_profile WHERE id = $1
`

// Get route
func (r *Routes) Get(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	if err := r.Db.Get(&user, getQuery, userID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
