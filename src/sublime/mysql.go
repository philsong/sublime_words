package sublime

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(userpost UserPost) bool {
	// 格式有点怪, @tcp 是指网络协议(难道支持udp?), 然后是域名和端口
	db, e := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/weiyan?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		print("ERROR?")
		return false
	}

	defer db.Close()

	// 提醒一句, 运行到这里, 并不代表数据库连接是完全OK的, 因为发送第一条SQL才会校验密码 汗~!
	rows, e := db.Query("insert into users (uname, upwd, uemail) values (?,?,?)",
		userpost.UName, userpost.UPwd, userpost.UEmail)
	if e != nil {
		fmt.Printf("query error!!%v\n", e)
		return false
	}
	if rows == nil {
		print("Rows is nil")
		return false
	}

	fmt.Println("InsertUser Done")
	return true
}

func QueryUser(usersingin UserSignin) (*UserPost, bool) {
	// 格式有点怪, @tcp 是指网络协议(难道支持udp?), 然后是域名和端口
	db, e := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/weiyan?charset=utf8")
	if e != nil { //如果连接出错,e将不是nil的
		print("ERROR?")
		return nil, false
	}

	defer db.Close()

	// 提醒一句, 运行到这里, 并不代表数据库连接是完全OK的, 因为发送第一条SQL才会校验密码 汗~!
	rows, e := db.Query("SELECT `uname`, `upwd`, `uemail` FROM users WHERE upwd = ? AND uemail = ?", usersingin.UPwd, usersingin.UEmail)
	if e != nil {
		fmt.Printf("query error!!%v\n", e)
		return nil, false
	}
	if rows == nil {
		fmt.Printf("Rows is nil")
		return nil, false
	}
	fmt.Println("DB rows.Next")

	user := new(UserPost)
	for rows.Next() { //跟java的ResultSet一样,需要先next读取
		// rows貌似只支持Scan方法 继续汗~! 当然,可以通过GetColumns()来得到字段顺序
		row_err := rows.Scan(&user.UName, &user.UPwd, &user.UEmail)
		if row_err != nil {
			fmt.Printf("Row error!!")
			return nil, false
		}

		fmt.Println(user) // 这里没有判断错误, 呵呵, 一般都不会有错吧
		fmt.Println("QueryUser Done")
		return user, true
	}
	fmt.Println(user) // 这里没有判断错误, 呵呵, 一般都不会有错吧
	fmt.Println("QueryUser failed")
	return nil, false
}
