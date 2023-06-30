package sqlserver

import (
	"github.com/wang-xuemin/gorm"
	"github.com/wang-xuemin/gorm/callbacks"
)

var updateFunc = callbacks.Update(&callbacks.Config{})

func Update(db *gorm.DB) {
	if db.Statement.Schema != nil && db.Statement.Schema.PrioritizedPrimaryField != nil && db.Statement.Schema.PrioritizedPrimaryField.AutoIncrement {
		db.Statement.Omits = append(db.Statement.Omits, db.Statement.Schema.PrioritizedPrimaryField.DBName)
	}

	updateFunc(db)
}
