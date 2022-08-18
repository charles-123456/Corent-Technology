package main

import (
	"fmt"
	"strings"
	"corent-go/google_chat_check"
	"golang.org/x/exp/slices"
	"sync"
	"github.com/kardianos/service"
	"os"
	"corent-go/corent/log"
	"flag"
	"github.com/magiconair/properties"
	"net/http"
	"path/filepath"
	"net"
	"github.com/mozilla/mig/modules/netstat"
)

// "github.com/magiconair/properties"
var TomcatName,TomcatPort,ChatSpaceName,ServiceAccPath string

type program struct{}

var(
	CurrentPath,_ = os.Getwd()
	ConfPath = CurrentPath+"\\conf.properties"
	ActivePort[]string
	DeadPort[]string
	wg sync.WaitGroup
	isFirst = true
	logger service.Logger
	PropertyFile[]string
)

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func readprops(){
	// TomcatName = "Analyzer Tomcat1,Ceapi Tomcat2,MarketPlace Tomcat3,Surpaas Tomcat4"
	// TomcatPort ="8080,9444,1455,9433"
	// ChatSpaceName="AAAAwlgqHZg"
	P, err := properties.LoadFiles(PropertyFile, properties.UTF8, true)
	if err != nil{
		log.Error(err)
	}
	TomcatName = P.MustGet("tomcat.name")
	// log.Info("TomcatName"+TomcatName)
	TomcatPort = P.MustGet("tomcat.port")
	// log.Info("TomcatPort"+TomcatPort)
	ServiceAccPath = P.MustGet("service.account.path")
	// log.Info("ServiceAccount"+ServiceAccPath)
	ChatSpaceName = P.MustGet("chat.space.name")
	// log.Info("Space"+ChatSpaceName)
	
}

func (p *program) run() {
	// Do work here
	readprops()
	value :="Hi Team 👋👋👋\nThis is Charlie😎,\n I'm hired by SaaSDev team🏣\nTo monitor Supaas Server Status🧐👁️‍🗨️."
	data := fmt.Sprintf("%v",value)
	google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName,ServiceAccPath)
	for {
		// log.Info("Entered Infinite time for loop!!!")
		TomcatPort := strings.Split(TomcatPort,",")
		TomcatName := strings.Split(TomcatName,",")
		for i,port := range TomcatPort {
			go NeverStop(port,TomcatName[i])
			wg.Add(1)
		}
		go http.ListenAndServe(":8090", nil)
		wg.Wait()
		isFirst =false
			}
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
// func openLogFile(path string)(*os.File,error){
// 	 LogFile,err := os.OpenFile(path,os.O_WRONLY | os.O_APPEND | os.O_CREATE,0644)
// 	 if err != nil {
//         return nil, err
//     }
// 	return LogFile, nil
// }

// func InfoLog(msg string){
// 	fileInfo, err := openLogFile("./Log.log")
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	infoLog := log.New(fileInfo, "[info]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
//     infoLog.Printf("%v",msg)
// }

// func errorLog(msg error){
// 	fileError,err := openLogFile("./Log.log")
// 	if err != nil {
//         log.Fatal(err)
//     }
// 	errorLog := log.New(fileError, "[error]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
// 	errorLog.Printf("%v",msg)
// }
func getHomePath() string {
	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	log.Info("BinPath"+binPath)
	if err != nil {
		panic(err)
	 }
	path, _ := filepath.Split(filepath.ToSlash(binPath))
	path = path[:len(path)-1]
	path = filepath.ToSlash(path)
	log.Info("Path"+path)
	return binPath
	}
	
	

func main() {
	// path := getHomePath()
	ConfPath := getHomePath()+"\\conf.properties"
	log.Info("COnf"+ConfPath)
	PropertyFile = []string{ConfPath}
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()
	svcConfig := &service.Config{
		Name:"Charlie2",
		DisplayName: "Charlie2",
		Description: "Charlie2",
		Arguments:[]string{ConfPath},
	}
	log.Info("New svcConfigline executed.")
	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Error(err)
	}
	log.Info("New Service creating line excuted")
	errs := make(chan error, 5)
	logger, err = s.Logger(errs)
	if err != nil {
		log.Error(err)
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
			log.Error(err)
			if strings.Contains(err.Error(), "Unknown action") {
				_, _ = fmt.Fprintf(os.Stderr, "Valid actions: %q\n", service.ControlAction)
			}
			os.Exit(1)
		}
		return
	}
	err = s.Run()
	if err != nil {
		log.Error(err)
	}
	log.Info("New Run Called Kardinos excuted")
	
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
	// log.Info("NeverStop Method Called!!!")
	ipAddr := findRemoteIP()
	conn,_, _ := netstat.HasListeningPort(port)
	pharse := StartOrRunningUpdate(isFirst)
	if conn{
		IsActivePort := slices.Contains(ActivePort,port)
		if !IsActivePort{
			data := fmt.Sprintf("%v is %v by this Host:%v",Name,pharse,ipAddr)
			google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName,ServiceAccPath)
			// log.Info("google chat method called!!!")
			ActivePort = append(ActivePort,port)
		}
		RemoveDeadPort()
	}else{
		IsDeadPort := slices.Contains(DeadPort,port)
		if !IsDeadPort{
			data := fmt.Sprintf("%v is stopped by this Host:%v",Name,ipAddr)
			google_chat_check.StartingPoint(map[string]string{"data": data},ChatSpaceName,ServiceAccPath)
			// log.Info("google chat method called!!!")
			DeadPort = append(DeadPort,port)
		}
		RemoveActivePort()
	}
}

func findRemoteIP()string {
	_,ele,_ := netstat.HasListeningPort("3399")
	ip := findLocalIp()
	var ipAddr string;
	for _,e := range ele{
		val := fmt.Sprintf("%v",e)
		isContained := strings.Contains(val,ip)
		if isContained{
			getfullIp := strings.Split(val,"3399")
			getIpwithspaces := getfullIp[len(getfullIp)-1]
			removeSpace := strings.Split(getIpwithspaces," ")
			ipAddr = removeSpace[1]
		}
	}
	return ipAddr

}

func findLocalIp() string {
	var ip string;
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
		
	}
	return ip
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