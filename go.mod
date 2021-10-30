module bitbucket.org/houmeteam/houme-go

go 1.16

require (
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/heroku/x v0.0.31
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace bitbucket.com/houmeteam/houme-go/forge => ../forge

replace bitbucket.com/houmeteam/houme-go/models => ../models

replace bitbucket.com/houmeteam/houme-go/configs => ../configs

replace bitbucket.com/houmeteam/houme-go/database => ../database

replace bitbucket.com/houmeteam/houme-go/dtos => ../dtos

replace bitbucket.com/houmeteam/houme-go/helpers => ../helpers

replace bitbucket.com/houmeteam/houme-go/langs => ../langs

replace bitbucket.com/houmeteam/houme-go/repositories => ../repositories

replace bitbucket.com/houmeteam/houme-go/services => ../services
