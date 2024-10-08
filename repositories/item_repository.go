package repositories

import (
	"github.com/pythonakoto/gin_crud/models"
)

var items = []models.Item{
	{ID: 1, Title: "First Item", Body: "This is the first item"},
	{ID: 2, Title: "Second Item", Body: "This is the second item"},
}

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

// GET all items  - this is a 'class method' for ItemRepository
func (repo *ItemRepository) GetAll() []models.Item {
	return items
}

// GET an item by ID -  this is like a Python 'class method' for ItemRepository
func (repo *ItemRepository) GetByID(id int) *models.Item {
	// Loop through the slice of items stored in the repository (in-memory data)
	for _, item := range items {
		// If the item's ID matches the requested ID
		if item.ID == id {
			// Return a pointer to the item
			return &item
		}
	}
	// If no item with the given ID is found, return nil (null) to indicate this
	return nil
}

// CREATE a new item
func (repo *ItemRepository) Create(item models.Item) models.Item {
	// increment the ID by 1 (set it to one more than the current number of items)
	item.ID = len(items) + 1
	// add item object to items slice
	items = append(items, item)
	// return the item you have just created
	return item
}

// UPDATE a item by ID
func (repo *ItemRepository) Update(id int, updatedItem models.Item) *models.Item {
	// Loop through the slice of items stored in the repository (in-memory data)
	for index, item := range items {
		// check if the ID in the struct matches the id we want to update
		if item.ID == id {
			// update the item Title & Body
			items[index].Title = updatedItem.Title
			items[index].Body = updatedItem.Body
			// return the pointer of the updasted item
			return &items[index]
		}
	}
	// If no item with the given ID is found, return nil (null) to indicate this
	return nil
}

// DELETE an item by ID
func (repo *ItemRepository) Delete(id int) bool {
	// Loop through the slice of items stored in the repository (in-memory data)
	for index, item := range items {
		// Check if the ID of the current item matches the ID we want to delete
		if item.ID == id {
			// Remove the item at the found index by slicing out the element
			// append(items[:index], items[index+1:]...) combines all elements
			// before the item and all elements after the item, effectively
			// removing the item at 'index'
			// ... unpacks the slice items[index+1:] and appends its elements to items[:index]
			items = append(items[:index], items[index+1:]...)
			// Return true to indicate the item was successfully deleted
			return true
		}
	}
	// If no item with the given ID is found, return false to indicate failure
	return false
}
