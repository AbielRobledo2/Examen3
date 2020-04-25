package modelos

import "time"

type Libro struct {
	Id               int       `json:"id" gorm:"primary_key, AUTO_INCREMENT, column:id"`
	Titulo           string    `json:"titulo" gorm:"column:titulo"`
	Autor            string    `json:"autor" gorm:"column:autor"`
	Descripcion      string    `json:"descripcion" gorm:"column:descripcion"`
	Editorial        string    `json:"editorial" gorm:"column:editorial"`
	FechaPublicacion time.Time `json:"fecha_publicacion" gorm:"column:fecha_publicacion"`
}
