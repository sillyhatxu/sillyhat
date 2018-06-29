package main

import (
	"time"
	"log"
	"strconv"
	"hash/crc32"
	"strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sillyhat-golang-tool/sillyhat_database"
)


const (
	dataSourceName      = "deja_cloud:Xj$nbyndfwlb#C@tcp(deja-production.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/inventory"
	productInventorySQL = "INSERT INTO inventory.product_inventory (id,auto_deleted,quantity,size,product_id) VALUES (?,TRUE,?,?,?)"
	//productInventorySQL = "INSERT INTO inventory.product_inventory (id,auto_deleted,quantity,size,product_id) VALUES (?,TRUE,?,?,?) ON DUPLICATE KEY UPDATE auto_deleted = TRUE,quantity = ?,size = ?,product_id = ?"
)

func main() {
	var mysqlClient sillyhat_database.MySQLClient
	mysqlClient = sillyhat_database.MySQLClient{DataSourceName:dataSourceName}
	mysqlClient.Init()

	log.Println("start project")
	start := time.Now()

	var inventoryArray []inventoryDTO
	inventoryArray = query(mysqlClient.GetConnection())
	log.Printf("total : %v",len(inventoryArray))
	for _,inventoryDTO := range inventoryArray{
		insert(mysqlClient.GetConnection(),inventoryDTO.shop_item_id,inventoryDTO.quantity,inventoryDTO.size)
	}
	end := time.Since(start)
	log.Println("Execute time : ", end)
	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//<-done
}

type inventoryDTO struct{
	shop_item_id int64
	quantity int
	size string
}

func query(db *sql.DB) []inventoryDTO {
	log.Println("query start")
	var inventoryArray [] inventoryDTO
	tx,_ := db.Begin()
	defer tx.Commit()
	rows,_ := tx.Query("SELECT shop_item_id,quantity,size FROM inventory.inventory_item")
	defer rows.Close()
	for rows.Next(){
		var shop_item_id int64
		var quantity int
		var size string
		if err := rows.Scan(&shop_item_id,&quantity,&size); err != nil {
			log.Fatal(err)
		}
		inventoryArray = append(inventoryArray,*&inventoryDTO{shop_item_id:shop_item_id,quantity:quantity,size:size})
	}
	log.Println("query end")
	return inventoryArray
}

func insert(db *sql.DB,shop_item_id int64,quantity int,size string) {
	id := hashCodeTrimUpper(strconv.FormatInt(shop_item_id,10) + "_" + size)
	log.Printf("id : %v ; shop item id : %v ; size : %v ; quantity : %v;",id,shop_item_id,size,quantity)
	stm,_ := db.Prepare(productInventorySQL)
	stm.Exec(id,quantity,size,shop_item_id)
	stm.Close()
}

func hashCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}

func hashCodeTrim(str string) int {
	return hashCode(strings.Replace(str, " ", "", -1))
}

func hashCodeTrimUpper(str string) int {
	return hashCodeTrim(strings.ToUpper(str))
}