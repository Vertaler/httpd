package main

import (
	"./handler"
	"./tcp_server"
	"flag"
	"runtime"
	"time"
	"fmt"
)

func main() {
	cpunum := *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	address := *flag.String("a", "localhost:8080", "-a HOST:PORT")
	root_dir := *flag.String("r", "../../http-test-suite-master", "-r ROOTDIR")
	log_enable := *flag.Bool("l", true, "-l")
	factory := handler.CreateFactory(root_dir, log_enable)
	server := tcp_server.CreateServer(cpunum, address, factory)
	server.Start()
	timer := time.NewTicker(time.Second)
	go func(){
		for{
			<- timer.C
			fmt.Println(runtime.NumGoroutine())
		}
	}()
}
