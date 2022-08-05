package main

import (
	"fmt"
	"strings"
	"github.com/mozilla/mig/modules/netstat"
	"corent-go/google_chat_check"
	"golang.org/x/exp/slices"
	"sync"
	"github.com/kardianos/service"
	"flag"
	"os"
	"time"
	"log"
)
	
// "github.com/magiconair/properties"
var TomcatName,TomcatPort,ChatSpaceName string

// var StoppedPort = []string{}
// var Time = time.Now()
// var Dt_fmt = Time.Format("01-02-2006 15:04:05")
var ActivePort[]string
var DeadPort[]string
// var isAllow = false
var wg sync.WaitGroup
var isFirst = true

func readprops(){
	TomcatName = "Analyzer Tomcat1,Ceapi Tomcat2,MarketPlace Tomcat3,Surpaas Tomcat4"
	TomcatPort ="8080,9444,1455,9433"
	ChatSpaceName="AAAAwlgqHZg"
	// var PropertyFile = []string{"./conf.properties"}
	// var P, _ = properties.LoadFiles(PropertyFile, properties.UTF8, true)
	//  TomcatName = P.MustGet("tomcat.name")
	//  TomcatPort = P.MustGet("tomcat.port")
	// var ServiceAccPath= P.MustGet("service.account.path")
	//  ChatSpaceName= P.MustGet("chat.space.name")
	
}


var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	InfoLog("Entered Start service method")
	if service.Interactive() {
		_ = logger.Info("Running in terminal.")
	} else {
		_ = logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	InfoLog("Now Calling run method")
	go p.run()
	return nil
}

func (p *program) run() {
	readprops()
	InfoLog("Entered run method")
	value :="Hi Team 👋👋👋\nThis is Charlie😎,\n I'm hired by SaaSDev team🏣\nTo monitor Supaas Server Status🧐👁️‍🗨️."
	data := fmt.Sprintf("%v",value)
	google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName)
	for range time.Tick(time.Second * 10)  {
		InfoLog("Entered Infinite time for loop!!!")
		TomcatPort := strings.Split(TomcatPort,",")
		TomcatName := strings.Split(TomcatName,",")
		for i,port := range TomcatPort {
			go NeverStop(port,TomcatName[i])
			wg.Add(1)
		}
		isFirst =false
		wg.Wait()	
			}
	InfoLog("Finshed run method")

}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	InfoLog("Now Calling stop method")
	return nil
}

func openLogFile(path string)(*os.File,error){
	 LogFile,err := os.OpenFile(path,os.O_WRONLY | os.O_APPEND | os.O_CREATE,0644)
	 if err != nil {
        return nil, err
    }
	return LogFile, nil
}

func InfoLog(msg string){
	fileInfo, err := openLogFile("./Log.log")
	if err != nil {
        log.Fatal(err)
    }
	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
    infoLog.Printf("%v",msg)
}

func errorLog(msg error){
	fileError,err := openLogFile("./Log.log")
	if err != nil {
        log.Fatal(err)
    }
	errorLog := log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	errorLog.Printf("%v",msg)
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()
	svcConfig := &service.Config{
		Name:"Charlie2",
		DisplayName: "Charlie2",
		Description: "Charlie2",
	}
	InfoLog("New svcConfigline executed.")
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		errorLog(err)
	}
	InfoLog("New Service creating line excuted")
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		errorLog(err)
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
			errorLog(err)
			if strings.Contains(err.Error(), "Unknown action") {
				_, _ = fmt.Fprintf(os.Stderr, "Valid actions: %q\n", service.ControlAction)
			}
			os.Exit(1)
		}
		return
	}
	err = s.Run()
	if err != nil {
		errorLog(err)
	}
	InfoLog("New Run Called Kardinos excuted")
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
	defer wg.Done()
	InfoLog("NeverStop Method Called!!!")
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