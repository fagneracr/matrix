package ginserver

import (
	"log"
	"matrix/internal/database"
	"matrix/internal/verifymatrix"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type jsonInput struct {
	Letters []string `json:"letters"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := database.InitRedis()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTION"},
		AllowHeaders:     []string{"Origin", "access-control-allow-origin", "content-type", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           5 * time.Minute,
	}))
	r.OPTIONS("/", Options)
	r.POST("/sequence", func(c *gin.Context) {
		var letters jsonInput
		err := c.ShouldBindJSON(&letters)
		if err != nil {
			log.Println("Error:" + err.Error())
			c.JSON(http.StatusOK, gin.H{"is_valid": false})
			return
		}
		// Verify if input is valid,only contain letters authorized
		isvalidInput := verifymatrix.ArrayisValid(letters.Letters)
		if !isvalidInput {
			c.JSON(http.StatusOK, gin.H{"is_valid": false})
			return
		}
		//Verify sequence, if found 2 or more is valid
		countfind := 0
		countfind = countfind + verifymatrix.FindSequence(letters.Letters)
		// Arrange to vertical
		matrixv := verifymatrix.BuildVertical(letters.Letters)
		countfind = countfind + verifymatrix.FindSequence(matrixv)
		// Arrange to diagonal
		diagnais := verifymatrix.FindDiagonais(letters.Letters)
		countfind = countfind + verifymatrix.FindSequence(diagnais)
		//Output if is valid or not
		if countfind >= 2 {
			db.Set(letters.Letters, true)
			c.JSON(http.StatusOK, gin.H{"is_valid": true})
			return
		}
		db.Set(letters.Letters, false)
		c.JSON(http.StatusOK, gin.H{"is_valid": false})
		return
	})
	r.GET("/stats", func(c *gin.Context) {
		// Get in database data ratio
		valid, notvalid, ratio := db.ReturnStats()
		c.JSON(http.StatusOK, gin.H{
			"count_valid":   valid,
			"count_invalid": notvalid,
			"ration":        ratio,
		})
		return

	})
	return r
}

/*Options - Default Cors Allow*/
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}
