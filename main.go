package main

import (
	"net/http"

	modelos "github.com/abielrobledo2/unidad-3/modelos"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func buscarLibros(c *gin.Context) {
	var libros []modelos.Libro
	db.Find(&libros)
	c.JSON(http.StatusOK, libros)
}

func buscarLibro(c *gin.Context) {
	id := c.Param("id")
	var libro modelos.Libro

	err := db.Where("id = ?", id).First(&libro).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"alerta": "El libro no esta registrado"})
		return
	}
	c.JSON(http.StatusOK, libro)
}

func guardarLibro(c *gin.Context) {
	var libro modelos.Libro
	c.BindJSON(&libro)

	err := db.Create(&libro).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"alerta": "No se pudo registrar el libro, intente de nuevo."})
		return
	}

	c.JSON(http.StatusOK, libro)
}

func actualizarLibro(c *gin.Context) {
	id := c.Param("id")
	var libro modelos.Libro

	err := db.Where("id = ?", id).First(&libro).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"alerta": "El libro no esta registrado"})
		return
	}
	var datos modelos.Libro
	c.BindJSON(&datos)

	err = db.Model(&libro).Updates(datos).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"alerta": "No se pudo actualizar el libro"})
		return
	}
	c.JSON(http.StatusOK, libro)
}

func BorrarLibro(c *gin.Context) {
	id := c.Param("id")
	var libro modelos.Libro

	err := db.Where("id = ?", id).First(&libro).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"alerta": "No se pudo borrar el libro"})
		return
	}
	db.Delete(&libro)
	c.JSON(http.StatusOK, libro)
}

func main() {
	// Conexion a la base de datos
	db, _ = gorm.Open("mysql", "root:@/biblioteca?charset=utf8&parseTime=True&loc=Local")

	// Cierra la conexion a la base de datos al finalizar la ejecucion
	defer db.Close()

	// Instancia del servidor
	r := gin.Default()

	// Rutas
	rLibros := r.Group("/libros")
	{
		rLibros.GET("/", buscarLibros)
		rLibros.GET("/:id", buscarLibro)
		rLibros.POST("/", guardarLibro)
		rLibros.PUT("/:id", actualizarLibro)
		rLibros.DELETE("/:id", BorrarLibro)
	}

	// Iniciar servidor en el puerto 3000
	r.Run(":3000")
}
