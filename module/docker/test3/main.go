package main

import (
	"strings"
	"hash/crc32"
	"log"
	"sillyhat-golang-tool/sillyhat_database"
	"time"
)

const (
	//dataSourceName      = "deja_cloud:Xj$nbyndfwlb#C@tcp(deja-production.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/inventory"
	dataSourceName      = "deja_cloud:deja_cloud@tcp(deja-dt.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/shopping_bag"
	shoppingbagSQL = "UPDATE shopping_bag.shopping_bag_item SET inventory_item_id = ? WHERE id = ?"
	//productInventorySQL = "INSERT INTO inventory.product_inventory (id,auto_deleted,quantity,size,product_id) VALUES (?,TRUE,?,?,?) ON DUPLICATE KEY UPDATE auto_deleted = TRUE,quantity = ?,size = ?,product_id = ?"
)


func main() {
	log.Println("mysqlClient init start")
	var mysqlClient sillyhat_database.MySQLClient
	mysqlClient = sillyhat_database.MySQLClient{DataSourceName:dataSourceName}
	mysqlClient.Init()
	log.Println("mysqlClient Init end")

	log.Println("query start")
	shoppingbagArray := query(mysqlClient)
	log.Println("query end")
	for _,shoppingbagDTO := range shoppingbagArray{
		time.Sleep(1 * time.Second)
		log.Printf("shoppingbagDTO id: %v ; shop item id : %v ; size : %v",shoppingbagDTO.id,shoppingbagDTO.shop_item_id,shoppingbagDTO.size)
	}
}


type shoppingbagDTO struct{
	id int64
	shop_item_id int64
	size string
}

func query(mysqlClient sillyhat_database.MySQLClient) []shoppingbagDTO {
	//resultRows,err := mysqlClient.QueryList("SELECT id,shop_item_id,size FROM shopping_bag.shopping_bag_item")
	//if err != nil{
	//	log.Println("query error",err.Error())
	//	return nil
	//}
	var shoppingbagArray []shoppingbagDTO

	tx,err := mysqlClient.GetConnection().Begin()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer tx.Commit()
	rows,err := tx.Query("SELECT id,shop_item_id,size FROM shopping_bag.shopping_bag_item")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	for rows.Next(){
		var id int64
		var shop_item_id int64
		var size string
		if err := rows.Scan(&id,&shop_item_id,&size); err != nil {
			log.Fatal(err)
		}
		shoppingbagArray = append(shoppingbagArray,*&shoppingbagDTO{shop_item_id:shop_item_id,id:id,size:size})
	}
	log.Println("query end")
	return shoppingbagArray
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