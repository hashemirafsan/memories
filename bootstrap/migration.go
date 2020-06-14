package bootstrap

import (
	"github.com/hashemirafsan/memories/model"
)

// Migration ...
type Migration struct {}

func (m *Migration) Run() {
	db := Database{}
	
	db.Connection().AutoMigrate(&model.ServiceClient{})
	db.Connection().AutoMigrate(&model.ActivityLog{})
}