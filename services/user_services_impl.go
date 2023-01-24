package services

import (
	"context"
	"errors"

	"example.com/gin-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserservicesImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserservice(usercollection *mongo.Collection, ctx context.Context) Userservices {
	return &UserservicesImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserservicesImpl) Createuser(user *models.User) error {
	_, error := u.usercollection.InsertOne(u.ctx, user)
	return error
}
func (u *UserservicesImpl) Getuser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserservicesImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("documents not found")
	}
	return users, nil
}

func (u *UserservicesImpl) Updateuser(User *models.User) error {
	query := bson.D{bson.E{Key: "user_name", Value: User.Name}}
	filter := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: User.Name}, bson.E{Key: "user_age", Value: User.Age}, bson.E{Key: "user_address", Value: User.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, query, filter)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}
func (u *UserservicesImpl) Deleteuser(name *string) error {

	query := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, query)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil

}
