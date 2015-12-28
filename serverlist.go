/*
	serverlist is a small cli script to look up for ec2 instances using tags on them
	It is an example for how to filter ec2 instances using tags on the ec2 instance
	This serverlist looks for tags Environment and Role and filter them according to the
	Value of the Environment and Role tag.
	It also filter and gets only the instances in runnning of pending state

	It looks up for awskey, secret and region from environment variable

		lyambem-MBP:go-intro lyambem$ ./serverlist -e dev -r api
		+-----------+------------------------------------------+------------+------------+
		|   NAME    |                 DNSNAME                  | INSTANCEID |    ZONE    |
		+-----------+------------------------------------------+------------+------------+
		| pprod-api | ec2-88-99-77-999.compute-1.amazonaws.com | i-271kk100 | us-east-1d |
		+-----------+------------------------------------------+------------+------------+

*/

package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type info struct {
	Name        string
	MachineType string
	InstanceId  string
	//SpotId         string
	Zone    string
	DnsName string
}

func drawtable(data []info) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "DnsName", "InstanceId", "Zone"})

	for _, d := range data {

		table.Append([]string{d.Name, d.DnsName, d.InstanceId, d.Zone})
	}
	table.Render()

}

func main() {

	env := flag.String("e", "", "Environment tag to look up")
	role := flag.String("r", "", "Role tag to look up")

	flag.Parse()

	if (*env == "") || (*role == "") {
		fmt.Println("Environment and Role names are needed. Exitting ... ")

	}

	ec2client := ec2.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("tag:Role"),
				Values: []*string{
					aws.String(*role),
				},
			},
			&ec2.Filter{
				Name: aws.String("tag:Environment"),
				Values: []*string{
					aws.String(*env),
				},
			},
			&ec2.Filter{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}

	resp, err := ec2client.DescribeInstances(params)

	check(err)

	var instances []info

	for idx, _ := range resp.Reservations {

		for _, inst := range resp.Reservations[idx].Instances {

			name := *inst.InstanceId

			//spotid := "None"

			for _, keys := range inst.Tags {
				if *keys.Key == "Name" {
					name = *keys.Value
				}

			}
			instance := info{
				Name:        name,
				MachineType: *inst.InstanceType,
				InstanceId:  *inst.InstanceId,
				Zone:        *inst.Placement.AvailabilityZone,
				//SpotId: *inst.SpotInstanceRequestId,
				DnsName: *inst.PublicDnsName,
			}

			instances = append(instances, instance)

		}
	}
	//fmt.Println(instances)
	drawtable(instances)

}
