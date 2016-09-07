package main

import (
	"flag"
	"runtime"
	"httpd"
)
func main() {
	 server := httpd.Server{}
	 server.cpu_number = *flag.Int("c", runtime.NumCPU(), "-c NCPU")
	 server.port = *flag.Int("p", 8080, "-p PORT")
	 server.root = *flag.String("r", "/", "-r ROOTDIR")
	 server.log_enable = *flag.Bool("e", false, "-e")
	 server.log_path = *flag.String("l","./","-l LOGPATH")
	 server.Start()
}
