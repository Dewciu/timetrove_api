package serializers

import (
	"github.com/dewciu/timetrove_api/pkg/database/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddressResponse struct {
	ID     uuid.UUID `json:"-"`
	Street string    `json:"street"`
	City   string    `json:"city"`
	State  string    `json:"state"`
} //@name Address

type AddressSerializer struct {
	C *gin.Context
	models.Address
}

func (s *AddressSerializer) Response() AddressResponse {
	response := AddressResponse{
		ID:     s.ID,
		Street: s.Street,
		City:   s.City,
		State:  s.State,
	}

	return response
}
