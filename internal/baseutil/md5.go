package baseutil

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string)string{

	h := md5.New()
	h.Write([]byte(str))

	return fmt.Sprintf("%x",h.Sum(nil))
}
