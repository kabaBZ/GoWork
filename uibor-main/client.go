package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const OnlineArg = "list"

type Client struct {
	ServerIp   string
	ServerPort int
	conn       net.Conn
	flag       int //当前客户端的选择模式
	Name       string
}

func NewClient(serverIp string, serverPort int) *Client {
	//创建客户端对象
	addr := fmt.Sprintf("%s:%d", serverIp, serverPort)
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       99,
		Name:       addr,
	}
	// 连接server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return nil
	}
	client.conn = conn
	return client
}

var (
	serverIp   string
	serverPort int
)

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "prot", 3000, "设置服务器端口(默认是3000)")
}
func (c *Client) Navigation() bool {
	var flag int
	fmt.Println("请选择以下数字进入相应的功能")
	fmt.Println("1.公聊模式")
	fmt.Println("2.修改用户名")
	fmt.Println("3.私聊模式")
	fmt.Println("4.查询在线用户")
	fmt.Println("0.退出聊天室")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 4 {
		c.flag = flag
		return true
	} else {
		fmt.Println("---------请输入合法范围内的数字------------")
		return false
	}
}

//处理服务端回复的消息,直接显示到标准输出既可
func (c *Client) DealResponse() {
	fmt.Println("----------------")
	//一旦有数据则copy到stdout标准输出
	io.Copy(os.Stdout, c.conn)

}

func (c *Client) Run() {

	for c.flag != 0 {
		for !c.Navigation() {
		}
		//根据不同的模式处理不同的业务
		switch c.flag {
		case 1:
			//进入公聊模式
			fmt.Println("进入 公聊 模式")
			c.PublicChat()
			break
		case 2:
			//更新用户名
			fmt.Println("进入 更新用户名 ")
			c.UpdateName()
			break
		case 3:
			//私聊模式
			fmt.Println("进入 私聊 模式")
			c.PrivateChat()
			break
		case 4:
			//查询在线用户
			fmt.Println("进入 查询当前在线用户")
			c.OnlineClient()
		}
	}
	fmt.Println("退出!")
}

func (c *Client) UpdateName() {
	fmt.Println(">>>>>>请输入更新后的用户名<<<<<<<")
	fmt.Scanln(&c.Name)
	//封装协议
	msg := "UpdateName|" + c.Name + "\n"
	c.SendMsg(msg)
}
func (c *Client) PrivateChat() {
	var msg string
	fmt.Scanln(&msg)
	c.SendMsg(msg)
}

func (c *Client) PublicChat() {
	var msg string
	fmt.Scanln(&msg)
	c.SendMsg(msg)
}

func (c *Client) OnlineClient() {
	// var msg string
	// fmt.Scanln(&msg)
	c.SendMsg(OnlineArg)
}

func (c *Client) SendMsg(msg string) {
	_, err := c.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
	}
}

func main() {

	//命令行解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>服务器连接失败")
	} else {
		fmt.Println(">>>>>服务器连接成功")
	}
	go client.DealResponse()
	//启动客户端业务
	client.Run()
	// select {}
}
