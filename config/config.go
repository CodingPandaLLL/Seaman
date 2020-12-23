package config

import (
	"encoding/json"
	"os"
)

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
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
	file, err := os.Open("D:\\GoPath\\src\\Seaman\\config.json")
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
