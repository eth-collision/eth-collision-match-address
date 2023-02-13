package tool

import (
	"math/big"
	"sync"
)

var locker = &sync.RWMutex{}

func BigIntAdd(a, b *big.Int) *big.Int {
	locker.Lock()
	defer locker.Unlock()
	return big.NewInt(0).Add(a, b)
}
