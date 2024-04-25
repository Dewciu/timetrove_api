package address

import (
	"github.com/dewciu/timetrove_api/pkg/common"
)

type Address struct {
	common.BaseModel
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}
