package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "test-go-project/docs"
	"test-go-project/src/controllers"

	_ "github.com/jackc/pgx/v4/stdlib" // Импорт драйвера

	"github.com/gin-contrib/cors"
)

var db *sql.DB

func initDB() {
	var err error
	connectionString := "user=user password=password dbname=albums_db host=db port=5432 sslmode=disable"
	db, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal("Ошибка при подключении к БД:", err)
	}
}

func main (){
	initDB()
	defer db.Close()

	fmt.Println("code started")
	router:=gin.Default()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
	}))

	c := controllers.NewController(db)

	router.GET("/albums", c.GetAlbums)
	router.POST("/albums", c.PostAlbum)
	router.DELETE("albums/:id", c.DeleteAlbum)

	router.Run("0.0.0.0:8080")
}