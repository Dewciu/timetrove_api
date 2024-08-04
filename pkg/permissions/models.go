package permissions

import (
	"github.com/dewciu/timetrove_api/pkg/common"
)

type PermissionModel struct {
	common.BaseModel
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
} //@name Permission

func (u *PermissionModel) TableName() string {
	return "permissions"
}

type PermissionGroupModel struct {
	common.BaseModel
	Name        string             `json:"name"`
	Permissions []*PermissionModel `gorm:"many2many:permission_groups;"`
} //@name PermissionGroup

func (u *PermissionGroupModel) TableName() string {
	return "permission_groups"
}
