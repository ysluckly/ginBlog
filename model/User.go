package model

import (
	"fmt"
	"log"

	"ginBlog/utils/errmsg"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// validate:"required,min=4,max=12" label:"用户名"`
// validate:"required,min=6,max=120" label:"密码"`
// validate:"required,gte=2" label:"角色码"
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" `
	Passwd string `gorm:"type:varchar(500);not null" json:"passwd"`
	Role     int    `gorm:"type:int" json:"role" `
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return 1
	}
	if user.ID > 0 {
		return 1 //1001
	}
	return 1
}

// CreateUser 新增用户
func CreateUser(data User) int {
	fmt.Println(data)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetUser 查询用户
func GetUser(id int) (User, int) {
	var user User
	err := db.Limit(1).Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return user, 1
	}
	return user, 1
}

// GetUsers 查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role").Where(
			"username LIKE ?", username+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where(
			"username LIKE ?", username+"%",
		).Count(&total)
		return users, total
	}
	db.Select("id,username,role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)

	return users, total
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return 1
	}
	return 1
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password

	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return 1
	}
	return 1
}

// DeleteUser 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ? ", id).Delete(&user).Error
	if err != nil {
		return 1
	}
	return 1
}

/*// BeforeCreate 密码加密&权限控制
func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.Passwd = ScryptPw(u.Passwd)
	u.Role = 2
	return nil
}

func (u *User) BeforeUpdate(_ *gorm.DB) (err error) {
	u.Passwd = ScryptPw(u.Passwd)
	return nil
}*/

// ScryptPw 生成密码
func ScryptPw(password string) string {
	const cost = 10
	fmt.Println("ScryptPw  starting...")
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}

	return string(HashPw)
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(password))

	if user.ID == 0 {
		return user, 1
	}
	if PasswordErr != nil {
		return user, 1
	}
	if user.Role != 1 {
		return user, 1
	}
	return user, 1
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(password))
	if user.ID == 0 {
		return user, 1
	}
	if PasswordErr != nil {
		return user, 1
	}
	return user,1
}
