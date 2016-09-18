package main

import (
	"./handler"
	"./tcp_server"
	"flag"
	"runtime"
)

func main() {
	cpunum := *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	address := *flag.String("a", "localhost:8080", "-a HOST:PORT")
	root_dir := *flag.String("r", "../../http-test-suite-master", "-r ROOTDIR")
	log_enable := *flag.Bool("l", false, "-l")
	factory := handler.CreateFactory(root_dir, log_enable)
	server := tcp_server.CreateServer(cpunum, address, factory)
	server.Start()
}
