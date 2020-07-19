package main

import (
	"flag"
	"fmt"

	"github.com/rfparedes/cloudimagechk/providers/aws"
	"github.com/rfparedes/cloudimagechk/utils"
)

func main() {

	providerPtr := flag.String("provider", "aws", "Cloud Service Provider to use")
	publisherPtr := flag.String("publisher", "SUSE", "Publisher images to use")
	providerListPtr := flag.Bool("listproviders", false, "List available providers")

	_ = providerPtr
	_ = publisherPtr

	flag.Parse()

	if *providerListPtr == true {
		for _, v := range utils.GetProviderList() {
			fmt.Println(v)
		}
	}

	aws.GetImages("SUSE")
}
