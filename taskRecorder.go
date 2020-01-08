package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	remote = "cameron:linux@tcp(66.45.159.238:3306)/work_records"
	LAN = "cameron:linux@tcp(10.0.0.104:3306)/work_records"
	localhost  = "cameron:linux@tcp(127.0.0.1:3306)/work_records"
)

func main() {
	db, err := sql.Open("mysql",
		LAN)
	if err != nil {
		log.Fatal(err)
	}

	name, description, time, difficulty := getInput()
	query := "INSERT INTO tasks (name, description, time_taken, difficulty) VALUES (\"" + name + "\",\"" + description + "\",\"" + time + "\",\"" + difficulty + "\");"
	fmt.Println("Query entered: \n	", query)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(query)
	defer db.Close()
}

func getInput() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of Task: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Print("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSuffix(description, "\n")

	fmt.Print("Time Taken: ")
	time, _ := reader.ReadString('\n')
	time = strings.TrimSuffix(time, "\n")

	fmt.Print("Difficulty: ")
	difficulty, _ := reader.ReadString('\n')
	difficulty = strings.TrimSuffix(difficulty, "\n")

	return name, description, time, difficulty
}
