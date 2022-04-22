package data_postgres

import (
	"github.com/chefsgo/chef"
	dp "github.com/chefsgo/data-postgres"
)

func init() {
	driver := dp.Driver()
	chef.Register("cockroach", driver)
	chef.Register("db", driver)
	chef.Register("crdb", driver)
}
