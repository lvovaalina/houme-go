module bitbucket.org/houmeteam/houme-go

// +heroku goVersion go1.17
go 1.17

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	github.com/heroku/x v0.0.31
	github.com/lib/pq v1.10.3 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2
	gorm.io/driver/postgres v1.2.1
	gorm.io/gorm v1.22.2
)

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/compute v1.2.0 // indirect
	cloud.google.com/go/iam v0.1.1 // indirect
	cloud.google.com/go/storage v1.20.0 // indirect
	github.com/appleboy/gin-jwt/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.10.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.8.1 // indirect
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.0.0-20220207234003-57398862261d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/api v0.67.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220207185906-7721543eae58 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
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

replace bitbucket.com/houmeteam/houme-go/converters => ../converters
