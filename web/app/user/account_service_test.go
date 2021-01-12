package user

import "testing"

func TestMd5Password(t *testing.T) {

	pwd := md5password("seven4x")
	println(len(pwd))
	println(pwd)

}
