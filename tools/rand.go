package tools

import (
	"crypto/rand"
	"math/big"
)

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("最小值不能成功最大值")
	}
	//if min < 0 {
	//	f64Min := math.Abs(float64(min))
	//	i64Min := int64(f64Min)
	//	result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))
	//	return result.Int64() - i64Min
	//} else {
	//	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	//	return min + result.Int64()
	//}

	result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return min + result.Int64()

}
