# go-aws-utils
Few aws tools/scripts in go

# serverlist.go

Looks up for instances in ec2 using tags. This is an example to look up for instance using tags whose Keys are Called Environment and Role. Serverlist list all the ec2 instances that Have the same value for Tag Keys Environment and Role

		lyambem-MBP:go-intro lyambem$ ./serverlist -e dev -r api
		+-----------+------------------------------------------+------------+------------+
		|   NAME    |                 DNSNAME                  | INSTANCEID |    ZONE    |
		+-----------+------------------------------------------+------------+------------+
		| pprod-api | ec2-88-99-77-999.compute-1.amazonaws.com | i-271kk100 | us-east-1d |
		+-----------+------------------------------------------+------------+------------+

