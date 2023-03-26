package appmodule

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	DeletedAt time.Time `gorm:"index;default:null"`
	GroupId   uuid.UUID `gorm:"default:null"`
	CreatorId uuid.UUID
	UpdaterId uuid.UUID
	DeletedId uuid.UUID
}

type SysModules struct {
	BaseModel
	ModuleName string
	Version    string
	ApiMap     string //Json module
}

type MethodParams struct {
	Method      interface{}
	Len         int
	FormHandler bool
}

type BaseModule struct {
	Db         *gorm.DB
	ModuleName string
	ModuleMap  map[string]MethodParams
	ModuleDeps []map[string]string //format dependency key map - module name, value - version
	Version    string              // module version format vX.X.X-YYY
}

func (m BaseModule) ResponseData(resError error, responseBody any, metaData ...interface{}) DataResponse {
	errMsg := ""
	if resError != nil {
		errMsg = resError.Error()
	}
	if responseBody == nil {
		return DataResponse{false, nil, nil, errMsg}
	}
	var meta any = nil
	if metaData != nil {
		meta = metaData
	}
	return DataResponse{true, responseBody, meta, errMsg}
}

func (m BaseModule) ResponseTree(resError error, responseBody any, metaData any) TreeResponse {
	errMsg := ""
	if resError != nil {
		errMsg = resError.Error()
	}
	if responseBody == nil {
		return TreeResponse{false, nil, nil, errMsg}
	}
	var meta any = nil
	if metaData != nil {
		meta = metaData
	}
	return TreeResponse{true, responseBody, meta, errMsg}
}

func (m BaseModule) Api() map[string][]map[string]any {
	api := map[string][]map[string]any{
		m.ModuleName: {},
	}
	var el map[string]any
	for k, v := range m.ModuleMap {
		el = map[string]any{"name": k, "len": v.Len}
		if v.FormHandler {
			el["formHandler"] = true
		}
		api[m.ModuleName] = append(api[m.ModuleName], el)
	}
	return api
}

func (m BaseModule) Route(method string, args ...interface{}) (any, error) {
	if val, ok := m.ModuleMap[method]; ok {
		if len(args) == 0 {
			return val.Method.(func() []map[string]string)(), nil
		} else {
			return val.Method.(func(string) []map[string]string)(args[0].(string)), nil
		}
	} else {
		return nil, fmt.Errorf("Method %s not defined in module  %s", method, m.ModuleName)
	}
}
