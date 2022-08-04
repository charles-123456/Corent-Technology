package main

import (
	"flag"
	"fmt"
	"github.com/kardianos/service"
	"net/http"
	"os"
	"strings"
)

var logger service.Logger

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {

	if service.Interactive() {
		_ = logger.Info("Running in terminal.")
	} else {
		_ = logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {
	http.HandleFunc("/", hi)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/status", Status)
	http.HandleFunc("/service", Service)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
	}
}

func hi(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "corent-service test")
}

func (p *program) Stop(s service.Service) error {

	return nil
}
func Status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Running SuccessFully.......!")
}
func Service(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "azure app service....!!!!")
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func main() {

	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()
	svcConfig := &service.Config{
		Name:        "corent-service-test",
		DisplayName: "corent-service-test",
		Description: "corent-service-test",
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
