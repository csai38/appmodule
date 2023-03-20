package appmodule

import (
	"github.com/google/uuid"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	RemoveAt  time.Time `gorm:"default:null"`
	GroupId   uuid.UUID `gorm:"default:0"`
	CreatorId uuid.UUID
	UpdaterId uuid.UUID
	RemoverId uuid.UUID
}

type BaseModule struct{
	db *gorm.db
	moduleMap map[string] interface{}
}
