package numeral

import "context"

// Log struct
type Log struct {
	ID        string `json:"id,omitempty"`
	Count     int    `json:"count,omitempty"`
	ImageKey  string `json:"image_key,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	DeviceID  string `json:"device_id,omitempty"`
	ShapeID   string `json:"shape_id,omitempty"`
}

// Repository interface
type Repository interface {
	CreateLog(ctx context.Context, countLog Log) (Log, error)
}
