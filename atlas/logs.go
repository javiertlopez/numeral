package atlas

import (
	"context"
	"fmt"
	"time"

	"github.com/javiertlopez/numeral"

	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection keeps the collection name
const Collection = "logs"

// log model for mongodb
type log struct {
	ID        string    `bson:"_id"`
	Count     int       `bson:"count,omitempty"`
	ImageKey  string    `bson:"image_key,omitempty"`
	DeviceID  string    `bson:"device_id,omitempty"`
	ShapeID   string    `bson:"shape_id,omitempty"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (r *repository) CreateLog(ctx context.Context, countLog numeral.Log) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	time := time.Now()

	uuid := guuid.New().String()

	insert := &log{
		ID:        uuid,
		ImageKey:  countLog.ImageKey,
		DeviceID:  countLog.DeviceID,
		ShapeID:   countLog.ShapeID,
		CreatedAt: time,
		UpdatedAt: time,
	}

	_, err := collection.InsertOne(ctx, insert)
	if err != nil {
		return numeral.Log{}, err
	}

	return insert.toModel(), nil
}

// GetByID retrieves a log with the ID
func (r *repository) GetByID(ctx context.Context, id string) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	var response log

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&response)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return numeral.Log{}, fmt.Errorf("not found")
		}

		return numeral.Log{}, err
	}

	return response.toModel(), nil
}

func (r *repository) UpdateLog(ctx context.Context, id string, countLog numeral.Log) (numeral.Log, error) {
	collection := r.mongo.Collection(Collection)
	time := time.Now()

	var insert log

	filter := bson.M{"_id": id}

	err := collection.FindOne(ctx, filter).Decode(&insert)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return numeral.Log{}, fmt.Errorf("not found")
		}

		return numeral.Log{}, err
	}

	insert.UpdatedAt = time
	insert.Count = countLog.Count

	update := bson.M{
		"$set": insert,
	}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return numeral.Log{}, err
	}

	return insert.toModel(), nil
}

func (l log) toModel() numeral.Log {
	return numeral.Log{
		ID:        l.ID,
		Count:     l.Count,
		ImageKey:  l.ImageKey,
		DeviceID:  l.DeviceID,
		ShapeID:   l.ShapeID,
		CreatedAt: l.CreatedAt.UTC().String(),
		UpdatedAt: l.UpdatedAt.UTC().String(),
	}
}
