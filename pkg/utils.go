package pkg

import (
	"encoding/hex"
	"regexp"
	"strings"
	"sync"
)

var (
	utilsOnce sync.Once
	utils     *Utils
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func GetUtils() *Utils {
	utilsOnce.Do(func() {
		utils = NewUtils()
	})
	return utils
}

func (Utils) ConvertToHex(token []byte) string {
	str := hex.EncodeToString(token)
	var re = regexp.MustCompile("[^a-fA-F0-9]")
	return re.ReplaceAllString(str, "")
}

// Is element of s contains in e
func (Utils) ContainsString(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(e, a) {
			return true
		}
	}
	return false
}
