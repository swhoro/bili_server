module b.carriage.fun/server

go 1.19

require (
	b.carriage.fun/datamodel v0.0.1
	github.com/gofiber/contrib/jwt v1.0.4
	github.com/gofiber/fiber/v2 v2.48.0
	github.com/golang-jwt/jwt/v5 v5.0.0
	go.uber.org/zap v1.25.0
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.3
)

replace b.carriage.fun/datamodel => ../datamodel

require (
	github.com/MicahParks/keyfunc/v2 v2.1.0 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/goccy/go-json v0.10.2
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.48.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
)
