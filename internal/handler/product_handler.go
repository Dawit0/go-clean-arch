package handler

import (
	"net/http"
	"strconv"

	"github.com/Dawit0/cleanArch/internal/domain"
	"github.com/Dawit0/cleanArch/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandeler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(roter *gin.Engine, usecase usecase.ProductUsecase) {
	handeler := &ProductHandeler{usecase}

	roter.POST("/product/", handeler.Create)
	roter.GET("/products", handeler.GetAll)
	roter.GET("/product/:id", handeler.GetById)
	roter.PUT("/product/:id", handeler.UpdateProduct)
	roter.DELETE("/product/:id", handeler.DeleteProduct)
}

func (p *ProductHandeler) DeleteProduct(c *gin.Context) {
	panic("unimplemented")
}

func (p *ProductHandeler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := p.usecase.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (p *ProductHandeler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := p.usecase.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (p *ProductHandeler) Create(c *gin.Context) {
	var product domain.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.usecase.Create(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)

}

func (p *ProductHandeler) GetAll(c *gin.Context) {
	products, err := p.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
