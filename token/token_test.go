/*
#
# @Time : 2019/9/10 16:20
# @Author : Qian Lu
*/

package token

import (
	"fmt"
	"os"
	"testing"
)

func TestToken(t *testing.T) {
	accessKeyId, accessKeySecret := os.Getenv("accessKeyId"), os.Getenv("accessKeySecret")
	token, err := GetAliToken(accessKeyId, accessKeySecret)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		fmt.Println(fmt.Sprintf("Testing token:%s", token))
		t.Log("OK")
	}
}
