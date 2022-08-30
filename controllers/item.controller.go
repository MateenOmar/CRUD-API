package controllers

import (
	"example/CRUD-APIs/models"
	"example/CRUD-APIs/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService services.ItemService
}

func New(itemService services.ItemService) ItemController {
	return ItemController{
		itemService: itemService,
	}
}

func (ic *ItemController) CreateItem(c *gin.Context) {
	var newItem models.Product

	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ic.itemService.CreateItem(&newItem)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func (ic *ItemController) GetItem(c *gin.Context) {
	var itemName string = c.Param("id")
	item, err := ic.itemService.GetItem(&itemName)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (ic *ItemController) GetStock(c *gin.Context) {
	items, err := ic.itemService.GetStock()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// func (ic *ItemController) Purchase(c *gin.Context) {
// 	var item models.Product

// 	if err := c.ShouldBindJSON(&item); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	err := ic.itemService.Purchase(&item)

// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, item)
// }

func (ic *ItemController) Purchase(c *gin.Context) {
	var itemName string = c.Param("id")
	item, err := ic.itemService.GetItem(&itemName)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err = ic.itemService.Purchase(item)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (ic *ItemController) Return(c *gin.Context) {
	c.JSON(200, "")
}

func (ic *ItemController) DeleteItem(c *gin.Context) {
	var itemName string = c.Param("id")
	err := ic.itemService.DeleteItem(&itemName)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func (ic *ItemController) RegisterItemRoutes(rg *gin.RouterGroup) {
	itemroute := rg.Group("/stock")
	itemroute.POST("/create", ic.CreateItem)
	itemroute.GET("/items/:id", ic.GetItem)
	itemroute.GET("/items", ic.GetStock)
	itemroute.PATCH("/purchase/:id", ic.Purchase)
	itemroute.PATCH("/return/:id", ic.Return)
	itemroute.DELETE("/delete/:id", ic.DeleteItem)
}
