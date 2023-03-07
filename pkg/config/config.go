package config

import (
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper" // 自定义包名，避免与内置 viper 实例冲突
)

// viper库实例
var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFunc map[string]ConfigFunc

func init() {
	// 1. 初始化 Viper 库
	viper = viperLib.New()

	//2. 配置类型
	viper.SetConfigType("env")
	//3. 环境变量配置文件查找的路径， 相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀， 用以区分Go的系统环境变量
	viper.SetEnvPrefix("appenv")
	//5. 读取环境变量
	viper.AutomaticEnv()

	ConfigFunc = make(map[string]ConfigFunc)

}

// InitConfig 初始化配置信息， 完成对环境变量 以及 config信息加载
func InitConfig(env string) {
	//1.加载环境变量
	loadEnv(env)

	//2.注册配置信息
	loadConfig()
}

func loadConfig() {

	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}

}

func loadEnv(envSuffix string) {
	//默认加载 .env 文件， 如果有传参 --env=name 的话， 加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filepath := ".env." + envSuffix
		if _, err := os.Stat(filepath); err == nil {
			//如 .env.testing 或 .env.stage
			envPath = filepath
		}
	}

	//加载 env
	viper.SetConfigName(envPath)
	if err := viper.ReadConfig(); err != nil {
		panic(err)
	}

	//监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}

// env 读取环境变量， 支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}

	return internalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFunc[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取， 如： app.name
// 第二个参数允许传参默认值
func Get(paht string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	//config 或者环境变量不存在的情况
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}

		return nil
	}
	return viper.Get(path)

}

// GetInt 获取Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}
