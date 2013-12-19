package sublime

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 结构体成员仅大写开头外界才能访问
type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func Mysqltest() {
	// 格式有点怪, @tcp 是指网络协议(难道支持udp?), 然后是域名和端口
	db, e := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/weiyan?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		print("ERROR?")
		return
	}

	defer db.Close()

	// 提醒一句, 运行到这里, 并不代表数据库连接是完全OK的, 因为发送第一条SQL才会校验密码 汗~!
	rows, e := db.Query("SELECT uid FROM `users` limit 1")
	if e != nil {
		fmt.Printf("query error!!%v\n", e)
		return
	}
	if rows == nil {
		print("Rows is nil")
		return
	}
	fmt.Println("DB rows.Next")
	for rows.Next() { //跟java的ResultSet一样,需要先next读取
		user := new(User)
		// rows貌似只支持Scan方法 继续汗~! 当然,可以通过GetColumns()来得到字段顺序
		row_err := rows.Scan(&user.User, &user.Password, &user.Host)
		if row_err != nil {
			print("Row error!!")
			return
		}
		b, _ := json.Marshal(user)
		fmt.Println(string(b))
	}
	fmt.Println("Done")
}
