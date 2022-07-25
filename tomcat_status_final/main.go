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
// var TomcatName = P.MustGet("tomcat.name")
var TomcatPort = P.MustGet("tomcat.port")
var StoppedPort = []string{}
var Time = time.Now()
var Dt_fmt = Time.Format("01-02-2006 15:04:05")

func main() {
	for range time.Tick(time.Second * 60) {
		if len(StoppedPort) > 0 {
			ExistPortCheck(StoppedPort)
		}
		PortListenerCheck(TomcatPort)
	}
}

func PortListenerCheck(Port string) {
	conn, _, err := netstat.HasListeningPort(Port)

	ErrorCheck(err)
	if !conn {
		data := fmt.Sprintf("%v is Stopped time - %v\n",Port,Dt_fmt)
		google_chat_check.StartingPoint(map[string]string{"data": data})
		StoppedPort = append(StoppedPort, Port)

	} else {
		data := fmt.Sprintf("%v is listening Now and Time - %v !!!\n",Port,Dt_fmt)
		google_chat_check.StartingPoint(map[string]string{"data": data})
	}
}

func ExistPortCheck(val []string) {
	for i, ele := range val {
		conn, _, _ := netstat.HasListeningPort(ele)
		if conn {
			data := fmt.Sprintf(" %v has been started and Time - %v!!!\n",ele,Dt_fmt)
			google_chat_check.StartingPoint(map[string]string{"data": data})
			StoppedPort = RemoveIndex(val, i)
		}
	}
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}


