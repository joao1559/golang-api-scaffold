package repositories

import (
	"context"
	"fmt"

	"github.com/joao1559/golang-api-scaffold/interfaces"
	"github.com/joao1559/golang-api-scaffold/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoNoteRepository struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

// NewMongoNoteRepository ...
func NewMongoNoteRepository(DBConn *mongo.Client, Ctx context.Context) interfaces.NoteRepository {
	return &mongoNoteRepository{
		Collection: DBConn.Database("scaffold").Collection("note"),
		Ctx:        Ctx,
	}
}

func (m *mongoNoteRepository) Get() ([]*models.Note, error) {
	var notes []*models.Note

	cursor, err := m.Collection.Find(m.Ctx, bson.D{{}})
	if err != nil {
		fmt.Println(err)
		return notes, err
	}

	defer cursor.Close(m.Ctx)
	for cursor.Next(m.Ctx) {
		var t models.Note
		err := cursor.Decode(&t)
		if err != nil {
			return notes, err
		}

		notes = append(notes, &t)
	}

	if err := cursor.Err(); err != nil {
		return notes, err
	}

	return notes, nil
}

func (m *mongoNoteRepository) Insert(note models.Note) error {
	_, err := m.Collection.InsertOne(m.Ctx, note)
	if err != nil {
		return err
	}

	return err
}
