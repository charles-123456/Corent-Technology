package main

import (
	"fmt"
	"strings"
	"github.com/magiconair/properties"
	"github.com/mozilla/mig/modules/netstat"
	"corent-go/google_chat_check"
	"golang.org/x/exp/slices"
	"sync"
	"github.com/kardianos/service"
	"flag"
	"os"
	"time"
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
var wg sync.WaitGroup
var isFirst = true

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	value :="Hi Team 👋👋👋\nThis is Charlie😎,\n I'm hired by SaaSDev team🏣\nTo monitor Supaas Server Status🧐👁️‍🗨️."
	data := fmt.Sprintf("%v",value)
	google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName)
	for range time.Tick(time.Second * 10)  {
		TomcatPort := strings.Split(TomcatPort,",")
		TomcatName := strings.Split(TomcatName,",")
		for i,port := range TomcatPort {
			NeverStop(port,TomcatName[i])
		}
		isFirst =false
			}

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()
	svcConfig := &service.Config{
		Name:"Charlie1",
		DisplayName: "Charlie1",
		Description: "Charlie1",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		panic(err)
	}
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
	go func() {
		for {
			err := <-errs
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "Unknown action") {
				_, _ = fmt.Fprintf(os.Stderr, "Valid actions: %q\n", service.ControlAction)
			}
			os.Exit(1)
		}
		return
	}
	err = s.Run()
	if err != nil {
		_ = logger.Error(err)
	}
	
}

func StartOrRunningUpdate(isFirst bool)string{
	if isFirst{
		val := "Running"
		return val
	}else{
		val := "Started"
		return val
	}
}

func NeverStop(port string,Name string) {
	// defer wg.Done()
	conn,_, _ := netstat.HasListeningPort(port)
	pharse := StartOrRunningUpdate(isFirst)
	if conn{
		IsActivePort := slices.Contains(ActivePort,port)
		if !IsActivePort{
			data := fmt.Sprintf("%v is %v",Name,pharse)
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