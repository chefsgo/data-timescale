package data_postgres

import (
	"github.com/chefsgo/data"
	dp "github.com/chefsgo/data-postgres"
)

func init() {
	driver := dp.Driver()
	data.Register("cockroach", driver)
	data.Register("db", driver)
	data.Register("crdb", driver)
}
