package user

import (
	"lifeprogress/api/user/weight"
	"lifeprogress/api/user/wsource"
	"lifeprogress/api/user/wsrecord"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Routes struct that contains db
type Routes struct {
	Db *sqlx.DB
}

// Register mounts all trips routes on the gin group.
func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	g := r.Group("/user")
	g.GET("", routes.GetAll)
	g.GET("/:id", routes.Get)
	wsource.Register(g, db)
	weight.Register(g, db)
	wsrecord.Register(g, db)
}
