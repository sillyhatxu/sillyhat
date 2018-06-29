package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"time"
)

func main() {
	db,err := sql.Open("mysql", "deja_cloud:deja_cloud@tcp(deja-dt.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/shopping_bag?timeout=30s")
	if err != nil{
		log.Println("Open error")
		log.Println(err.Error())
	}
	db.SetConnMaxLifetime(time.Duration(60) * time.Second)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	tx,err := db.Begin()
	if err != nil {
		log.Println("Begin error")
		log.Println(err.Error())
	}
	defer tx.Commit()

	rows,err := tx.Query("SELECT id,shop_item_id,size FROM shopping_bag.shopping_bag_item")
	if err != nil {
		log.Println("Query error")
		log.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		var id int64
		var shop_item_id int64
		var size string
		if err := rows.Scan(&id,&shop_item_id,&size); err != nil {
			log.Fatal(err)
		}
		log.Printf("shoppingbagDTO id: %v ; shop item id : %v ; size : %v",id,shop_item_id,size)
	}
}
