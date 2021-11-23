package mysql_xorm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"fmt"
	"errors"
	"io/ioutil"
	"xorm.io/core"
	"time"
)

var EngineGroup = make(map[string]*xorm.EngineGroup)

// 通过文件地址，读取文件初始化
func NewClientByFile(confFile string) (*xorm.EngineGroup, error) {
	var c *xorm.EngineGroup

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return c, errors.New("mysql 配置文件读取失败：" + err.Error())
	}
	var conf []XmsyqlConf
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return c, errors.New("mysql 配置文件解析失败：" + err.Error())
	}
	return NewClients(conf)
}

// 通过struct 配置文件
func NewClients(conf []XmsyqlConf) (*xorm.EngineGroup, error) {
	var c *xorm.EngineGroup
	for k := range conf {
		var err error
		EngineGroup[conf[k].Key], err = NewClient(conf[k])
		if err != nil && EngineGroup[conf[k].Key] == nil {
			return c, errors.New("NewEngineGroup is error：" + err.Error())
		}
		if conf[k].Key == "default" {
			c = EngineGroup[conf[k].Key]
		}
	}
	return c, nil
}

// 根据XmsyqlConf的struce初始化一个引擎
func NewClient(c XmsyqlConf) (*xorm.EngineGroup, error) {
	if EngineGroup[c.Key] != nil {
		return EngineGroup[c.Key], errors.New("连接名已存在")
	}
	var master *xorm.Engine
	slaves := []*xorm.Engine{}
	var policies_weight []int
	policise := xorm.RoundRobinPolicy()
	// 检查是否配置主从
	if len(c.Master_slave) > 0 {
		for i := range c.Master_slave {
			temp, err := setEngine(c.Master_slave[i])
			if err != nil {
				return nil, err
			}
			if c.Master_slave[i].Key == "master" {
				master = temp
			} else {
				policies_weight = append(policies_weight, c.Master_slave[i].Policies_weight)
				slaves = append(slaves, temp)
			}
		}
		if c.Policies != 0 {
			policise = mappingPolicies(c.Policies, policies_weight)
		}
	} else {
		temp, err := setEngine(c)
		if err != nil {
			return nil, err
		}
		master = temp
		slaves = append(slaves, temp)
	}
	engineGroup, err := xorm.NewEngineGroup(master, slaves, policise)
	if err != nil {
		return nil, errors.New("NewEngineGroup is error：" + err.Error())
	}
	return engineGroup, nil
}

// 获取客户端
func GetMysqlClient(key string) (*xorm.EngineGroup) {
	if EngineGroup[key] != nil {
		return EngineGroup[key]
	}
	return nil
}


// 设置引擎
func setEngine(conf XmsyqlConf) (*xorm.Engine, error) {
	conf.initDefault()
	engine, err := xorm.NewEngine(conf.Driver,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
			conf.Charset))
	if err != nil {
		return nil, errors.New("engine is loading error：" + err.Error())
	}
	// 设置表前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, conf.Prefix)
	engine.SetTableMapper(tbMapper)
	// 连接池的空闲数大小
	engine.SetMaxIdleConns(conf.MaxIdleConns)
	// 最大打开连接数
	engine.SetMaxOpenConns(conf.MaxOpenConns)
	// 连接的最大生存时间
	engine.SetConnMaxLifetime(conf.ConnMaxLifetime)
	// 默认设置时区
	engine.TZLocation, _ = time.LoadLocation(conf.Tzlocation)
	err = engine.Ping()
	if (err != nil){
		return nil, errors.New("engine is ping error：" + err.Error())
	}
	return engine, nil
}


// 设置从库访问策略
func mappingPolicies(policies int,weights []int) xorm.GroupPolicyHandler {
	if policies == 1 {
		// 随机访问负载策略
		return xorm.RandomPolicy()
	}else if(policies == 2){
		// 权重随机访问
		return xorm.WeightRandomPolicy(weights)
	}else if(policies == 3){
		// 权重轮询
		return xorm.WeightRoundRobinPolicy(weights)
	}else if(policies == 4){
		// 最小连接数
		return  xorm.LeastConnPolicy()
	}
	// 轮询访问负载策略
	return xorm.RoundRobinPolicy()
}