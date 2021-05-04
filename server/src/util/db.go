package util

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var configPath string = os.Getenv("CONFIG_PATH")

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

//数据库的配置
const driverName = "mysql"

//DB数据库连接池
var DB *sql.DB

func InitDB() {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	config := &Config{}
	err = json.Unmarshal(buf, config)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=uft8"
	//注意：要想解析time.Time类型，必须要设置parseTime=True
	path := strings.Join([]string{config.Username, ":", config.Password, "@tcp(", config.IP, ":", config.Port, ")/", config.DBName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	//打开数据库，前者是驱动名，所以要导入:_"github.com/go-sql-driver/mysql"
	DB, _ = sql.Open(driverName, path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		log.Panic(err)
	}
	log.Println("database connect success")
}
