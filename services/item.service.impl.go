package services

import (
	"context"
	"errors"
	"example/CRUD-APIs/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemServiceImpl struct {
	itemCollection *mongo.Collection
	c              context.Context
}

func NewItemService(itemCollection *mongo.Collection, c context.Context) ItemService {
	return &ItemServiceImpl{
		itemCollection: itemCollection,
		c:              c,
	}
}

func (i *ItemServiceImpl) CreateItem(item *models.Product) error {
	_, err := i.itemCollection.InsertOne(i.c, item)
	return err
}

func (i *ItemServiceImpl) GetItem(id *string) (*models.Product, error) {
	var item *models.Product
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := i.itemCollection.FindOne(i.c, query).Decode(&item)
	return item, err
}

func (i *ItemServiceImpl) GetStock() ([]*models.Product, error) {
	var items []*models.Product
	cursor, err := i.itemCollection.Find(i.c, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(i.c) {
		var item models.Product
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(i.c)
	if len(items) == 0 {
		return nil, errors.New("no items found")
	}
	return items, nil
}

func (i *ItemServiceImpl) Purchase(item *models.Product) error {
	filter := bson.D{bson.E{Key: "id", Value: item.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "id", Value: item.ID}, bson.E{Key: "name", Value: item.Name}, bson.E{Key: "quantity", Value: item.Quantity}}}}
	result, _ := i.itemCollection.UpdateOne(i.c, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no match found")
	}
	return nil
}

func (i *ItemServiceImpl) Return(item *models.Product) error {
	return nil
}

func (i *ItemServiceImpl) DeleteItem(id *string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, _ := i.itemCollection.DeleteOne(i.c, filter)
	if result.DeletedCount != 1 {
		return errors.New("no match found")
	}
	return nil
}
