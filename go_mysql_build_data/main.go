package main

import (
	"database/sql"
	"fmt"
	"go_demo/go_mysql_build_data/name"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func initInstance() *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", "root", "123456", "127.0.0.1", 3306, "test_db", "utf8mb4")
	db, err := sql.Open("mysql", connString)
	if err != nil || db == nil {
		log.Fatalf("connect to mysql failed, err:%+v, db:%+v", err, db)
	}

	return db
}

func queryOneRow(db *sql.DB) {
	sql := "select name, age, sex from student where id = 10"
	var name string
	var age, sex int
	err := db.QueryRow(sql).Scan(&name, &age, &sex)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("name: %s, age: %d, sex: %d\n", name, age, sex)
	}
}

func queryMultiRows(db *sql.DB) {
	sql := "select kugouid, packet_key, packet_value from student limit 10"
	r, err := db.Query(sql)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var name string
	var age, sex int
	for r.Next() {
		r.Scan(&name, &age, &sex)
		fmt.Printf("name: %s, age: %d, sex: %d\n", name, age, sex)
	}
}

/*
CREATE TABLE `student` (

		`id` int(11) NOT NULL AUTO_INCREMENT,
		`name` varchar(100) NOT NULL DEFAULT '',
		`age` int(11) NOT NULL DEFAULT '0',
		`sex` tinyint(4) NOT NULL DEFAULT '0',
		`weight` int(11) NOT NULL DEFAULT '0',
		PRIMARY KEY (`id`)
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
*/
func generateData(db *sql.DB) {
	total := 300
	rand.Seed(time.Now().UnixNano())

	insertSql := "insert into student (name,age,sex,weight) values(?,?,?,?)"

	for i := 0; i < total; i++ {
		name := name.RandomName()
		age := rand.Intn(20) + 15    // 随机生成15 ~ 35
		sex := rand.Intn(2)          // 0为男, 1为女
		weight := rand.Intn(60) + 60 // 随机生成60 ~ 120

		if i%200 == 0 {
			time.Sleep(time.Millisecond * 100)
		}

		r, err := db.Exec(insertSql, name, age, sex, weight)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			i, _ := r.LastInsertId()
			fmt.Printf("i: %v\n", i)
		}
	}

}

func main() {
	db := initInstance()
	// queryOneRow(db)
	// queryMultiRows(db)
	generateData(db)
}
