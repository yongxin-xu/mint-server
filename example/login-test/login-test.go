package login_test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	mintcommon "mint-server/common"
	handler "mint-server/handler"
	"testing"
)

func TestUserProtocString(t *testing.T) {
	au := handler.NewUser("yx123", "yongxin", "yongxin123")
	buffer, _ := proto.Marshal(au)
	t.Log(buffer)
	data := &mintcommon.AppUser{}
	err := proto.Unmarshal(buffer, data)
	if err != nil {
		fmt.Println(err)
	}
	t.Log(data)
}