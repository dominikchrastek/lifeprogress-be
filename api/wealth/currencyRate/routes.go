package currencyRate

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{Db: db}

	g := r.Group("/currency-rate")

	// g.POST("", routes.Post)
	// g.PUT("/:id", routes.Put)
	// g.DELETE("/:id", routes.Delete)
	// g.GET("", routes.GetAll)
	g.GET("/task", routes.Task)

}
