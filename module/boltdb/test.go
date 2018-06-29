package main

import (
	"log"
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

func getDB() (*bolt.DB, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	return db,nil
}

func testSet(root,key,value string)  {
	db,err := getDB()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(root))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), []byte(value))
	})
}

func testGet(root,key string) string {
	db,errDB := getDB()
	if errDB != nil {
		log.Println(errDB)
	}
	defer db.Close()

	var result []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(root))
		if b != nil{
			v := b.Get([]byte(key))
			result = v
			return nil
		}
		return errors.New("Doesn't have key")
	})
	if err != nil{
		log.Println(err.Error())
		return ""
	}
	return string(result)
}

//func testGet(root,key string) byte[] {
//	db,err := getDB()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(root))
//		v := b.Get([]byte(key))
//		return v
//	})
//	return nil
//}

func main() {
	testSet("enviroment","test","1")
	log.Println(testGet("enviroment","test"))
	log.Println(testGet("heihei","test"))
}