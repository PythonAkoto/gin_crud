package controllers

// this is the HTTP layer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pythonakoto/gin_crud/models"
	"github.com/pythonakoto/gin_crud/services"
)

// creating a struct to group the togtehr under a single type
type ItemController struct {
	// points to an ItemService pointer struct that lives in the services package
	service *services.ItemService
}

func NewItemController(service *services.ItemService) *ItemController {
	return &ItemController{service: service}
}

// HTML rendering for items
func (control *ItemController) RenderItemsHTML(c *gin.Context) {
	items := control.service.GetItems()

	// render the index.html template with the items data
	c.HTML(http.StatusOK, "index.html", gin.H{
		"items": items, // pass the items list to the template
	})
}

// GET /items (get ALL items)
func (control *ItemController) GetItems(c *gin.Context) {
	items := control.service.GetItems()
	c.JSON(http.StatusOK, items)
}

// GET /items/:id (get items by ID)
func (control *ItemController) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi((c.Param("id")))
	item := control.service.GetItem(id)
	if item != nil {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSONP(http.StatusNotFound, gin.H{"Message": "Message not found"})
	}
}

// POST /items (create a new item)
func (control *ItemController) CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindBodyWithJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := control.service.CreateItem(newItem)
	c.JSON(http.StatusCreated, item)
}

// PUT /items/:id (update an item via ID)
func (control *ItemController) UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	item := control.service.UpdatedItem(id, updatedItem)
	if item != nil {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
	}
}

// DELETE /items/:id (delete item via ID)
func (control *ItemController) DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if control.service.DeleteItem(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
	}
}
