package main

import "fmt"
type Server struct{
	Ip string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip: ip,
		Port: port,
	}
	return server
}

//启动服务器的接口
func (this *Server)start(){
	//socket listen
	listener,err := net.Listen("tcp", fmt.Sprintf("%s: %d",this.Ip,this.Port))
	if
}