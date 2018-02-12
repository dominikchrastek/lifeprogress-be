package user

import (
	"lifeprogress/api/user/userWeight"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

// Register mounts all trips routes on the gin group.
func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	g := r.Group("/user")
	g.GET("", routes.GetAll)
	g.GET("/:id", routes.Get)
	userWeight.Register(g, db)
}
