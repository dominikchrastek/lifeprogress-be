package wealth

import (
	"lifeprogress/api/wealth/currency"
	"lifeprogress/api/wealth/currencyRate"
	"lifeprogress/api/wealth/wsource"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Routes struct {
	Db *sqlx.DB
}

func Register(r *gin.RouterGroup, db *sqlx.DB) {
	wsource.Register(r, db)
	currency.Register(r, db)
	currencyRate.Register(r, db)
}
