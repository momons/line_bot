package Util

import (
	"crypto/sha512"
	"fmt"
)

// ハッシュユーティル.
type HashUtil struct {
}

// ハッシュユーティルインスタンス.
var hashUtilInstance *HashUtil

// インスタンス取得.
func GetHashUtil() *HashUtil {
	if hashUtilInstance == nil {
		// インスタンス作成.
		hashUtilInstance = &HashUtil{}
	}
	return hashUtilInstance
}

func (util *HashUtil) Hash(
	header string,
	center string,
	footer string,
) string {
	hashStr := center
	for i := 0; i < 10; i++ {
		hashStr = fmt.Sprintf("%x", sha512.Sum512([]byte(header+center+footer)))
	}
	return hashStr
}
