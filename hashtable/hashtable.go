package hashtable

import (
	"hash/fnv"
)

const (
	ChainNodeSize uint32 = 8 // 假设Key单项列表默认的长度是8个Node
)

var KeyChains *KeyChainNode

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

type Hashtable struct{}

func HashingKey(str string) uint32 {
	hashModel := fnv.New32a()
	hashModel.Write([]byte(str))
	return hashModel.Sum32() % ChainNodeSize
}

func init() {
	KeyChains = InitKeyChains()
}

/* 初始化一个长度为ChainNodeSize的node */
func InitKeyChains() *KeyChainNode {
	var prevNode, tmpNode *KeyChainNode
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

	mapKey := HashingKey(key)

	node := FindKeyChainNode(mapKey, InitKeyChains())
	// 不冲突
	if node.Refcount == 0 {
		node.Bucket = &Bucket{
			Key:        key,
			Data:       value,
			NextBucket: nil,
		}
	} else {
		//
		for {
			if node.Bucket.NextBucket == nil {
				node.Bucket.NextBucket = &Bucket{
					Key:        key,
					Data:       value,
					NextBucket: nil,
				}
				break
			}
			node.Bucket = node.Bucket.NextBucket
		}
	}
	// 当前node的bucket个数
	node.Refcount++
}

func (hashtable *Hashtable) Delete(key string, value interface{}) interface{} {
	return nil
}

func (hashtable *Hashtable) Update(key string, newValue interface{}) interface{} {
	return nil
}

func (hashtable *Hashtable) Get(key string) interface{} {
	return nil
}
