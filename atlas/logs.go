package atlas

import (
	"context"
	"time"

	"github.com/javiertlopez/numeral"

	guuid "github.com/google/uuid"
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

func (l log) toModel() numeral.Log {
	return numeral.Log{
		ID:        l.ID,
		Count:     l.Count,
		ImageKey:  l.ImageKey,
		DeviceID:  l.DeviceID,
		ShapeID:   l.ShapeID,
		Timestamp: l.CreatedAt.UTC().Unix(),
	}
}
