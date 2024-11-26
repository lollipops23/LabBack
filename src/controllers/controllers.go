package controllers

import (
	"database/sql"
	"net/http"

	"test-go-project/src"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	db *sql.DB
}

func NewController(db *sql.DB) *Controller {
	return &Controller{db: db}
}

func (c *Controller) GetAlbums(context *gin.Context) {
	field := context.Query("field")
	sqlOp := context.Query("sql_op")
	value := context.Query("value")

	var query string
	var args []interface{}

	// Проверяем, поступили ли необходимые параметры
	if field != "" && sqlOp != "" && value != "" {
		// Генерация SQL-запроса в зависимости от sql_op
		if sqlOp == "IN" {
			query = "SELECT id, fullname, record_book, birth_date, create_date FROM albums WHERE " + field + " = $1"
			args = append(args, value)
		} else if sqlOp == "NOT_IN" {
			query = "SELECT id, fullname, record_book, birth_date, create_date FROM albums WHERE " + field + " != $1"
			args = append(args, value)
		} else {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Недопустимая операция фильтрации"})
			return
		}
	} else {
		// Если параметры не указаны, вернуть все записи
		query = "SELECT id, fullname, record_book, birth_date, create_date FROM albums"
	}

	rows, err := c.db.Query(query, args...)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo": err.Error()})
		return
	}
	defer rows.Close()

	var albums []src.ALBUM
	for rows.Next() {
		var album src.ALBUM
		if err := rows.Scan(&album.ID, &album.Fullname, &album.RecordBook, &album.BirthDate, &album.CreateDate); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo":err.Error()})
			return
		}
		albums = append(albums, album)
	}

	if albums!=nil{
		context.IndentedJSON(http.StatusOK, albums)
		return
	}

	context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Фильтрация не обнаружила объектов"})

	
}

// Метод для добавления нового альбома
func (c *Controller) PostAlbum(context *gin.Context) {
	var newAlbum src.ALBUM

	if err := context.ShouldBindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo": err})
		return
	}

	_, err := c.db.Exec("INSERT INTO albums (fullname, record_book, birth_date, create_date) VALUES ($1, $2, $3, $4)",
		newAlbum.Fullname, newAlbum.RecordBook, newAlbum.BirthDate, newAlbum.CreateDate)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo": err})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Студент успешно зачислен"})
}

// Метод для удаления альбома по ID
func (c *Controller) DeleteAlbum(context *gin.Context) {
	id := context.Param("id")

	result, err := c.db.Exec("DELETE FROM albums WHERE id = $1", id)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo": err})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Что пошло не так, проверьте логи системы", "errorInfo": err})
		return
	}

	if rowsAffected == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Ошибка"})
		return
	}

	context.IndentedJSON(200, gin.H{"message": "Студент успешно отчислен"})
}
