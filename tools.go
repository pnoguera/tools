// +build tools

package main

import (
	_ "github.com/davecheney/httpstat"                         // http status
	_ "github.com/monochromegane/the_platinum_searcher/cmd/pt" // platinum searcher
	_ "github.com/parkghost/gohttpbench"                       // http benchmarking tool
	_ "github.com/tsenart/vegeta"                              // http load tool
)
