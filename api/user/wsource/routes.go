package wsource

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Routes struct that contains db
type Routes struct {
	Db *sqlx.DB
}

// Register wsource aroutes
func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	g := r.Group("/:id/wsource")
	g.POST("/:ws_id", routes.Post)
	g.DELETE("/:ws_id", routes.Delete)
	g.GET("", routes.GetAll)
}
