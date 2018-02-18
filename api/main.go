package api

import (
	"lifeprogress/api/life"
	"lifeprogress/api/meta"
	"lifeprogress/api/user"
	"lifeprogress/api/wealth"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Register do fap
func Register(app *gin.Engine, db *sqlx.DB) {
	apiGroup := app.Group("/api")

	meta.Register(apiGroup, db)
	user.Register(apiGroup, db)
	wealth.Register(apiGroup, db)
	life.Register(apiGroup, db)
}
