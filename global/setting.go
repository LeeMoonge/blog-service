package global

// 包全局变量，在pkg/setting中仅读取配置信息是不够的，还需要将配置信息和应用程序关联起来，才能使用它

import "github.com/go/blog-service/pkg/setting"

// 这里对最初预估的从三个区段进行了配置并声明了全局变量，以便在接下来的步骤中将其关联起来，提供给应用程序内部调用
// !!!全局变量的初始化是会随着应用程序的不断演化而不断改变的，也就是说，这里展示的不一定是最终结果
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
