package mysql_xorm

import "time"

type XmsyqlConf struct {
	Key		string	`json:"key"`
	Driver	string	`json:"driver"`
	Host	string	`json:"host"`
	Port	int		`json:"port"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Charset string	`json:"charset"`
	Prefix string	`json:"prefix"`
	Policies int	`json:"policies"` // 从库的负载策略 默认轮询访问 0:轮询访问; 1：随机访问;2：权重随机;3：权重轮询;4：最小连接数
	Policies_weight int	`json:"policies_weight"`// 从库策略 权重比
	MaxIdleConns int	`json:"max_idle_conns"`//连接池的空闲数大小 默认 10个
	MaxOpenConns int	`json:"max_open_conns"`//最大打开连接数 默认设置100个
	ConnMaxLifetime time.Duration	`json:"conn_max_lifetime"`//连接的最大生存时间 默认永久
	Tzlocation string `json:"tzlocation"` // 设置时区
	Master_slave []XmsyqlConf	`json:"master_slave"`
}

func (x *XmsyqlConf) initDefault() {
	if x.Key == "" {
		x.Key = "default"
	}
	if x.Driver == "" {
		x.Driver = "mysql"
	}
	if x.Port == 0 {
		x.Port = 3306
	}
	if x.Charset == "" {
		x.Charset = "utf8"
	}
	if x.MaxIdleConns == 0 {
		x.MaxIdleConns = 10
	}
	if x.MaxOpenConns == 0 {
		x.MaxOpenConns = 100
	}
	if x.Tzlocation == "" {
		x.Tzlocation = "Local"
	}
}