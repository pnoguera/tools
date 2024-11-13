// +build tools

package main

import (
	_ "github.com/SuddenGunter/jwtdec" // JWT decoder
	// _ "github.com/anchore/grype"                               // Vulnerability scanner
	_ "github.com/cloudflare/cfssl/cmd/cfssl" // PKI/TLS swiss army nife
	_ "github.com/davecheney/httpstat"        // http status
	_ "github.com/dundee/gdu/v5/cmd/gdu"      // DU tool
	_ "github.com/elecprog/tldr"              // tldr is a man similar tool

	_ "github.com/gobuffalo/cli/cmd/buffalo"                   // Go Buffalo
	_ "github.com/haggishunk/hclfmt"                           // HCL code formatter. Supports terraform 0.12
	_ "github.com/monochromegane/the_platinum_searcher/cmd/pt" // platinum searcher
	_ "github.com/xo/usql"                                     //universal SQL client

	_ "github.com/mr-karan/doggo/cmd/doggo"  // DNS tool
	_ "github.com/parkghost/gohttpbench"     // http benchmarking tool
	_ "github.com/rakyll/hey"                // another http benchmarking tool
	_ "github.com/securego/gosec/cmd/gosec"  // Go code security inpection
	_ "github.com/six-ddc/plow"              // HTTP benchmark
	_ "github.com/tomwright/dasel/cmd/dasel" // jq for JSON,YAML and others

	// _ "github.com/tsenart/vegeta"                // http load tool
	_ "filippo.io/age/cmd/age"                                 // Modern encryption tool
	_ "filippo.io/age/cmd/age-keygen"                          // Modern encryption tool
	_ "github.com/awslabs/eks-node-viewer/cmd/eks-node-viewer" // viasualization of dynamic node usage in EKS

	// _ "github.com/v-byte-cpu/sx"                               // Network scanner
	_ "github.com/wallix/awless"                     // aws cli wrapper
	_ "github.com/warrensbox/terraform-switcher"     // Terraform version switch tool
	_ "github.com/yannh/kubeconform/cmd/kubeconform" // kubernetes manifest validation tool
	_ "github.com/ycd/dstp/cmd/dstp"                 // common network tests agaist a domain
	_ "golang.org/x/vuln/cmd/govulncheck"            // GO vulnerability check
	_ "mvdan.cc/sh/v3/cmd/shfmt"                     // format shell programs
	_ "sigs.k8s.io/kind"                             // K8s in Docker
)
