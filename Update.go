package HandleBolt

import (
	"bytes"
	"github.com/boltdb/bolt"
	"sync"
	"fmt"
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

//TODO:创建数据库
func (D *Db)CreateBucket(BucketName string)error{
	err:=D.Db.Update(func(tx *bolt.Tx) error {
		_,err:=tx.CreateBucket([]byte(BucketName))
		return err
	})
	return err
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
			//fmt.Printf("key is %v,value is %s\n", bytesToInt(k), v)
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
//范围搜索
func (D *Db) SearchRangeData(BucketName,startData,endData string)map[string]string{
	result:=make(map[string]string,0)
	 if err:=D.Db.View(func(tx *bolt.Tx) error {
		// Assume our events bucket exists and has RFC3339 encoded time keys.
		c := tx.Bucket([]byte(BucketName)).Cursor()

		// Our time range spans the 90's decade.
		min := []byte(startData)
		max := []byte(endData)

		// Iterate over the 90's.
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			//fmt.Printf("%s: %s\n", k, v)
			result[string(k)]=string(v)
		}

		return nil
	});err!=nil{

	 }
	return result
}



//一般没啥用 存储格式为  1----value   2-----value
func (D *Db)SaveDataBySequence(BucketName string,datas []string)error{
	for i := 0; i < len(datas); i++ {
		if err := D.Db.Update(func(tx *bolt.Tx) error {
			if _, err := tx.CreateBucketIfNotExists([]byte(BucketName)); err != nil {
				return err
			}
			b := tx.Bucket([]byte(BucketName))
			id, _ := b.NextSequence()
			err := b.Put(u64tob(id), []byte(datas[i]))
			fmt.Println(u64tob(id))
			return err
		}); err != nil {
			return err
		}

	}
	return nil
}

func (D *Db)SearchAllSequenceData(BucketName string)(map[int64]string)  {
	//sync map只是随便练手写的，不适配当前情况
	data:=sync.Map{}
	_=D.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		cur := b.Cursor()
		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			data.Store(bytesToInt64(k),string(v))
			//fmt.Printf("key is %v,value is %s\n", bytesToInt(k), v)
		}
		return nil
	})
	result:=make(map[int64]string,0)
	data.Range(func(key, value interface{}) bool {
		result[key.(int64)]=value.(string)
		return true
	})
	return result
}


//插入大量数据
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

//TODO：批量删除相关数据

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
