package action

import "time"

// 系统配置
type SettingsConf struct {
	Global     GlobalConf                `yaml:"global"`
	DataSource map[string]DataSourceConf `yaml:"data_sources"`
}

// 全局配置
type GlobalConf struct {
	QueryInterval time.Duration `yaml:"query_interval"` // 查询间隔
	QueryTimeout  time.Duration `yaml:"query_timeout"`  // 查询超时
}

// 数据源
type DataSourceConf struct {
	URL        string                 `yaml:"url"`                  // 数据库连接
	Username   string                 `yaml:"username"`             // 用户名
	Password   string                 `yaml:"password"`             // 用户密码
	Properties map[string]interface{} `yaml:"properties,omitempty"` // 其它数据源配置
}

// 查询参数
type QueryConf struct {
	SQL        string                 `yaml:"sql"`                // 查询SQL
	Params     map[string]interface{} `yaml:"params,omitempty"`   // 查询参数
	Metrics    map[string]interface{} `yaml:"metrics"`            // 暴露指标，非指标字段都将作为标签存在
	DataSource string                 `yaml:"data_source"`        // 查询数据源，数据源使用名称关联
	Interval   time.Duration          `yaml:"interval,omitempty"` // 查询间隔
	Timeout    time.Duration          `yaml:"timeout,omitempty"`  // 查询超时
}
