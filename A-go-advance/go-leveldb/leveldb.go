package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	//Use bloom filter:
	o := &opt.Options{
		Filter: filter.NewBloomFilter(2),
	}
	db, err := leveldb.OpenFile("store/db", o)

	if err != nil {
		fmt.Println("leveldb.OpenFile err:", err)
	}
	defer db.Close()

	// 批量操作
	//fmt.Println("***************************************************")
	//batch := new(leveldb.Batch)
	//batch.Put([]byte("foo"), []byte("value"))
	//batch.Put([]byte("bar"), []byte("another value"))
	//batch.Delete([]byte("baz"))
	//err = db.Write(batch, nil)
	//fmt.Println("***************************************************")

	//data, err := read("zhunh", db)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("read= %s\n", data)

	//it(db)
	//seekIt("皆可",db)
	//rangeIt("key02", "zhunh", db)
	preIt("key", db)
}

//data, err := db.Get([]byte("zhunh"), nil)
func read(key string, db *leveldb.DB) ([]byte, error) {
	return db.Get([]byte(key), nil)
}

//添加
func put(key, value string, db *leveldb.DB) error {
	//err = db.Put([]byte("皆可"), []byte("啊啊咖啡机"), nil)
	return db.Put([]byte(key), []byte(value), nil)
}

//删除
func delete(key string, db *leveldb.DB) error {
	return db.Delete([]byte(key), nil)
}

//迭代
func it(db *leveldb.DB) {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%s:%s\n", key, value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Printf("iter.Error:", err)
	}
}

//查找然后迭代
func seekIt(key string, db *leveldb.DB) {
	iter := db.NewIterator(nil, nil)
	for ok := iter.Seek([]byte(key)); ok; ok = iter.Next() {
		// Use key/value.
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%s:%s\n", key, value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Printf("iter.Error:", err)
	}
}

//遍历子集范围,区间查找
func rangeIt(start, end string, db *leveldb.DB) {
	iter := db.NewIterator(&util.Range{Start: []byte(start), Limit: []byte(end)}, nil)
	for iter.Next() {
		fmt.Printf("[%s]:%s\n", iter.Key(), iter.Value())
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Printf("iter.Error:", err)
	}
}

//前缀查找
func preIt(prefix string, db *leveldb.DB) {
	iter := db.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	for iter.Next() {
		// Use key/value.
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("%s:%s\n", key, value)
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		fmt.Println("iter.Error:", err)
	}
}
