package protocalRpc

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"testing"
)

func TestProtoc(t *testing.T) {
	//创建实例
	p := Person{
		Id:     1234,
		Name:   "XWJ",
		Email:  "xwj@stu,pku.edu.cn",
		Phones: []*Person_PhoneNumber{{Number: "123456", Type: Person_HOME}}}
	// 使用 proto 库的 Marshal 函数来序列换protocol buffer 数据
	out, err := proto.Marshal(&p)
	if err != nil {
		log.Fatalln("编码失败：", err)
	}
	fname := "./protoc_text.txt"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("写入失败:", err)
	} else {
		//	读取数据
		in, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Fatalln("读取失败:", err)
		}
		person := &Person{}
		if err := proto.Unmarshal(in, person); err != nil {
			log.Fatalln("解析失败:", err)
		} else {
			fmt.Println(person)
		}
	}
}
