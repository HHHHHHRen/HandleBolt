package HandleBolt

import (
	"github.com/boltdb/bolt"
	"time"
)

var(
	Db *bolt.DB
)
func initDB(BucketName string) *bolt.DB {
	db,_:=bolt.Open("mydb.db",0600,&bolt.Options{Timeout:1*time.Second})
	Db=db
	Db.Update(func(tx *bolt.Tx) error {
		_,err:=tx.CreateBucket([]byte(BucketName))
		return err
	})
	return Db
}
