module b.carriage.fun/migrate

go 1.20

require (
	b.carriage.fun/datamodel v0.0.1
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.3
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
)

replace b.carriage.fun/datamodel => ../datamodel
