package wsource

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{Db: db}

	g := r.Group("/wsource")

	g.GET("/:id", routes.Get)
	g.GET("", routes.GetAll)
	g.POST("", routes.Post)
}
