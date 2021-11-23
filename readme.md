## Installation mysql_xorm

解决mysql的读写分离，多个链接的管理问题
Env：

golang >= 1.13.0

Install:

```
// 设置环境变量使得go支持私有库
GOPRIVATE=code.zm.shzhanmeng.com
// 安装
go get -u code.zm.shzhanmeng.com/go-common/mysql_xorm
```

Import:

```go
import "code.zm.shzhanmeng.com/go-common/mysql_xorm"
```