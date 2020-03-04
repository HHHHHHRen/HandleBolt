### github 类库HandleBolt
#### 使用方法
 名称 | 传递参数 | 返回值
:--:|:--:|:--:
InitDb |BucketName | Db类型的结构体和err

#### 新增和修改功能的实现

* #### *ChangeData*
```go
func (D *Db)ChangeData(BucketName,key,value interface{})error
```
#### 查询表中的全部数据
* #### *SearchAllData*
```go
func (D *Db)SearchAllData(BucketName string)map[string]string
```

#### 范围搜索
* #### *SearchRangeData*
```go
func (D *Db) SearchRangeData(BucketName,startData,endData string)map[string]string
```

#### key为自增序列的插入数据
* #### *SaveDataBySequence*
```go
func (D *Db)SaveDataBySequence(BucketName string,datas []string)error
```

#### 新增数据库
* #### *CreateBucket*
```go
func (D *Db)CreateBucket(BucketName string)error
```

#### 查询表中的自增序列值
* #### *SearchAllSequenceData*
```go
func (D *Db)SearchAllSequenceData(BucketName string)(map[int64]string)
```

#### 批量插入数据
* #### *InertDatas*
```go
func (D *Db)InertDatas(BucketName string,datas map[string]string)
```

#### 按照key值查询数据
* #### *ReadOneData*
```go
func(D *Db) ReadOneData(BucketName , key string)(string)
```

#### 删除数据库中的对应key-value

* #### *RemoveData*
```go
func(D *Db) RemoveData(BucketName,key string)error
```

#### 删除Bucket
* #### *RemoveBucket*
```go
func (D *Db)RemoveBucket (BucketName string)error
```
