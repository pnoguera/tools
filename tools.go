// +build tools

package main

import (
	_ "github.com/SuddenGunter/jwtdec"                         // JWT decoder
	_ "github.com/davecheney/httpstat"                         // http status
	_ "github.com/gjbae1212/gossm"                             // SSM SSH
	_ "github.com/gobuffalo/buffalo/buffalo"                   // Go Buffalo
	_ "github.com/haggishunk/hclfmt"                           // HCL code formatter. Supports terraform 0.12
	_ "github.com/monochromegane/the_platinum_searcher/cmd/pt" // platinum searcher
	_ "github.com/mr-karan/doggo/cmd/doggo/cli@v0.4.0"         // #name:doggo# DNS tool
	_ "github.com/parkghost/gohttpbench"                       // http benchmarking tool
	_ "github.com/rakyll/hey"                                  // another http benchmarking tool
	_ "github.com/securego/gosec/cmd/gosec"                    // Go code security inpection
	_ "github.com/tsenart/vegeta/v12"                          // http load tool
	_ "github.com/warrensbox/terraform-switcher"               // Terraform version switch tool
)
