package config

import (
	"api/controller"
	"api/system"
)

/**
ip
port
home address
 */
const (
	IP_ADDRESS = "127.0.0.1"
	IP_PORT    = "8080"

	DB_USER = "root"
	DB_PASSWORD="123456"
	DB_PORT="3306"
	DB_NAME="test"

	HOME_URI   = "/home"


)

/**
Route table
 */
var RouteTable = system.Router{
	"HOME": system.Cfg{&controller.HomeController{}, "GET", system.MethodMap{ "Index" : "" , }},
}
