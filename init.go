package HandleBolt

import (
	"github.com/boltdb/bolt"
	//"time"
)

type Db struct {
	Db *bolt.DB
}

//初始化数据库
func InitDb(BucketName string)(*Db,error){
	//打开数据库
	db,_:=bolt.Open("mydb.db",0600,nil)//&bolt.Options{Timeout:1*time.Second})
	var D=new(Db)
	D.Db=db
	//新建数据库表
	err:=D.Db.Update(func(tx *bolt.Tx) error {
		_,err:=tx.CreateBucket([]byte(BucketName))
		return err
	})

	return D,err
}
