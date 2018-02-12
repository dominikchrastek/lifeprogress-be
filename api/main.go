package api

import (
	"lifeprogress/api/meta"
	"lifeprogress/api/user"
	"lifeprogress/api/weight"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Register do fap
func Register(app *gin.Engine, db *sqlx.DB) {
	apiGroup := app.Group("/api")
	weight.Register(apiGroup, db)
	meta.Register(apiGroup, db)
	user.Register(apiGroup, db)
}
