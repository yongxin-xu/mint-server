package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	handler "mint-server/handler"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("Client start error", err)
	}

	resp := &handler.SrvResponse{}

	time.Sleep(time.Second * 1)

	for {
		au := handler.NewUser("yongxin123", "yongxin", "12345678")
		rq := handler.WrapAccountRequest(au, "SignUp")
		//
		//buffer, _ := proto.Marshal(rq)
		//fmt.Println(buffer)
		//
		//if _, err := conn.Write(buffer); err != nil {
		//      fmt.Println("Write error", err)
		//}
		//
		//buf := make([]byte, 1024)
		//cnt, err := conn.Read(buf)
		//if err != nil {
		//      fmt.Println("Read error", err)
		//}
		//
		//if err := proto.Unmarshal(buf[:cnt], resp); err != nil {
		//      fmt.Println("Unmarshal error", err)
		//} else {
		//      fmt.Println(resp.Info, resp.Au.ID)
		//}
		//
		//time.Sleep(time.Second * 1)

		//au = handler.NewUser("yx123f", "yongxin", "yongxin123")
		//rq = handler.WrapAccountRequest(au, "SignUp")
		//
		//buffer, _ = proto.Marshal(rq)
		//fmt.Println(buffer)
		//
		//if _, err := conn.Write(buffer); err != nil {
		//      fmt.Println("Write error", err)
		//}
		//
		//buf = make([]byte, 1024)
		//cnt2, err := conn.Read(buf)
		//if err != nil {
		//      fmt.Println("Read error", err)
		//}
		//
		//if err := proto.Unmarshal(buf[:cnt2], resp); err != nil {
		//      fmt.Println("Unmarshal error", err)
		//} else {
		//      fmt.Println(resp.Info, resp.Au.Name, resp.Au.Account)
		//}
		//
		//time.Sleep(time.Second * 1)
		//
		//au = handler.NewUser("yx123", "yongxin", "yongxin123")
		//rq = handler.WrapAccountRequest(au, "SignUpf")
		//
		//buffer, _ = proto.Marshal(rq)
		//fmt.Println(buffer)
		//
		//if _, err := conn.Write(buffer); err != nil {
		//      fmt.Println("Write error", err)
		//}
		//
		//buf = make([]byte, 1024)
		//cnt3, err := conn.Read(buf)
		//if err != nil {
		//      fmt.Println("Read error", err)
		//}
		//
		//if err := proto.Unmarshal(buf[:cnt3], resp); err != nil {
		//      fmt.Println("Unmarshal error", err)
		//} else {
		//      fmt.Println(resp.Info, resp.Au.Name, resp.Au.Account)
		//}

		//time.Sleep(time.Second * 1)

		rq = handler.WrapAccountRequest(au, "SignIn")

		buffer2, _ := proto.Marshal(rq)
		fmt.Println(buffer2)

		if _, err := conn.Write(buffer2); err != nil {
			fmt.Println("Write error", err)
		}

		buf2 := make([]byte, 1024)
		cnt4, err := conn.Read(buf2)
		if err != nil {
			fmt.Println("Read error", err)
		}

		if err := proto.Unmarshal(buf2[:cnt4], resp); err != nil {
			fmt.Println("Unmarshal error", err)
		} else {
			fmt.Println(resp.Info, resp.Au.Name, resp.Au.Account)
		}

		time.Sleep(time.Second * 1)

		//rq = handler.WrapAccountRequest(au, "SignIn")
		//
		//buffer2, _ = proto.Marshal(rq)
		//fmt.Println(buffer2)
		//
		//if _, err := conn.Write(buffer2); err != nil {
		//      fmt.Println("Write error", err)
		//}
		//
		//buf2 = make([]byte, 1024)
		//cnt5, err := conn.Read(buf2)
		//if err != nil {
		//      fmt.Println("Read error", err)
		//}
		//
		//if err := proto.Unmarshal(buf2[:cnt5], resp); err != nil {
		//      fmt.Println("Unmarshal error", err)
		//} else {
		//      fmt.Println(resp.Info, resp.Au.Name, resp.Au.Account)
		//}
		//
		//time.Sleep(time.Second * 1)

		break
	}

}