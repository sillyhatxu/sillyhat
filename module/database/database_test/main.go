package main

import (
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
	"log"
)

var db = &sql.DB{}

func init(){
	//db,_ = sql.Open("mysql", "deja_cloud:deja_cloud@tcp(deja-dt.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/ocb_syncer")
	db,_ = sql.Open("mysql", "deja_cloud:deja_cloud@tcp(deja-dt.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/payment")
	connMaxLifetime := 10
	maxIdleConns := 10
	maxOpenConns := 100
	fmt.Printf("connMaxLifetime:%d\n", connMaxLifetime)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	//SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//If n <= 0, no idle connections are retained.
	fmt.Printf("maxIdleConns:%d\n", maxIdleConns)
	db.SetMaxIdleConns(maxIdleConns)

	//SetMaxOpenConns sets the maximum number of open connections to the database.
	//If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit
	//If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
	fmt.Printf("maxOpenConns:%d\n", maxOpenConns)
	db.SetMaxOpenConns(maxOpenConns)
}

//var db = &sql.DB{}

//func init(){
//	db, _ := sql.Open("mysql", "deja_cloud:deja_cloud@tcp(deja-dt.ccf2gesv8s9h.ap-southeast-1.rds.amazonaws.com:3306)/ocb_syncer")
	//connMaxLifetime := 10
	//maxIdleConns := 10
	//maxOpenConns := 100
	//fmt.Printf("connMaxLifetime:%d\n", connMaxLifetime)
	//db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	//
	////SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	////If n <= 0, no idle connections are retained.
	//fmt.Printf("maxIdleConns:%d\n", maxIdleConns)
	//db.SetMaxIdleConns(maxIdleConns)
	//
	////SetMaxOpenConns sets the maximum number of open connections to the database.
	////If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit
	////If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
	//fmt.Printf("maxOpenConns:%d\n", maxOpenConns)
	//db.SetMaxOpenConns(maxOpenConns)
//}

