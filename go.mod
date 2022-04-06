module go_api

go 1.15

require github.com/gorilla/mux v1.8.0

require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence

require internal/web/rest v1.0.0

replace internal/web/rest => ./internal/web/rest
