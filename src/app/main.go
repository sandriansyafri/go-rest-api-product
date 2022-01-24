package main

import (
	"sans-api-product/handler"
	"sans-api-product/model"
	"sans-api-product/product"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/go_api-project_product?charset=utf8mb4&parseTime=True&loc=Local" // config DB connection
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})                                                //  use gin-gonic

	// LAYERING
	repositoryProduct := product.NewRepository(db)
	serviceProduct := product.NewService(repositoryProduct)
	handlerProduct := handler.NewHandler(serviceProduct)

	// migrate to model / table  to DB
	products := model.Product{}
	db.AutoMigrate(&products)

	r := gin.Default()      // create router
	v1 := r.Group("api/v1") //  group API by version

	//  ROUTES PRODUCT
	v1.GET("/products", handlerProduct.GetProducts)          // GET All Product
	v1.GET("/products/:id", handlerProduct.GetProduct)       // GET Single Product
	v1.POST("/products", handlerProduct.CreateProduct)       // CREATE Product
	v1.PUT("/products/:id", handlerProduct.UpdateProduct)    // UDPATE Product
	v1.DELETE("/products/:id", handlerProduct.DeleteProduct) // DELETE Product

	r.Run(":8000") //  listen and serve on localhost:8000
}
