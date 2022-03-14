package mysql_xorm_test

import (
	"path/filepath"
	"fmt"
	"testing"
	"github.com/DavidXia1989/mysql_xorm"
)

func Test_NewClientByFile_1(t *testing.T) {
	mysqlServerConf, _ := filepath.Abs(filepath.Join("testdata", "xmysql.yaml"))
	fmt.Println(mysqlServerConf)
	fmt.Println(123)
	a, err :=mysql_xorm.NewClientByFile(mysqlServerConf)
	fmt.Println(a)
	fmt.Println(err)


}
