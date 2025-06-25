package serv

import (
	"database/sql"
	"fuk/fukkk"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(args *string, db *sql.DB) {

	r := gin.Default() // log and recovery

	r.GET("/", func(c *gin.Context) {
		c.String(200, "use /args")
	})

	r.GET("/agrs", func(c *gin.Context) {
		text := *args
		c.String(http.StatusOK, "%v", text)

		if text != "" {
			err := fukkk.Logs(db, text, "вывел на локалхосте е боюшки")
			if err != nil {
				c.String(http.StatusInternalServerError, "db err oops\n %v", db)
			}
		}
	})
	r.Run(":8080")
}
