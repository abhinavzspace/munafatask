package awsutil

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// Basic parameters for working with aws spot instances
type Parameters struct {
	Count int
	Instance string
	Price string
	Region string
	Zone string
}


// Implements Stringer to pretty print Parameters
func (p Parameters) String() string {
	return fmt.Sprintf("count=%d, instance=%s, price=%s, region=%s, zone=%s",
		p.Count, p.Instance, p.Price, p.Region, p.Zone)
}


// It finds Spot price history by zone and instance
// Other input parameters can be configured by uncommenting the appropriate portion of the code
func (p *Parameters) DescribeSpotPriceHistory(svc *ec2.EC2) (string, error) {
	paramsForReq := &ec2.DescribeSpotPriceHistoryInput{
		AvailabilityZone: aws.String(p.Zone),
		//		DryRun:           aws.Bool(true),
		//		EndTime:          aws.Time(time.Now()),
		//		Filters: []*ec2.Filter{
		//			{ // Required
		//				Name: aws.String("String"),
		//				Values: []*string{
		//					aws.String("String"), // Required
		//					// More values...
		//				},
		//			},
		//			// More values...
		//		},
		InstanceTypes: []*string{
			aws.String(p.Instance), // Required
			// More values...
		},
		//		MaxResults: aws.Int64(1),
		//		NextToken:  aws.String("String"),
		//		ProductDescriptions: []*string{
		//			aws.String("String"), // Required
		//			// More values...
		//		},
		//		StartTime: aws.Time(time.Now()),
	}
	resp, err := svc.DescribeSpotPriceHistory(paramsForReq)

	if err != nil {
		return "", err
	}
	return fmt.Sprintln(resp), nil
}


// Creates a Spot instance request by using Price, imageId, instance and zone.
// Other input parameters can be configured by uncommenting the appropriate portion of the code
func (p *Parameters) RequestSpotInstances(svc *ec2.EC2, imageId string) (string, error) {
	paramsForReq := &ec2.RequestSpotInstancesInput{
		SpotPrice:             aws.String(p.Price), // Required
		//		AvailabilityZoneGroup: aws.String(p.Zone),
		//		BlockDurationMinutes:  aws.Int64(1),
		//		ClientToken:           aws.String("String"),
		//		DryRun:                aws.Bool(true),
		//		InstanceCount:         aws.Int64(1),
		//		LaunchGroup:           aws.String("String"),
		LaunchSpecification: &ec2.RequestSpotLaunchSpecification{
			//			AddressingType: aws.String("String"),
			//			BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			//				{// Required
			//					DeviceName: aws.String("String"),
			//					Ebs: &ec2.EbsBlockDevice{
			//						DeleteOnTermination: aws.Bool(true),
			//						Encrypted:           aws.Bool(true),
			//						Iops:                aws.Int64(1),
			//						SnapshotId:          aws.String("String"),
			//						VolumeSize:          aws.Int64(1),
			//						VolumeType:          aws.String("VolumeType"),
			//					},
			//					NoDevice:    aws.String("String"),
			//					VirtualName: aws.String("String"),
			//				},
			//				// More values...
			//			},
			//			EbsOptimized: aws.Bool(true),
			//			IamInstanceProfile: &ec2.IamInstanceProfileSpecification{
			//				Arn:  aws.String("String"),
			//				Name: aws.String("String"),
			//			},
			ImageId:      aws.String(imageId),
			InstanceType: aws.String(p.Instance),
			//			KernelId:     aws.String("String"),
			//			KeyName:      aws.String("String"),
			//			Monitoring: &ec2.RunInstancesMonitoringEnabled{
			//				Enabled: aws.Bool(true), // Required
			//			},
			//			NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			//				{// Required
			//					AssociatePublicIpAddress: aws.Bool(true),
			//					DeleteOnTermination:      aws.Bool(true),
			//					Description:              aws.String("String"),
			//					DeviceIndex:              aws.Int64(1),
			//					Groups: []*string{
			//						aws.String("String"), // Required
			//						// More values...
			//					},
			//					NetworkInterfaceId: aws.String("String"),
			//					PrivateIpAddress:   aws.String("String"),
			//					PrivateIpAddresses: []*ec2.PrivateIpAddressSpecification{
			//						{// Required
			//							PrivateIpAddress: aws.String("String"), // Required
			//							Primary:          aws.Bool(true),
			//						},
			//						// More values...
			//					},
			//					SecondaryPrivateIpAddressCount: aws.Int64(1),
			//					SubnetId:                       aws.String("String"),
			//				},
			//				// More values...
			//			},
			Placement: &ec2.SpotPlacement{
				AvailabilityZone: aws.String(p.Zone),
				//				GroupName:        aws.String("String"),
			},
			//			RamdiskId: aws.String("String"),
			//			SecurityGroupIds: []*string{
			//				aws.String("String"), // Required
			//				// More values...
			//			},
			//			SecurityGroups: []*string{
			//				aws.String("String"), // Required
			//				// More values...
			//			},
			//			SubnetId: aws.String("String"),
			//			UserData: aws.String("String"),
		},
		//		Type:       aws.String(p.Instance),
		//		ValidFrom:  aws.Time(time.Now()),
		//		ValidUntil: aws.Time(time.Now()),
	}
	resp, err := svc.RequestSpotInstances(paramsForReq)

	if err != nil {
		return "", err
	}
	return fmt.Sprintln(resp), nil
}


// Describes all available Spot instance requests for the user
// Other input parameters can be configured by uncommenting the appropriate portion of the code
func (p *Parameters) DescribeSpotInstanceRequests(svc *ec2.EC2) (string, error) {
	//	paramsforReq := &ec2.DescribeSpotInstanceRequestsInput{
	//		DryRun: aws.Bool(true),
	//		Filters: []*ec2.Filter{
	//			{ // Required
	//				Name: aws.String("String"),
	//				Values: []*string{
	//					aws.String("String"), // Required
	//					// More values...
	//				},
	//			},
	//			// More values...
	//		},
	//		SpotInstanceRequestIds: []*string{
	//			aws.String("String"), // Required
	//			// More values...
	//		},
	//	}
	resp, err := svc.DescribeSpotInstanceRequests(nil)

	if err != nil {
		return "", err
	}
	return fmt.Sprintln(resp), nil
}


// Cancels Spot instance requests for given slice of spotInstanceRequestIds
// Other input parameters can be configured by uncommenting the appropriate portion of the code
func (p *Parameters) CancelSpotInstanceRequests(svc *ec2.EC2, spotInstanceRequestIds []*string) (string, error) {
	paramsForReq := &ec2.CancelSpotInstanceRequestsInput{
		SpotInstanceRequestIds: spotInstanceRequestIds,
		//		DryRun: aws.Bool(true),
	}
	resp, err := svc.CancelSpotInstanceRequests(paramsForReq)

	if err != nil {
		return "", err
	}
	return fmt.Sprintln(resp), nil
}