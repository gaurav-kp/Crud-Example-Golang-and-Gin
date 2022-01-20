package dbhelper

import "os"

// sample mysql connection string
// "dbusername:dbuserpassword@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

var DsnMySql = "root:" + os.Getenv("env_mysql_pwd") + "@tcp(127.0.0.1:3310)/test?charset=utf8mb4&parseTime=True&loc=Local"
