package model

import (
	"database/sql"
	"fmt"
	"log"
)

//创建表
func CreateTable(DB *sql.DB) {
	sql := `CREATE TABLE IF NOT EXISTS users(
   id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
   username VARCHAR(64),
   password VARCHAR(64),
   status INT(4),
   createtime INT(10)
   ); `

	if _, err := DB.Exec(sql); err != nil {
		fmt.Println("create table failed:", err)
		return
	}
	fmt.Println("create table successd")
}

func InsertData(db *sql.DB) {
	rows, err := db.Query(`INSERT INTO user (id, name, age) VALUES (2, "ys", 210)`)
	defer rows.Close()

	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}

	var result int
	rows.Scan(&result)
	log.Printf("insert result %v\n", result)
}

func SelectData(db *sql.DB) {
	var id int
	var name string
	var age int
	rows, err := db.Query(`SELECT * From user where id = 2`)
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&id, &name,&age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println( id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}


//更新数据

func UpdateData(DB *sql.DB) {
	result, err := DB.Exec("UPDATE users set password=? where id=?", "111111", 3)

	if err != nil {
		fmt.Printf("Insert failed,err:%v\n", err)

		return

	}

	fmt.Println("update data successd:", result)

	rowsaffected, err := result.RowsAffected()

	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n", err)

		return

	}

	fmt.Println("Affected rows:", rowsaffected)

}

//删除数据

func DeleteData(DB *sql.DB) {
	result, err := DB.Exec("delete from users where id=?", 2)

	if err != nil {
		fmt.Printf("Insert failed,err:%v\n", err)

		return

	}

	fmt.Println("delete data successd:", result)

	rowsaffected, err := result.RowsAffected()

	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v\n", err)

		return

	}

	fmt.Println("Affected rows:", rowsaffected)

}

// 事务
func TXInsert(db *sql.DB) {
	/*tx := db.MustBegin()
	tx.MustExec(`INSERT INTO student VALUES ('1', 'Jack', 'Jack', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '1', '2018-06-26 17:08:35');`)
	tx.MustExec(`INSERT INTO student VALUES ('2', 'Emily', 'Emily', 'England', '', '', 'http://img2.imgtn.bdimg.com/it/u=3588772980,2454248748&fm=27&gp=0.jpg', '2', null);`)
	err = tx.Commit()
	if err != nil {
	log.Fatalln(err)
	}*/
}


// GORM
// https://learnku.com/docs/gorm/v1/query/3786
// http://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/gorm/gorm%E7%94%A8%E6%B3%95%E4%BB%8B%E7%BB%8D.html

// UserInfo 用户信息
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Hobby string
}

func GORM(){
	u1 := UserInfo{1, "枯藤", "男", "篮球"}
	u2 := UserInfo{2, "topgoer.com", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	db.First(&u, 1) // 找到id为1的产品
	db.First(&u, "hobby = ?", "L1212") // 找出 hobby 为 l1212 的产品
	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}