package contract

import (
	"math/rand"
	"time"
)

var managerKSPool, validatorKSPool RoundRobin

type KSItem struct {
	Key    string
	Secret string
}

var managerKS = []*KSItem{}
var validatorKS = []*KSItem{}

// func init() {
// 	managerKS = shuffle(managerKS)
// 	validatorKS = shuffle(validatorKS)
// 	managerKSPool, _ = NewRoundRobin(managerKS)
// 	validatorKSPool, _ = NewRoundRobin(validatorKS)
// }

func shuffle(vals []*KSItem) []*KSItem {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]*KSItem, len(vals))
	n := len(vals)
	for i := 0; i < n; i++ {
		randIndex := r.Intn(len(vals))
		ret[i] = vals[randIndex]
		vals = append(vals[:randIndex], vals[randIndex+1:]...)
	}
	return ret
}

func GetManagerKS() ([]byte, string) {
	item := managerKSPool.Next()
	return []byte(item.Key), item.Secret
}

func GetValidatorKS() ([]byte, string) {
	item := validatorKSPool.Next()
	return []byte(item.Key), item.Secret
}
