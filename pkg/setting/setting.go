package setting

import "github.com/spf13/viper"

// 组件：该文件对读取配置的行为进行封装，以便应用程序的使用

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	/* 该方法用于初始化本项目配置的基础属性，即设定配置文件的名称问config，配置类型为yaml，
	并且设置其配置路径为相对路径configs/，以确保在项目目录下能够成功编写组件 */
	vp := viper.New()
	// # viper是允许设置多个配置路径的，这样可以尽可能的尝试解决路径查找问题，也就是说，可以不断的调研AddConfigPath方法！！！
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()

	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
