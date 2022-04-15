package main

type ServerConfig struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

type MySqlConfig struct {
	UserName string  `yaml:"userName"`
	PassWord string  `yaml:"passWord"`
	DataBase string  `yaml:"DataBase"`
	Host     string  `yaml:"host"`
	Port     int     `yaml:"port"`
	TimeOut  float32 `yaml:"timeOut"`
}

// Config 总的
type Config struct {
	Server ServerConfig `yaml:"server"`
	MySql  MySqlConfig  `yaml:"mySql"`
}
