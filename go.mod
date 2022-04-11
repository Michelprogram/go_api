module go_api

go 1.15

require github.com/gorilla/mux v1.8.0

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence

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

require internal/persistence/bolt v1.0.0

replace internal/persistence/bolt => ./internal/persistence/bolt
