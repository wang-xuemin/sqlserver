# GORM SQL Server Driver

## USAGE

```go
import (
  "github.com/wang-xuemin/sqlserver"
  "github.com/wang-xuemin/gorm"
)

// github.com/microsoft/go-mssqldb
dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
```

Checkout [https://gorm.io](https://gorm.io) for details.
