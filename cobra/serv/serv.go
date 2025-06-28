package serv

import (
	"database/sql"
	"fmt"
	"fuk/fukkk"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server(db *sql.DB) {

	r := gin.Default()
	r.LoadHTMLGlob("front/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "calendar.html", nil)
	})
	// r.GET("/calendar", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "calendar.html", nil)
	// })

	// r.GET("/test", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "test ok")
	// })
	// r.POST("/submit-date", func(c *gin.Context) {
	// 	var data struct {
	// 		Date string `json:"date"`
	// 	}

	// 	if err := c.BindJSON(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"message": "ur date: " + data.Date})
	// })

	r.POST("/log", func(c *gin.Context) {
		fmt.Println("POST /log received")

		var dataDB struct {
			Date string `json:"date"`
			Task string `json:"task"`
		}

		if err := c.BindJSON(&dataDB); err != nil || dataDB.Task == "" {
			fmt.Println("Ошибка при записи в БД:", err)
			c.JSON(http.StatusBadRequest, gin.H{"err": "invalid format"})
			return
		}

		input := dataDB.Task
		date := dataDB.Date
		output := "correct"

		err := fukkk.Logs(db, input, output, date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "cannot put logs in db"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"mess":   "Data added",
			"input":  input,
			"output": output,
			"date":   date,
		})
	})
	r.GET("/log", func(c *gin.Context) {
		logs, err := fukkk.GetAllLogs(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": "cannot get logs from db"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"logs": logs})
	})

	r.Run(":8080")
}
