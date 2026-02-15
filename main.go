package main

import db "github.com/Pankaj-Gupta25/Go_file_upload/db"

func main() {

	db.Connect()
	r := setupRoutes()
	r.Run(":4000")
}
