package numeral

import "context"

// Log struct
type Log struct {
	ID        string `json:"id,omitempty"`
	Count     int    `json:"count,omitempty"`
	ImageKey  string `json:"image_key,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	ShapeID   string `json:"shape_id,omitempty"`
}

// Repository interface
type Repository interface {
	CreateLog(ctx context.Context, countLog Log) (Log, error)
	GetByID(ctx context.Context, id string) (Log, error)
	UpdateLog(ctx context.Context, id string, countLog Log) (Log, error)
}
