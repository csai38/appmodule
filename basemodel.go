package appmodule

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	RemoveAt  time.Time `gorm:"default:null"`
	GroupId   uuid.UUID `gorm:"default:null"`
	CreatorId uuid.UUID
	UpdaterId uuid.UUID
	RemoverId uuid.UUID
}

type BaseModule struct {
	Db         *gorm.DB
	ModuleMap  map[string]interface{}
	ModuleDeps []map[string]string //format dependency key map - module name, value - version
	Version    string              // module version format vX.X.X-YYY
}
