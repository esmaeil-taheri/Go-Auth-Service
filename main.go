package main

import (
	"time"
	"fmt"

	"gameauth_v2/config"
	"gameauth_v2/repository/mysql"
	"gameauth_v2/delivery/httpserver"
	"gameauth_v2/repository/migrator"
	"gameauth_v2/service/userservice"
	"gameauth_v2/service/authservice"
	"gameauth_v2/validator/uservalidator"
)

const (
	JwtSignKey = "jwt_secret"
	AccessTokenSubject = "ac"
	RefreshTokenSubject = "rt"
	AccessTokenExpireDuration = time.Hour * 24
	RefereshTokenExpireDuration = time.Hour * 7 * 24
)



func main() {

	// TODO - read config path from command line
	cfg2 := config.Load("config.yml")
	fmt.Printf("cfg2: %+v\n", cfg2)

	cfg := config.Config {
		HTTPServer: config.HTTPServer{Port: 8080},
		Auth: authservice.Config{
			SignKey: JwtSignKey,
			AccessExpirationTime: AccessTokenExpireDuration,
			RefreshExpirationTime: RefereshTokenExpireDuration,
			AccessSubject: AccessTokenSubject,
			RefreshSubject: RefreshTokenSubject,
		},
		Mysql: mysql.Config{
			Username: "gameapp_user",
			Password: "gameapp_pass",
			Port: 3306,
			Host: "localhost",
			DBName: "gameapp",
		},
	}

	// TODO - add command for migrations
	mgr := migrator.New(cfg.Mysql)
	mgr.Up()

	authSvc, userSvc, userValidator := setupServices(cfg)

	server := httpserver.New(cfg, authSvc, userSvc, userValidator)

	fmt.Println("start echo server")
	server.Serve()

}

func setupServices(cfg config.Config) (authservice.Service, userservice.Service, uservalidator.Validator) {
	authSvc := authservice.New(cfg.Auth)

	MysqlRepo := mysql.New(cfg.Mysql)
	userSvc := userservice.New(authSvc, MysqlRepo)

	uV := uservalidator.New(MysqlRepo)

	return authSvc, userSvc, uV
}
