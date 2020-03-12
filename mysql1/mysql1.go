package main


// 1 连接 mysql


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"time"
	"fmt"
) //利用匿名导入，因为只需要它的驱动

func main() {

	db, err := sql.Open("mysql", "tester:test123@(127.0.0.1:3306)/testdb2?parseTime=true")

	if err != nil{
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}


	{ // Create a new table
		query := `
            CREATE TABLE IF NOT EXISTS users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            )ENGINE=InnoDB DEFAULT CHARSET=utf8;`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ //Inserting our first user
		username := "plh"
		password := "123456"
		createdAt := time.Now()
		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

		if err != nil{
			log.Fatal(err)
		}
		userID, err:= result.LastInsertId()
		fmt.Println(userID)



	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)  // 这个地方应该换成一个 类

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 2).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
	 type user struct {
	 	id int
	 	username string
	 	password string
	 	createdAt time.Time
	 }

	 rows , err := db.Query(`SELECT id, username, password, created_at FROM users`)
	 if err != nil{
	 	log.Fatal(err)
	 }
	 defer rows.Close()

	 var users [] user // 这个地方是切片
	 for rows.Next(){
	 	var u user
	 	err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

	 	if err != nil{
	 		log.Fatal(err)
		}
		users = append(users, u)

	 }
		// check err
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
		//这个地方打印的确不够美观
		for i, u := range users{
			fmt.Printf("index %d\t id %d\t  username %s \t password %s\t createdAt %s\n", i+1, u.id, u.username, u.password, u.createdAt.String())
		}

	}

	{ // 删除
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}

	}




}