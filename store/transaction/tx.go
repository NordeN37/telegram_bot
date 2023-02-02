package transaction

import (
	"database/sql"

	"gorm.io/gorm"
)

type ITransaction interface {
	Begin(opts ...*sql.TxOptions) *gorm.DB
	Commit() *gorm.DB
	Rollback() *gorm.DB
}
