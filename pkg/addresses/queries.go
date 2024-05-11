package addresses

import (
	"fmt"

	"github.com/dewciu/timetrove_api/pkg/common"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateAddressQuery(address AddressModel) error {
	r := common.DB.Create(&address)
	if r.Error != nil {
		err := r.Error.(*pgconn.PgError)

		if err.Code == "23505" {
			fmt.Println(err.Detail)
			column := common.GetColumnFromUniqueErrorDetails(err.Detail)
			return &common.AlreadyExistsError{Column: column}
		}
	}

	return r.Error
}
