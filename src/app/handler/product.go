package handler

import (
	"net/http"
	"sans-api-product/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

// create private model for handler and ref Interface Service
type handler struct {
	service product.Service
}

// get private model handler
func NewHandler(s product.Service) *handler {
	return &handler{service: s}
}

// GET ALL
func (h *handler) GetProducts(c *gin.Context) {
	s := c.Query("s")

	// is Query Params
	if s != "" {
		products, err := h.service.FindLIKE(s)

		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"ok":      false,
				"status":  http.StatusInternalServerError,
				"message": "FAILED TO FETCH DATA",
				"errors":  err,
			})

			return
		}

		c.JSON(200, gin.H{
			"ok":      true,
			"status":  http.StatusOK,
			"message": "FETCH DATA SUCCESS",
			"data":    products,
		})

		return

	}

	// isNot Query Params
	result, err := h.service.FindAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"status":  http.StatusInternalServerError,
			"message": "FAILED TO FETCH DATA",
			"errors":  err,
		})

		return
	}

	c.JSON(200, gin.H{
		"ok":      true,
		"status":  http.StatusOK,
		"message": "FETCH DATA SUCCESS",
		"data":    result,
	})
}

// GET
func (h *handler) GetProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	products, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"status":  http.StatusInternalServerError,
			"message": "NO FOUND PRODUCT",
			"data":    err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"status":  http.StatusOK,
		"message": "FETCH SUCCESS",
		"data":    products,
	})
}

// CREATE
func (h *handler) CreateProduct(c *gin.Context) {

	productRequest := product.ProductRequest{}
	c.ShouldBindJSON(&productRequest)

	result, err := h.service.Create(productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"status":  http.StatusInternalServerError,
			"message": "FAILED TO CREATE DATA",
			"errors":  err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"status":  http.StatusOK,
		"message": "CREATED DATA SUCCESS",
		"data":    result,
	})
}

// UPDATE
func (h *handler) UpdateProduct(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	productRequest := product.ProductRequest{}

	c.ShouldBindJSON(&productRequest)
	product, err := h.service.Update(id, productRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":     false,
			"status": http.StatusInternalServerError,
			"err":    err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"status":  http.StatusOK,
		"message": "UPDATED DATA SUCCESS",
		"data":    product,
	})

}

// DELETE
func (h *handler) DeleteProduct(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	deletedProduct, err := h.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"status":  http.StatusInternalServerError,
			"message": "NO FOUND DATA",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"status":  http.StatusOK,
		"message": "DELETE DATA SUCCESS",
		"data":    deletedProduct,
	})

}
