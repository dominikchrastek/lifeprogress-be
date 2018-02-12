package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	// g := r.Group("/meta")
	r.GET("/meta", routes.Get)
}
