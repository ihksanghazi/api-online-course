package main

import (
	"fmt"

	"github.com/ihksanghazi/api-online-course/databases"
)

func main() {
	databases.ConnectDB()

	fmt.Println("Sukses Koneksi")
}
