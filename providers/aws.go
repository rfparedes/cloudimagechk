package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//GetRegions return all AWS regions
func GetRegions() []string {
	var regionSlice []string
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create a new EC2 client
	svc := ec2.New(sess)

	resultRegions, err := svc.DescribeRegions(nil)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	for _, reg := range resultRegions.Regions {
		regionSlice = append(regionSlice, *reg.RegionName)
	}
	return regionSlice
}

// GetImages return AWS images according to a filter
func GetImages(publisher string) {
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	// Create a new EC2 client
	svc := ec2.New(sess)
	input := &ec2.DescribeImagesInput{
		ImageIds: []*string{
			aws.String("ami-07174570cb87e95c5"),
		},
	}

	resultImages, _ := svc.DescribeImages(input)

	fmt.Println(resultImages)
}