func update(){
	//方式1 update
	start := time.Now()
	for i := 1001;i<=1100;i++{
		db.Exec("UPdate ocb_syncer.userinfo set age=? where uid=? ",i,i)
	}
	end := time.Now()
	fmt.Println("方式1 update total time:",end.Sub(start).Seconds())

	//方式2 update
	start = time.Now()
	for i := 1101;i<=1200;i++{
		stm,_ := db.Prepare("UPdate ocb_syncer.userinfo set age=? where uid=? ")
		stm.Exec(i,i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 update total time:",end.Sub(start).Seconds())

	//方式3 update
	start = time.Now()
	stm,_ := db.Prepare("UPdate ocb_syncer.userinfo set age=? where uid=?")
	for i := 1201;i<=1300;i++{
		stm.Exec(i,i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 update total time:",end.Sub(start).Seconds())

	//方式4 update
	start = time.Now()
	tx,_ := db.Begin()
	for i := 1301;i<=1400;i++{
		tx.Exec("UPdate ocb_syncer.userinfo set age=? where uid=?",i,i)
	}
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 update total time:",end.Sub(start).Seconds())

	//方式5 update
	start = time.Now()
	for i := 1401;i<=1500;i++{
		tx,_ := db.Begin()
		tx.Exec("UPdate ocb_syncer.userinfo set age=? where uid=?",i,i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 update total time:",end.Sub(start).Seconds())

}

func delete(){
	//方式1 delete
	start := time.Now()
	for i := 1001;i<=1100;i++{
		db.Exec("DELETE FROM ocb_syncer.userinfo WHERE uid=?",i)
	}
	end := time.Now()
	fmt.Println("方式1 delete total time:",end.Sub(start).Seconds())

	//方式2 delete
	start = time.Now()
	for i := 1101;i<=1200;i++{
		stm,_ := db.Prepare("DELETE FROM ocb_syncer.userinfo WHERE uid=?")
		stm.Exec(i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 delete total time:",end.Sub(start).Seconds())

	//方式3 delete
	start = time.Now()
	stm,_ := db.Prepare("DELETE FROM ocb_syncer.userinfo WHERE uid=?")
	for i := 1201;i<=1300;i++{
		stm.Exec(i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 delete total time:",end.Sub(start).Seconds())

	//方式4 delete
	start = time.Now()
	tx,_ := db.Begin()
	for i := 1301;i<=1400;i++{
		tx.Exec("DELETE FROM ocb_syncer.userinfo WHERE uid=?",i)
	}
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 delete total time:",end.Sub(start).Seconds())

	//方式5 delete
	start = time.Now()
	for i := 1401;i<=1500;i++{
		tx,_ := db.Begin()
		tx.Exec("DELETE FROM ocb_syncer.userinfo WHERE uid=?",i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 delete total time:",end.Sub(start).Seconds())

}

func query(){

	//方式1 query
	start := time.Now()
	rows,_ := db.Query("SELECT uid,username FROM ocb_syncer.userinfo")
	defer rows.Close()
	for rows.Next(){
		var name string
		var id int
		if err := rows.Scan(&id,&name); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end := time.Now()
	fmt.Println("方式1 query total time:",end.Sub(start).Seconds())

	//方式2 query
	start = time.Now()
	stm,_ := db.Prepare("SELECT uid,username FROM ocb_syncer.userinfo")
	defer stm.Close()
	rows,_ = stm.Query()
	defer rows.Close()
	for rows.Next(){
		var name string
		var id int
		if err := rows.Scan(&id,&name); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end = time.Now()
	fmt.Println("方式2 query total time:",end.Sub(start).Seconds())


	//方式3 query
	start = time.Now()
	tx,_ := db.Begin()
	defer tx.Commit()
	rows,_ = tx.Query("SELECT uid,username FROM ocb_syncer.userinfo")
	defer rows.Close()
	for rows.Next(){
		var name string
		var id int
		if err := rows.Scan(&id,&name); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("name:%s ,id:is %d\n", name, id)
	}
	end = time.Now()
	fmt.Println("方式3 query total time:",end.Sub(start).Seconds())
}

func insert1() {

	//方式1 insert
	//strconv,int转string:strconv.Itoa(i)
	start := time.Now()
	for i := 10010;i<=11000;i++{
		//每次循环内部都会去连接池获取一个新的连接，效率低下
		db.Exec("INSERT INTO ocb_syncer.userinfo(uid,username,age) values(?,?,?)",i,"user"+strconv.Itoa(i),i-1000)
	}
	end := time.Now()
	fmt.Println("方式1 insert total time:",end.Sub(start).Seconds())
}
func insert2()  {
	//方式2 insert
	start := time.Now()
	for i := 11010;i<=12000;i++{
		//Prepare函数每次循环内部都会去连接池获取一个新的连接，效率低下
		stm,_ := db.Prepare("INSERT INTO ocb_syncer.userinfo(uid,username,age) values(?,?,?)")
		stm.Exec(i,"user"+strconv.Itoa(i),i-1000)
		stm.Close()
	}
	end := time.Now()
	fmt.Println("方式2 insert total time:",end.Sub(start).Seconds())
}

func insert3()  {
	//方式3 insert
	start := time.Now()
	stm,_ := db.Prepare("INSERT INTO ocb_syncer.userinfo(uid,username,age) values(?,?,?)")
	for i := 12010;i<=13000;i++{
		//Exec内部并没有去获取连接，为什么效率还是低呢？
		stm.Exec(i,"user"+strconv.Itoa(i),i-1000)
	}
	stm.Close()
	end := time.Now()
	fmt.Println("方式3 insert total time:",end.Sub(start).Seconds())
}

func insert4()  {
	//方式4 insert
	start := time.Now()
	//Begin函数内部会去获取连接
	tx,_ := db.Begin()
	for i := 13010;i<=14000;i++{
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec("INSERT INTO ocb_syncer.userinfo(uid,username,age) values(?,?,?)",i,"user"+strconv.Itoa(i),i-1000)
	}
	//最后释放tx内部的连接
	tx.Commit()

	end := time.Now()
	fmt.Println("方式4 insert total time:",end.Sub(start).Seconds())
}
func insert5()  {
	//方式5 insert
	start := time.Now()
	for i := 14010;i<=15000;i++{
		//Begin函数每次循环内部都会去连接池获取一个新的连接，效率低下
		tx,_ := db.Begin()
		tx.Exec("INSERT INTO ocb_syncer.userinfo(uid,username,age) values(?,?,?)",i,"user"+strconv.Itoa(i),i-1000)
		//Commit执行后连接也释放了
		tx.Commit()
	}
	end := time.Now()
	fmt.Println("方式5 insert total time:",end.Sub(start).Seconds())
}

func testInsert1() {
	stm,_ := db.Prepare("INSERT INTO payment.payment (id, currency, customer_id, order_group_id, stripe_source_id, stripe_charge_id, total_amount, created_date, last_modified_date) VALUES (?,?,?,?,?,?,?,now(),now())")
	//Exec内部并没有去获取连接，为什么效率还是低呢？
	stm.Exec("2","SGD",1,"1","1","1",8000)
	stm.Close()
}
func testInsert2(args ...interface{}) {
	stm,_ := db.Prepare("INSERT INTO payment.payment (id, currency, customer_id, order_group_id, stripe_source_id, stripe_charge_id, total_amount, created_date, last_modified_date) VALUES (?,?,?,?,?,?,?,now(),now())")
	//Exec内部并没有去获取连接，为什么效率还是低呢？
	stm.Exec(args...)
	stm.Close()
}
func main() {
	//testInsert1()
	testInsert2("5","SGD",1,"1","1","1",8000)
//	#####insert3()    update3    delete3   query3
//	insert1()
//	insert2()
//	insert3()
//	insert4()
//	insert5()
	//query()
	//update()
	//query()
	//delete()
}