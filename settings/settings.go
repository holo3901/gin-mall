package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init2() (err error) {
	//1.相对路径
	viper.SetConfigFile("./config.yaml")
	//2.绝对路径viper.SetConfigFile("golang/go-project/src/goweb开发进阶/3.goweb开发常用组建/goweb开发脚本架/config.yaml")

	//3.viper.SetConfigName("config") //指定配置文件名称(不需要带后缀)
	//viper.SetConfigType("yaml")   //指定配置文件类型
	//viper.AddConfigPath(".")   //指定查找配置文件的路径(这里使用相对路径)

	//4.os获取文件路径
	err = viper.ReadInConfig() //读取配置信息
	if err != nil {
		//读取配置信息失败
		fmt.Printf("viper.ReadConfig() failed,err:%v\n", err)
		return
		//panic(fmt.Errorf("Fatal error config file:%s \n", err))
	}
	viper.WatchConfig() //配置文件实时监控,当配置文件发生变化之后，会实时更新
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarsjal failed,err:%v\n", err)
		}
	})
	return
}
