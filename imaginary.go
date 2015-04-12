package main

import (
	"flag"
	"fmt"
	. "github.com/tj/go-debug"
	"os"
	"runtime"
	"strconv"
)

var debug = Debug("imaginary")

var (
	aAddr  = flag.String("a", "", "bind address")
	aPort  = flag.Int("p", 8088, "port to listen")
	aVers  = flag.Bool("v", false, "")
	aVersl = flag.Bool("version", false, "")
	aHelp  = flag.Bool("h", false, "")
	aHelpl = flag.Bool("help", false, "")
	aCors  = flag.Bool("cors", false, "")
	aGzip  = flag.Bool("gzip", false, "")
	aKey   = flag.String("key", "", "")
	aCpus  = flag.Int("cpus", runtime.GOMAXPROCS(-1), "")
)

const usage = `imaginary server %s

Usage:
  imaginary -p 80
  imaginary -cors -gzip
  imaginary -h | -help
  imaginary -v | -version

Options:
  -a <addr>     bind address [default: *]
  -p <port>     bind port [default: 8088]
  -h, -help     output help
  -v, -version  output version
  -cors         Enable CORS support [default: false]
  -gzip         Enable gzip compression [default: false]
  -key <key>    Define API key
  -cpus <num>   Number of used cpu cores.
                (default for current machine is %d cores)
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage, Version, runtime.NumCPU()))
	}
	flag.Parse()

	if *aHelp || *aHelpl {
		showUsage()
	}

	if *aVers || *aVersl {
		fmt.Println(Version)
		return
	}

	runtime.GOMAXPROCS(*aCpus)

	port := getPort(*aPort)
	opts := ServerOptions{
		Port:    port,
		Address: *aAddr,
		Gzip:    *aGzip,
		CORS:    *aCors,
		ApiKey:  *aKey,
	}

	debug("imaginary server listening on port %d", port)

	err := Server(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot start the server: %s\n", err)
		os.Exit(1)
	}
}

func getPort(port int) int {
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		newPort, _ := strconv.Atoi(portEnv)
		if newPort > 0 {
			port = newPort
		}
	}
	return port
}

func showUsage() {
	flag.Usage()
	os.Exit(1)
}