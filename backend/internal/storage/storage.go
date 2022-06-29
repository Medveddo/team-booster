package storage

import (
	"context"
	"log"

	"github.com/Medveddo/team-booster/backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ISkillStorage interface {
	InsertSkill(skill *models.Skill) error
	GetSkills() ([]models.Skill, error)
}

func NewSkillStorage() *skillStorage {
	ctx, f := context.WithCancel(context.TODO())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:mongo@localhost:27017"))
	if err != nil {
		panic(err)
	}
	disconnect_func := func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection := client.Database("tmbster").Collection("skills")

	return &skillStorage{
		collection:      collection,
		context:         &ctx,
		disconnect_func: disconnect_func,
		cancel_ctx_func: f,
	}
}

type skillStorage struct {
	collection      *mongo.Collection
	context         *context.Context
	disconnect_func func()
	cancel_ctx_func func()
}

func (ss *skillStorage) InsertSkill(skill *models.Skill) error {
	_, err := ss.collection.InsertOne(*ss.context, *skill)
	if err != nil {
		panic(err)
	}
	return nil
}

func (ss *skillStorage) GetSkills() ([]models.Skill, error) {
	cur, err := ss.collection.Find(*ss.context, bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cur.Close(*ss.context)
	var skills []models.Skill = make([]models.Skill, 0)
	for cur.Next(*ss.context) {
		var result models.Skill
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		skills = append(skills, result)
	}
	return skills, nil
}

func (ss *skillStorage) Disconnect() {
	ss.cancel_ctx_func()
	ss.disconnect_func()
}
