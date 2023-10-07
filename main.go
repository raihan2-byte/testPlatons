package main

import (
	"database/sql"
	"log"
	"net/http"
	"platons/platons"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/platons")
	if err != nil {
		log.Fatal("Error Connection to db", err)
	}
	defer db.Close()

	platonRepo := platons.NewPlatonRepository(db)

	r.POST("/biteship/callback/price", func (c *gin.Context){
		var input platons.Platons
		if err := c.ShouldBindJSON(&input);
		err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return 
		}
		err := platonRepo.Create(input)
		if err != nil {
			log.Println("Error create data", err)
			return
		}
		c.JSON(http.StatusCreated, gin.H{"Message": "Platons Created"})
	})

	r.GET("/biteship/callback/:DestinationPostalCode", func (c *gin.Context){
		postalCodestr := c.Param("DestinationPostalCode")
		postalCode, err := strconv.Atoi(postalCodestr)
		if err != nil {
			log.Println("invalid postalcode", err)
			return
		}

		platonData, err := platonRepo.GetPlatonByPostalCode(postalCode)
		if err != nil {
			log.Println("Error get data", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if platonData == nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"platon data not found"})
			return
		}
		c.JSON(http.StatusOK, platonData)
		
	})
	r.Run(":8080")
}