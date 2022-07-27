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
var ServiceAccPath= P.MustGet("service.account.path")
var ChatSpaceName= P.MustGet("chat.space.name")
// var StoppedPort = []string{}
var Time = time.Now()
var Dt_fmt = Time.Format("01-02-2006 15:04:05")
var IsExistCheck = false
var ExistConnId = false
var ActivePort[]string
var DeadPort[]string
// var isAllow = false

func main() {
		for range time.Tick(time.Second * 10){
			PortListenerCheck(TomcatPort,ExistConnId)
			
		}		
}

var (
    globalMap = make(map[string]int)
)

func PortListenerCheck(Port string,ExistConnId bool) {
	TomcatPort := strings.Split(TomcatPort,",")
	TomcatName := strings.Split(TomcatName,",")
	for i,port := range TomcatPort {
		conn,_, _ := netstat.HasListeningPort(port)
		if conn {
				IsActivePort := slices.Contains(ActivePort,TomcatPort[i])
				if !IsActivePort{
					ActivePort = append(ActivePort,TomcatPort[i])
					IsFirst := true
					CheckPort(TomcatPort[i],1,TomcatName[i],IsFirst)
					RemoveDeadPort()
				}else{
					if cap(ActivePort) > 0 {
						IsActivePort :=  slices.Contains(ActivePort,TomcatPort[i])
						if !IsActivePort{
							RemoveDeadPort()
							IsFirst := false
							CheckPort(TomcatPort[i],1,TomcatName[i],IsFirst)
							ActivePort = append(ActivePort,TomcatPort[i])
						}
					}else{
						RemoveDeadPort()
						IsFirst := false
						CheckPort(TomcatPort[i],1,TomcatName[i],IsFirst)
						ActivePort = append(ActivePort,TomcatPort[i])
						// ActivePort = ActivePort[:0]
						}
				}			
				}else{
				IsDeadPort := slices.Contains(DeadPort,TomcatPort[i])
				if !IsDeadPort{
					DeadPort = append(DeadPort,TomcatPort[i])
					IsFirst := true
					CheckPort(TomcatPort[i],0,TomcatName[i],IsFirst)
					RemoveActivePort()
				}else{
				RemoveActivePort()
				IsFirst := false
				CheckPort(TomcatPort[i],0,TomcatName[i],IsFirst)
				// DeadPort = append(DeadPort,TomcatPort[i])
				// for id,_ := range DeadPort{
				// 	CheckPort(DeadPort[id],0,TomcatName[i])
				// 	// DeadPort = nil
				// }
				}
				 }
			}	
	}	

func CheckPort(port string,mapVal int,Name string,IsFirst bool){
	globalMap[port]=mapVal
	if !AlwaysAlertCon(ActivePort,DeadPort,port) || IsFirst {
		for id,IsVal := range globalMap{
			if id == port  && IsVal == 1{
				data := fmt.Sprintf("%v = %v is listening Now and Time - %v !!!\r\n",Name,port,Dt_fmt)
				google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
			}
			if id == port  && IsVal == 0{
				data := fmt.Sprintf("%v = %v is Stopped Now and Time - %v !!!\r\n",Name,port,Dt_fmt)
				google_chat_check.StartingPoint(map[string]string{"data": data},ServiceAccPath,ChatSpaceName)
			}
		}
	}
	delete(globalMap,port)
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

func AlwaysAlertCon(ActivePort,DeadPort []string,port string) bool {
	isActive := slices.Contains(ActivePort,port)
	isDead := slices.Contains(DeadPort,port)
	if isActive || isDead{
		return true
	}else{
		return false
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