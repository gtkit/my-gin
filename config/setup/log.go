package setup

import (
	"github.com/spf13/viper"

	"ydsd_gin/config"
)

func init() {
	config.Add("log", func() map[string]interface{} {
		return map[string]interface{}{
			"division": viper.GetString("log.division"),
			// 是否打印 执行的sql
			"sql": viper.GetInt("log.sql"),
			// 是否压缩，压缩日志不方便查看，我们设置为 false（压缩可节省空间）
			"compress":      viper.GetInt("log.compress"),
			"consolestdout": viper.GetInt("log.consolestdout"),
			"filestdout":    viper.GetInt("log.filestdout"),
			// 日志级别，必须是以下这些选项：
			// "debug" —— 信息量大，一般调试时打开。系统模块详细运行的日志，例如 HTTP 请求、数据库请求、发送邮件、发送短信
			// "info" —— 业务级别的运行日志，如用户登录、用户退出、订单撤销。
			// "warn" —— 感兴趣、需要引起关注的信息。 例如，调试时候打印调试信息（命令行输出会有高亮）。
			// "error" —— 记录错误信息。Panic 或者 Error。如数据库连接错误、HTTP 端口被占用等。一般生产环境使用的等级。
			// 以上级别从低到高，level 值设置的级别越高，记录到日志的信息就越少
			// 开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"
			"level":     viper.GetString("log.level"),
			"localtime": viper.GetInt("log.localtime"),
			// 最多保存多少天，7 表示一周前的日志会被删除，0 表示不删
			"maxage": viper.GetInt("log.maxage"),
			// 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
			"maxbackups": viper.GetInt("log.maxbackups"),
			// 每个日志文件保存的最大尺寸 单位：M
			"maxsize": viper.GetInt("log.maxsize"),
			"path":    viper.GetString("log.path"),
			"errpath": viper.GetString("log.errpath"),
		}
	})
}
