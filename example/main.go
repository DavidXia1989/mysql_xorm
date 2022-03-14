package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"path/filepath"
	"github.com/DavidXia1989/mysql_xorm"
	"time"
)

func main() {
	mysqlServerConf, _ := filepath.Abs(filepath.Join("conf", "xmysql.yaml"))
	a, err :=mysql_xorm.NewClientByFile(mysqlServerConf)
	fmt.Println(a)
	fmt.Println(err)
	fmt.Println( mysql_xorm.GetMysqlClient("h5_admin"))
	sql := "select * from `release`"
	result, _ :=  mysql_xorm.GetMysqlClient("zm_package").Query(sql)
	//release := new(Release)
	fmt.Println(result[0]["created_at"])

	release := make([]Release, 0)
	mysql_xorm.GetMysqlClient("zm_package").Table("release").Find(&release)
	fmt.Println(release)

	data := Release{
		AppId: "123",
		VerStr: "123",
		VerNum: "123",
		Hash: "123",
		Name: "123",
		NameCn: "123",
		NameEn: "123",
		Remark: "123",
		Path: "123",
		CosDomain: "123",
		Status: 1,
	}
	affect ,err := mysql_xorm.GetMysqlClient("zm_package").Table("release").Insert(data)
	fmt.Println(affect)
	fmt.Println(err)

}


type Release struct {
	Id		int		`json:"id" xorm:"pk autoincr"` // 设置后 insert
	AppId	string 	`json:"app_id"`
	VerStr	string	`json:"ver_str"`
	VerNum 	string	`json:"ver_num"`
	Hash 	string		`json:"hash"`
	Name 	string		`json:"name"`
	NameCn 	string		`json:"name_cn"`
	NameEn 	string		`json:"name_en"`
	Remark 	string		`json:"remark"`
	Path 	string		`json:"path"` // 文件路径
	CosDomain	string 	`json:"cos_domain"`// release源码包cos的下载地址
	Status	int			`json:"status"`// 0启用 1禁用
	CreatedAt time.Time `json:"created_at" xorm:"created"` // xorm:"created" 默认设置当前时间
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}