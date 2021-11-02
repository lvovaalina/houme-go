module bitbucket.org/houmeteam/houme-go

go 1.16

require (
	github.com/alexbrainman/sspi v0.0.0-20180613141037-e580b900e9f5 // indirect
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/heroku/x v0.0.31
	github.com/jcmturner/gokrb5/v8 v8.2.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/lib/pq v1.10.3 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/jcmturner/aescts.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/dnsutils.v1 v1.0.1 // indirect
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0 // indirect
	gopkg.in/jcmturner/rpc.v1 v1.1.0 // indirect
	gorm.io/driver/postgres v1.2.1 // indirect
	gorm.io/gorm v1.22.2 // indirect
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
