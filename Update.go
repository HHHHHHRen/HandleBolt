package HandleBolt

import (
	"github.com/boltdb/bolt"
	"sync"
)
//TODO:需要支持key,value的值是任意类型

//新增修改单个的数据项
func ChangeData(BucketName,key,value interface{})error{
 	if err:=Db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BucketName.(string))); err != nil {
			return err
		}
 		b := tx.Bucket([]byte(BucketName.(string)))
		if err := b.Put([]byte(key.(string)), []byte(value.(string))); err != nil {
			//log.Fatal(err)
			return err
		}
		return nil
	});err!=nil{
		return err
	}
	return nil
}

//查询数据库中的全部数据，返回值采用map
func SearchAllData(BucketName string)map[string]string{
	data:=sync.Map{}
	_=Db.View(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BucketName)); err != nil {
			return err
		}
		b := tx.Bucket([]byte(BucketName))
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			data.Store(k,v)
			//fmt.Printf("key is %s,value is %s\n", k, v)
		}
		return nil
	})
	result:=make(map[string]string,0)
	data.Range(func(key, value interface{}) bool {
		result[key.(string)]=value.(string)
		return true
	})
	return result
}

//TODO:查询范围内的数据

//TODO:按照自增序列存储数据  1----value  2---value

//TODO:获取所有数据



//同时插入大量数据
func InertDatas(BucketName string,datas map[string]string)  {
	var wg sync.WaitGroup
	wg.Add(len(datas))
	for k,v:=range datas{
		go func(key,value string) {
			wg.Done()
			_=Db.Batch(func(tx *bolt.Tx) error {
				return tx.Bucket([]byte(BucketName)).Put([]byte(key), []byte(value))
			})
		}(k,v)
	}
	wg.Wait()
}

//查询单条数据
func ReadOneData(BucketName , key string)(string){
	var result []byte
	if err:=Db.View(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BucketName)); err != nil {
			return err
		}
		b := tx.Bucket([]byte(BucketName))
		result=b.Get([]byte(key))
		//return string(result),nil
	});err!=nil{
		return ""
	}
	return string(result)
}

//删除数据库中某个内容
func RemoveData(BucketName,key string)error{
	if err:=Db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(BucketName)); err != nil {
			return err
		}
		b := tx.Bucket([]byte(BucketName))
		if err := b.Delete([]byte(key)); err != nil {
			//log.Fatal(err)
			return err
		}
		return nil
	});err!=nil{
		return err
	}
	return nil
}

//删库
func RemoveBucket (BucketName string)error{
	if err:=Db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(BucketName))
		if err:=b.DeleteBucket([]byte(BucketName));err!=nil{
			return err
		}
		return nil
	});err!=nil{
		return err
	}
	return nil
}
