package iniconfig

import (
	"testing"
	"fmt"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig `ini:"mysql"`
}

type ServerConfig struct {
	Ip string `ini:"ip"`
	Port int `ini:"port"`
}

type MysqlConfig struct {
	Username string `ini:"username"`
	Passwd string `ini:"passwd"`
	Database string `ini:"database"`
	Host string `ini:"host"`
	Port int `ini:"port"`
}

func TestIniFileConfig(t *testing.T) {
	filename := "./config.ini"
	/*UnMarshal the config file*/
	var conf Config
	err := UnMarshalFileConfig(filename, &conf)
	if err != nil {
		t.Errorf("UnMarshal err=%v\n", err)
		return
	}
	
	fmt.Printf("UnMarshal success, %#v\n", conf)

	/*Marshal the config file*/
	err = MarshalFileConfig("./config2.ini", conf)
	if err != nil {
		t.Errorf("Marshal err=%v\n", err)
		return
	}	
	fmt.Printf("Marshal success\n");
	
}



