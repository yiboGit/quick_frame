/**
* @Author: yibo_LastDay
* @Date: 2022/10/15 15:41
 */

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"quick_frame/const_data"
	ylog "quick_frame/log"
)

type HttpServer struct {
	HttpGin     *gin.Engine
	HttpPort    string
	HttpAddress string
}

func NewServer() *HttpServer {
	server := &HttpServer{
		HttpGin:     gin.New(),
		HttpPort:    const_data.Port,
		HttpAddress: const_data.Address,
	}
	return server
}

func (server *HttpServer) Start() {
	server.RegisterRoute()
	ylog.Log.Infof("http server start! %v:%v \n", server.HttpAddress, server.HttpPort)
	server.HttpGin.Run(fmt.Sprintf("%s:%s", server.HttpAddress, server.HttpPort))
}

func (server *HttpServer) RegisterRoute() {
	groupStudent := server.HttpGin.Group("student")

	groupManager := server.HttpGin.Group("manager")

	groupSalesperson := server.HttpGin.Group("salesperson")
}
