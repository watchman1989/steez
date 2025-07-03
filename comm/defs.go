package comm

import (
	"os"
	"os/signal"
	"syscall"
	"github.com/watchman1989/utils/logger"
	"gopkg.in/yaml.v2"
	"net"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	
)

const (
	defaultConfigFile = "./conf/config.yaml"
)

var (
	Sig  = make(chan os.Signal, 1)
	Quit = make(chan struct{})
	GContext GlobalContext
)

func Init() {
	signal.Notify(Sig, syscall.SIGINT, syscall.SIGTERM)
	//init config file
	initConfig(defaultConfigFile)
	//
	host := ""
	//init log
	lg := logger.NewDefaultLogger(logger.InitArgs{
		SrvName: GContext.GConfig.Server.Name,
		SrvHost:host,
	})
	GContext.Logger = lg
	//init mysql
	db, err := InitMysqlClient(GContext.GConfig.Mysql)
	if err != nil {
		panic(err)
	}
	GContext.Mysql = db

	GContext.Logger.Infof("init over")
}



func initConfig(configPath string) {
	_, err := os.Stat(configPath)
	if err != nil {
		panic(err)
	}
	configContent, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configContent, &GContext.GConfig)
	if err != nil {
		panic(err)
	}
}

type GlobalContext struct {
	GConfig GlobalConfig
	Logger *logger.Logger
	Mysql *gorm.DB
}

type GlobalConfig struct {
	Server ServerConfig `yaml:"server"`
	Mysql MysqlConfig `yaml:"mysql"`
}

type ServerConfig struct {
	Name string `yaml:"name"`
	Port int `yaml:"port"`
}

type MysqlConfig struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname string `yaml:"dbname"`
}


func GetLocalIp(name string) (ipAddr string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		//fmt.Printf("fail to get net interface: %s\n", err.Error())
		return
	}

	for _, iface := range interfaces {
		addrs, err1 := iface.Addrs()
		if err1 != nil {
			continue
		}
		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && name == iface.Name {
				ipAddr = ipNet.IP.String()
				return
			}
		}
	}
	if ipAddr == "" {
		err = fmt.Errorf("get ip addr failed")
	}
	return
}



const (
	dsnFmt = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10s"
)

func InitMysqlClient(config MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(dsnFmt, config.User, config.Password, config.Host, config.Port, config.Dbname)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("####init mysql error: %s\n", err.Error())
		return nil, err
	}
	return DB, nil
}