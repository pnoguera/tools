// +build tools

package main

import (
	_ "github.com/SuddenGunter/jwtdec"                         // JWT decoder
	_ "github.com/anchore/grype"                               // Vulnerability scanner
	_ "github.com/cloudflare/cfssl/cmd/cfssl"                  // PKI/TLS swiss army nife
	_ "github.com/davecheney/httpstat"                         // http status
	_ "github.com/dundee/gdu/v5/cmd/gdu"                       // DU tool
	_ "github.com/elecprog/tldr"                               // tldr is a man similar tool
	_ "github.com/gjbae1212/gossm"                             // SSM SSH
	_ "github.com/gobuffalo/cli/cmd/buffalo@v0.17.7"           // Go Buffalo
	_ "github.com/haggishunk/hclfmt"                           // HCL code formatter. Supports terraform 0.12
	_ "github.com/monochromegane/the_platinum_searcher/cmd/pt" // platinum searcher
	_ "github.com/mr-karan/doggo/cmd/doggo/cli@v0.4.0"         // #name:doggo# DNS tool
	_ "github.com/parkghost/gohttpbench"                       // http benchmarking tool
	_ "github.com/rakyll/hey"                                  // another http benchmarking tool
	_ "github.com/securego/gosec/cmd/gosec"                    // Go code security inpection
	_ "github.com/six-ddc/plow@latest"                         // HTTP benchmark
	_ "github.com/tsenart/vegeta/v12"                          // http load tool
	_ "github.com/warrensbox/terraform-switcher"               // Terraform version switch tool
	_ "github.com/ycd/dstp/cmd/dstp"                           // common network tests agaist a domain
	_ "mvdan.cc/sh/v3/cmd/shfmt@latest"                        // format shell programs
	_ "sigs.k8s.io/kind@v0.11.1"                               // K8s in Docker
)
