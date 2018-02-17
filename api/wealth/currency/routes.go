package currency

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{Db: db}

	g := r.Group("/currency")

	g.POST("", routes.Post)

}
