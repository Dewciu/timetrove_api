package addresses

import "github.com/dewciu/timetrove_api/pkg/common"

type AddressModel struct {
	common.BaseModel
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
}

func (s *AddressModel) TableName() string {
	return "addresses"
}
