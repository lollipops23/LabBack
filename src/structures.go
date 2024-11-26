package src

import (
	"time"

	"github.com/go-playground/validator/v10"
)



type ALBUM struct {
	ID         int       `json:"id"`
	Fullname   string    `json:"fullname" validate:"required"`
	RecordBook int64     `json:"record_book" validate:"required"`
	BirthDate  time.Time `json:"birth_date" validate:"required"`
	CreateDate time.Time `json:"create_date" validate:"required"`
}

// var Albums = []ALBUM{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

var validate *validator.Validate

func init(){
	validate=validator.New()
}

func ValidateAlbum(album ALBUM) error {
	return validate.Struct(album)
}
