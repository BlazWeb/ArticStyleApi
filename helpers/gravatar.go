// Script By ArticDev
package helpers

import (
	"crypto/md5"
	"fmt"
)

func RegisterAvatar(email string) string {
	hash := md5.Sum([]byte(email))
	idefault := "identicon"
	return fmt.Sprintf("https://www.gravatar.com/avatar/%x?d=%s", hash, idefault)
}
