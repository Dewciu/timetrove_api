package services

import (
	"github.com/dewciu/timetrove_api/pkg/common"
)

type Service struct {
	common.BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
