package services

// This is the BUISNESS LOGIC LAYER of the application

import (
	"github.com/pythonakoto/gin_crud/models"
	"github.com/pythonakoto/gin_crud/repositories"
)

type ItemService struct {
	repo *repositories.ItemRepository
}

// NewItemService is a constructor function that creates a new instance of ItemService.
// It takes an ItemRepository as an argument and returns a pointer to an ItemService.
func NewItemService(repo *repositories.ItemRepository) *ItemService {
	// Return a pointer to a new ItemService instance, with the provided repository assigned to its 'repo' field.
	return &ItemService{
		repo: repo, // The repository is passed and stored in the ItemService for later use.
	}
}

// GET all items
func (service *ItemService) GetItems() []models.Item {
	return service.repo.GetAll()
}

// GET an item by ID
func (service *ItemService) GetItem(id int) *models.Item {
	return service.repo.GetByID(id)
}

// CREATE an item
func (service *ItemService) CreateItem(item models.Item) models.Item {
	return service.repo.Create(item)
}

// UPDATE an item
func (service *ItemService) UpdatedItem(id int, item models.Item) *models.Item {
	return service.repo.Update(id, item)
}

// DELETE an item
func (service *ItemService) DeleteItem(id int) bool {
	return service.repo.Delete(id)
}
