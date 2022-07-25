package main

import (
	"fmt"
	"time"
	"github.com/magiconair/properties"
	"github.com/mozilla/mig/modules/netstat"
	"corent-go/google_chat_check"
)

var PropertyFile = []string{"./conf.properties"}
var P, _ = properties.LoadFiles(PropertyFile, properties.UTF8, true)
var TomcatName = P.MustGet("tomcat.name")
var TomcatPort = P.MustGet("tomcat.port")
var ServiceAccPath= P.MustGet("service.account.path")
var ChatSpaceName= P.MustGet("chat.space.name")
// var StoppedPort = []string{}
var Time = time.Now()
var Dt_fmt = Time.Format("01-02-2006 15:04:05")
var IsExistCheck = false
var ExistConnId = false
func main() {
		PortListenerCheck(TomcatPort,ExistConnId)
}

func PortListenerCheck(Port string,ExistConnId bool) {
	conn, _, err := netstat.HasListeningPort(Port)

	ErrorCheck(err)
	if !conn {
		if !ExistConnId {
			data := fmt.Sprintf("%v = %v is already Stopped -->time - %v \n",TomcatName,Port,Dt_fmt)
			google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
		}
		// data := fmt.Sprintf("%v = %v is alread Stopped -->time - %v \n",TomcatName,Port,Dt_fmt)
		// google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
		for range time.Tick(time.Second * 10){
			port_conn,_,_ := netstat.HasListeningPort(Port)
			if port_conn{
				IsExistCheck = true
				ExistingPortCheck(IsExistCheck)
			}
		}
		// StoppedPort = append(StoppedPort, Port)		
	} else {
		ExistingPortCheck(IsExistCheck)
	}
}

// func ExistPortCheck(val []string) {
// 	for i, ele := range val {
// 		conn, _, _ := netstat.HasListeningPort(ele)
// 		if conn {
// 			data := fmt.Sprintf(" %v = %v has been started and Time - %v!!!\n",TomcatName,ele,Dt_fmt)
// 			if !RunOnly{
// 				google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
// 				RunOnly = true
// 			}
// 			StoppedPort = RemoveIndex(val, i)
// 		}
// 	}
// }

// func RemoveIndex(s []string, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }

func ExistingPortCheck(IsExistCheck bool){
	data := fmt.Sprintf("%v = %v is listening Now and-->Time - %v !!!\n",TomcatName,TomcatPort,Dt_fmt)
		google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
		for range time.Tick(time.Second * 10){
			port_conn, _, _ := netstat.HasListeningPort(TomcatPort)
			if !port_conn{
				data := fmt.Sprintf("%v = %v is Stopped Now --> time - %v   \n",TomcatName,TomcatPort,Dt_fmt)
				google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
				if IsExistCheck{
					ExistConnId = true
					PortListenerCheck(TomcatPort,ExistConnId)
					ExistConnId = false
				} else {
					ExistConnId = true
					PortListenerCheck(TomcatPort,ExistConnId)
				}
				
			}
			
		}
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}


