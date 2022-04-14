module go_api

go 1.15

require github.com/gorilla/mux v1.8.0

require internal/entities v1.0.0

require internal/persistence/daobolt v1.0.0

replace internal/persistence/daobolt => ./internal/persistence/daobolt

require internal/persistence/daomemory v1.0.0

replace internal/persistence/daomemory => ./internal/persistence/daomemory

replace internal/entities => ./internal/entities

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence

require internal/persistence/interfaces v1.0.0

replace internal/persistence/interfaces => ./internal/persistence/interfaces

require (
	go.mongodb.org/mongo-driver v1.9.0 // indirect
	internal/web/rest v1.0.0
)

replace internal/web/rest => ./internal/web/rest

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	internal/persistence/mongodb v1.0.0
)

replace internal/persistence/mongodb => ./internal/persistence/mongodb

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
	github.com/swaggo/gin-swagger v1.1.0 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	github.com/urfave/cli v1.20.0 // indirect
	github.com/urfave/cli/v2 v2.4.0 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/tools v0.1.10 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	internal/persistence/bolt v1.0.0
)

replace internal/persistence/bolt => ./internal/persistence/bolt
