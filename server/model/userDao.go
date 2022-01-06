package model

import (
	"chatroom/common/message"
	"database/sql"
	"fmt"
)

//我们在服务器启动后，就初始化一个userDao实例，
//把它做成全局的变量，在需要和redis操作时，就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao 结构体体
//完成对User 结构体的各种操作.
type UserDao struct {
	Db *sql.DB
}

//思考一下在UserDao 应该提供哪些方法给我们
//1. 根据用户id 返回 一个User实例+err
func (this *UserDao) getUserById(id int) (user *User, err error) {
	sqlStr := `select userId,userPwd,userName from user where userId=?;`
	rowOBj := this.Db.QueryRow(sqlStr, id)
	var u User
	err = rowOBj.Scan(&u.UserId, &u.UserPwd, &u.UserName)
	user = &u
	if err != nil {
		err = ERROR_USER_NOTEXISTS
		return
	}
	return
}

//完成登录的校验 Login
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//
	////先从UserDao 的连接池中取出一根连接
	//conn := this.pool.Get()
	//defer conn.Close()
	user, err = this.getUserById(userId)
	if err != nil {
		return
	}
	//这时证明这个用户是获取到.
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
func (this *UserDao) Register(user *message.User) (err error) {
	_, err = this.getUserById(user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//这时，说明id在redis还没有，则可以完成注册
	//入库
	sqlStr := `insert into user(userId,userPwd,userName)values(?,?,?)`
	_, err = this.Db.Exec(sqlStr, user.UserId, user.UserPwd, user.UserName)
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}
