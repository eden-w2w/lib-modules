package databases

import (
	"github.com/eden-framework/sqlx"
)

var Config = struct {
	DB *sqlx.Database
}{
	DB: &sqlx.Database{},
}
