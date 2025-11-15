package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title"         json:"title"`
	Content   string             `bson:"content"       json:"content"`
	CreatedAt time.Time          `bson:"createdAt"     json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"     json:"updatedAt"`
	ExpiresAt *time.Time         `bson:"expiresAt,omitempty" json:"expiresAt,omitempty"`
}

type StatsResponse struct {
	TotalNotes    int64   `json:"totalNotes" bson:"totalNotes"`
	AvgContentLen float64 `json:"avgContentLen" bson:"avgContentLen"`
	MaxContentLen int64   `json:"maxContentLen" bson:"maxContentLen"`
	MinContentLen int64   `json:"minContentLen" bson:"minContentLen"`
}
