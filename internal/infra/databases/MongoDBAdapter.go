package databases

// import (
// 	"context"
// 	"fmt"
// 	"games/internal/app/data/repository"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type GameRepositoryDB struct {
// 	collection *mongo.Collection
// }

// func MongoDBAdapter(mongoClient *mongo.Client, dbName string, collectionName string) repository.Connection {
// 	collection := mongoClient.Database(dbName).Collection(collectionName)
// 	return &GameRepositoryDB{collection: collection}
// }

// func (db *GameRepositoryDB) Close(ctx context.Context) error {
// 	db.collection.Database().Client().Disconnect(ctx)
// 	return nil
// }
// func (db *GameRepositoryDB) Insert(ctx context.Context, statement ...any) error {
// 	_, err := db.collection.InsertOne(ctx, statement)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (db *GameRepositoryDB) Update(ctx context.Context, statement ...any) error {
// 	_, err := db.collection.UpdateOne(ctx, statement, statement)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (db *GameRepositoryDB) Delete(ctx context.Context, statement ...any) error {
// 	_, err := db.collection.DeleteOne(ctx, statement)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (db *GameRepositoryDB) FindAll(ctx context.Context, statement ...any) (*mongo.Cursor, error) {
// 	res, err := db.collection.Find(ctx, statement)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }
// func (db *GameRepositoryDB) FindOne(ctx context.Context, statement ...any) error {
// 	err := db.collection.FindOne(ctx, statement)
// 	if err != nil {
// 		return fmt.Errorf("error do find by id: %s", err)
// 	}
// 	return nil
// }
