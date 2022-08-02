package main

import (
	"fmt"
	"time"
	"strings"
	"github.com/magiconair/properties"
	"github.com/mozilla/mig/modules/netstat"
	"corent-go/google_chat_check"
	"golang.org/x/exp/slices"
)
	

var PropertyFile = []string{"./conf.properties"}
var P, _ = properties.LoadFiles(PropertyFile, properties.UTF8, true)
var TomcatName = P.MustGet("tomcat.name")
var TomcatPort = P.MustGet("tomcat.port")
// var ServiceAccPath= P.MustGet("service.account.path")
var ChatSpaceName= P.MustGet("chat.space.name")
// var StoppedPort = []string{}
// var Time = time.Now()
// var Dt_fmt = Time.Format("01-02-2006 15:04:05")
var ActivePort[]string
var DeadPort[]string
// var isAllow = false


func main() {
	for range time.Tick(time.Second * 10){
		TomcatPort := strings.Split(TomcatPort,",")
		TomcatName := strings.Split(TomcatName,",")
		for i,port := range TomcatPort {
			NeverStop(port,TomcatName[i])
		}		
			}

}


func NeverStop(port string,Name string) {
	conn,_, _ := netstat.HasListeningPort(port)
	if conn{
		IsActivePort := slices.Contains(ActivePort,port)
		if !IsActivePort{
			data := fmt.Sprintf("%v is Started",Name)
			google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName)
			ActivePort = append(ActivePort,port)
		}
		RemoveDeadPort()
	}else{
		IsDeadPort := slices.Contains(DeadPort,port)
		if !IsDeadPort{
			data := fmt.Sprintf("%v is stopped",Name)
			google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName)
			DeadPort = append(DeadPort,port)
		}
		RemoveActivePort()
	}
}


func RemoveActivePort(){
	common, _ := intersection(ActivePort,DeadPort)
	for i,_ := range common{
		for j,val := range ActivePort{
			if common[i] == val{
				ActivePort = RemoveIndex(ActivePort,j)
			}

		}
	}
}



func RemoveDeadPort(){
	common, _ := intersection(ActivePort,DeadPort)
	for i,_ := range common{
		for j,val := range DeadPort{
			if common[i] == val{
				 DeadPort = RemoveIndex(DeadPort,j)
			}
		}
	}

}


func intersection(a, b []string) ([]string, error) {

	// uses empty struct (0 bytes) for map values.
	m := make(map[string]struct{}, len(b))

	// cached
	for _, v := range b {
	  m[v] = struct{}{}
	}

	var s []string
	for _, v := range a {
	  if _, ok := m[v]; ok {
		s = append(s, v)
	  }
	}

	return s, nil
  }

func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}