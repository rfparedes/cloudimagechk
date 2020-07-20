package main

import (
	"flag"
	"fmt"

	"github.com/rfparedes/cloudimagechk/providers/aws"
	"github.com/rfparedes/cloudimagechk/utils"
)

func main() {

	providerPtr := flag.String("provider", "none", "Cloud Service Provider to use")
	publisherPtr := flag.String("publisher", "SUSE", "Publisher images to use")
	providerListPtr := flag.Bool("listproviders", false, "List available providers")
	flag.Parse()
	if *providerListPtr == true {
		for _, v := range utils.GetProviderList() {
			fmt.Println(v)
		}
	}
	// AWS
	if *providerPtr == "aws" {
		var (
			imageInfo  []aws.Image
			regionInfo []string
		)
		regionInfo = aws.GetRegions()
		for _, reg := range regionInfo {
			imageInfo = aws.GetImages(*publisherPtr, reg)
			fmt.Println(imageInfo)
		}
	}
}
