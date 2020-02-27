package HandleBolt

import (
	"github.com/boltdb/bolt"
	"sync"
)

//TODO:需要支持key,value的值是任意类型
//目前支持的情况只是存储json字符串。
//新增修改单个的数据项
func (D *Db)ChangeData(BucketName,key,value interface{})error{
 	if err:=D.Db.Update(func(tx *bolt.Tx) error {
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
func (D *Db)SearchAllData(BucketName string)map[string]string{
	//sync map只是随便练手写的，不适配当前情况
	data:=sync.Map{}
	_=D.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			data.Store(string(k),string(v))
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
func (D *Db)InertDatas(BucketName string,datas map[string]string)  {
	var wg sync.WaitGroup
	wg.Add(len(datas))
	for k,v:=range datas{
		go func(key,value string) {
			defer wg.Done()
			_=D.Db.Batch(func(tx *bolt.Tx) error {
				return tx.Bucket([]byte(BucketName)).Put([]byte(key), []byte(value))
			})
		}(k,v)
	}
	wg.Wait()
}

//查询单条数据
func(D *Db) ReadOneData(BucketName , key string)(string){
	var result []byte
	if err:=D.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		result=b.Get([]byte(key))
		return nil
		//return string(result),nil
	});err!=nil{
		return ""
	}
	return string(result)
}

//删除数据库中某个内容
func(D *Db) RemoveData(BucketName,key string)error{
	if err:=D.Db.Update(func(tx *bolt.Tx) error {
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
func (D *Db)RemoveBucket (BucketName string)error{
	if err:=D.Db.Update(func(tx *bolt.Tx) error {
		//b:=tx.Bucket([]byte(BucketName))
		err:=tx.DeleteBucket([]byte(BucketName))
		if err!=nil{
			return err
		}
		return nil
	});err!=nil{
		return err
	}
	return nil
}
