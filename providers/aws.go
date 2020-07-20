package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//Image type
type Image struct {
	CreationDate string
	Name         string
	ImageID      string
}

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
func GetImages(publisher string, region string) []Image {

	var (
		owners     = []*string{aws.String(publisher)}
		imageSlice []Image
	)
	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(region)},
	}))
	// Create a new EC2 client
	svc := ec2.New(sess)
	input := &ec2.DescribeImagesInput{
		Owners: owners,
	}
	resultImages, _ := svc.DescribeImages(input)
	for _, image := range resultImages.Images {
		imageSlice = append(imageSlice, Image{*image.CreationDate, *image.Name, *image.ImageId})
	}
	return imageSlice
}
