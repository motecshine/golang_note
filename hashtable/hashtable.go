package hashtable

import (
	"hash/fnv"
	"log"
	"sync"
)

var ChainNodeSize uint32 = 8 // 假设Key单项列表默认的长度是8个Node
var KeyChains *KeyChainNode
var Mask uint32 = 8 // 当resize hashtable 时这个值也会跟随 ChainNodeSize 变化

type Bucket struct {
	Data       interface{}
	Key        string
	NextBucket *Bucket //  解决hash冲突
}

type KeyChainNode struct {
	Key              uint32
	NextKeyChainNode *KeyChainNode
	Bucket           *Bucket
	Refcount         uint32 //计算冲突量有多少 默认是0
}

type Hashtable struct {
	mu sync.RWMutex
}

func HashingKey(str string) uint32 {
	hashModel := fnv.New32a()
	hashModel.Write([]byte(str))
	return hashModel.Sum32() % Mask
}

func init() {
	KeyChains = InitKeyChains()
}

/* 初始化一个长度为ChainNodeSize的node */
func InitKeyChains() *KeyChainNode {
	var prevNode, tmpNode *KeyChainNode
	var locking sync.RWMutex
	locking.Lock()
	for i := ChainNodeSize; i >= 1; i-- {
		if i == ChainNodeSize {
			tmpNode = &KeyChainNode{
				Key:              i,
				NextKeyChainNode: nil,
				Bucket:           nil,
				Refcount:         0,
			}
		} else {
			tmpNode = &KeyChainNode{
				Key:              i,
				NextKeyChainNode: prevNode,
				Bucket:           nil,
				Refcount:         0,
			}
		}
		prevNode = tmpNode
	}
	locking.Unlock()
	return tmpNode
}

func FindKeyChainNode(key uint32, KeyChains *KeyChainNode) *KeyChainNode {
	var i uint32
	for i = 1; i <= ChainNodeSize; i++ {
		if KeyChains.Key == key {
			return KeyChains
		}
		KeyChains = KeyChains.NextKeyChainNode
	}
	return nil
}

func (hashtable *Hashtable) Add(key string, value interface{}) {
	defer hashtable.mu.Unlock()
	mapKey := HashingKey(key)
	node := FindKeyChainNode(mapKey, KeyChains)
	if node == nil {
		log.Fatal("cant find node")
	}
	hashtable.mu.Lock()
	if node.Refcount == 0 {
		node.Bucket = &Bucket{
			Key:        key,
			Data:       value,
			NextBucket: nil,
		}
		// 当前node的bucket个数
		node.Refcount++
	} else {
		tmpNode := node
		for {
			if tmpNode.Bucket.Key == key {
				hashtable.Update(key, value)
				break
			}

			// 如果 key相同 找到bucket 相同的key 然后覆盖值
			if tmpNode.Bucket.NextBucket == nil {
				tmpNode.Bucket.NextBucket = &Bucket{
					Key:        key,
					Data:       value,
					NextBucket: nil,
				}
				// 当前node的bucket个数 即有多少冲突
				tmpNode.Refcount++
				break
			}
			if tmpNode.Bucket == nil {
				break
			}
			tmpNode.Bucket = tmpNode.Bucket.NextBucket
		}
	}

}

func (hashtable *Hashtable) Delete(key string) {
	mapKey := HashingKey(key)
	node := FindKeyChainNode(mapKey, KeyChains)
	if node == nil {
		log.Fatal("cant find node\n")
	}
	bucket := node.Bucket

	for {
		// 要删除的bucket在链表的头部
		if bucket.Key == key {
			// 当前的bucket只有一个，未存在冲突
			if node.Refcount == 1 {
				// 释放当前bucket的应用
				node.Bucket = nil
				break
			} else {
				// 如果Refcount > 2
				node.Bucket = nil
				node.Bucket = bucket.NextBucket
				break
			}
		}

		// 要删除的bucket在链表的中间或尾部
		if bucket.NextBucket != nil && bucket.NextBucket.Key == key {
			// 需要切断要删除的bucket 对下一个bucket的引用
			tmpBucket := bucket.NextBucket
			bucket.NextBucket = tmpBucket.NextBucket
			tmpBucket.NextBucket = nil
			break
		}
		if bucket.NextBucket == nil {
			break
		}
		bucket = bucket.NextBucket
	}
	node.Refcount--
}

func (hashtable *Hashtable) Update(key string, newValue interface{}) {
	mapKey := HashingKey(key)
	node := FindKeyChainNode(mapKey, KeyChains)
	if node == nil {
		log.Fatal("cant find node\n")
	}
	bucket := node.Bucket
	for {
		// 如果 key相同 找到bucket 相同的key 然后覆盖值
		if bucket.Key == key {
			bucket.Data = newValue
			break
		}
		if bucket.NextBucket == nil {
			break
		}
		bucket = bucket.NextBucket
	}
}

func (hashtable *Hashtable) Get(key string) interface{} {
	mapKey := HashingKey(key)
	node := FindKeyChainNode(mapKey, KeyChains)
	if node == nil {
		log.Fatal("cant find node")
	}

	bucket := node.Bucket
	if bucket == nil {
		return nil
	}

	for {
		if bucket.Key == key {
			return bucket.Data
		}
		if bucket.NextBucket == nil {
			return nil
		}
		bucket = bucket.NextBucket
	}
}
