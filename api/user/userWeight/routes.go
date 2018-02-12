package userWeight

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	g := r.Group("/:id/weight")
	g.GET("", routes.Get)
	g.POST("", routes.Post)
}
