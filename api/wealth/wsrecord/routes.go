package wsrecord

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Routes struct that contains db
type Routes struct {
	Db *sqlx.DB
}

// Register wsrecord routes
func Register(r *gin.RouterGroup, db *sqlx.DB) {
	routes := &Routes{db}

	g := r.Group("/wsrecord/:id")
	g.DELETE("", routes.Delete)
}
