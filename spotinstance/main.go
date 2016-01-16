package main

import (
    "fmt"
    "flag"
	"os"
	"strings"

	"github.com/abhinavzspace/munafatask/awsutil"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)


func main() {

	var params awsutil.Parameters

	flag.IntVar(&params.Count, "count", 1, "This is used for the number of instance requests you need")
	flag.StringVar(&params.Instance, "instance", "t1.micro", "This Represents the Instance Type")
	flag.StringVar(&params.Price, "price", "0", "This Represents the Price to use for the Requested instance")
	flag.StringVar(&params.Region, "region", "us-east-1", "This Represents the Region to Use")
	flag.StringVar(&params.Zone, "zone", "us-east-1c", "This Represents the zone to use for the respective Region")

	flag.Parse()

    fmt.Println(params)
    // Create an EC2 service object in the "us-west-2" region
    // Note that you can also configure your region globally by
    // exporting the AWS_REGION environment variable
    svc := ec2.New(session.New(), &aws.Config{Region: aws.String(params.Region)})

	var taskNo int
	var tryAgainMessage = "You can try again!"

	for {
		fmt.Println("Select one of the following:")
		fmt.Println("1. Get the price history")
		fmt.Println("2. Create the Spot Instance Request")
		fmt.Println("3. Check the Spot Instance Request History")
		fmt.Println("4. Cancel the Spot Instance Request")
		fmt.Println("5. Exit")

		fmt.Scanf("%d\n", &taskNo)

		switch taskNo {
		case 1:
			resp, err := params.DescribeSpotPriceHistory(svc)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(tryAgainMessage)
				break
			}
			fmt.Println(resp)
		case 2:
			fmt.Println("Please enter ImageId, eg. ami-f0091d91")
			var imageId string
			fmt.Scanf("%s\n", &imageId)
			resp, err := params.RequestSpotInstances(svc, imageId);
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(tryAgainMessage)
				break
			}
			fmt.Println(resp)
		case 3:
			resp, err := params.DescribeSpotInstanceRequests(svc)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(tryAgainMessage)
				break
			}
			fmt.Println(resp)
		case 4:
			fmt.Println("Enter spot instance request ids delimited by ',' which are needed to be canceled")
			var idsString string
			var ids []*string
			fmt.Scanf("%s\n", &idsString)
			idsString = strings.Trim(idsString, ", ")		// Trim ',' and ' '
			for _, v := range strings.Split(idsString, ",") {
				ids = append(ids, aws.String(v))
			}
			resp, err := params.CancelSpotInstanceRequests(svc, ids)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println(tryAgainMessage)
				break
			}
			fmt.Println(resp)
		case 5:
			fmt.Println("Thanks for checking out aws spot instances. Now exiting!")
			os.Exit(0)
		default:
			fmt.Println("Please enter a digit between 1 to 5")
		}
	}

}