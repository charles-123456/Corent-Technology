package main

import (
   "os/exec"
	"log"
)

func main(){
	
cmd := exec.Command("corent-go.exe")
if err := cmd.Start(); err != nil{
    log.Fatal(err)
}
}