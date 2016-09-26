package main

import (
	"./handler"
	"./tcp_server"
	"flag"
	"fmt"
	"runtime"
)

func main() {
	cpunum := flag.Int("c", runtime.NumCPU(), "-c NCPU")
	port := flag.String("p", "80", "-p PORT")
	host := flag.String("h", "localhost", "-h HOST")
	root_dir := flag.String("r", "../../http-test-suite-master", "-r ROOTDIR")
	log_enable := flag.Bool("l", false, "-l")
	flag.Parse()
	fmt.Println("Num cpu: ", *cpunum)
	runtime.GOMAXPROCS(*cpunum)
	factory := handler.CreateFactory(*root_dir, *log_enable)
	server := tcp_server.CreateServer(*port, *host, factory)
	server.Start()
}
