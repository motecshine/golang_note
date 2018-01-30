package hashtable

import "hash/fnv"


type Bucket struct {
	Data       interface{}
	NextBucket *Bucket //  解决hash冲突
}

type KeyChainNode struct {
	Key              int
	NextKeyChainNode *KeyChainNode
	Bucket           *Bucket
	Refcount         int //计算冲突量有多少 默认是0
}

func strHashToBucket(str string) uint32 {
	hashModel := fnv.New32a()
	hashModel.Write([]byte(str))
	return hashModel.Sum32() % 8
}
