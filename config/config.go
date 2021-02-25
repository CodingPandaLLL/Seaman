package config

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	AppScope    string   `json:"app_scope"`
	TenantId    string   `json:"tenant_id"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
}

/**
 * mysql配置
 */
type DataBase struct {
	Drive    string `json:"drive"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

//初始化服务器配置
func InitConfig() *AppConfig {
	var config *AppConfig
	//拼接路径
	path := filepath.Join(GetAppPath(), "config.json")
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}
	//config = &AppConfig{}
	return config
}

//获取项目路径
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	return path[:index]
}
