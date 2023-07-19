package validations

import (
	"fmt"
)

const (
	singleAZCount = 1
	MultiAZCount  = 3
)

func MinReplicasValidator(minReplicas int, multiAZ bool, isHostedCP bool, privateSubnetsCount int) error {
	if minReplicas < 0 {
		return fmt.Errorf("min-replica must be greater than zero")
	}
	if isHostedCP {
		// This value should be validated in a previous step when checking the subnets
		if privateSubnetsCount < 1 {
			return fmt.Errorf("Hosted clusters require at least a private subnet")
		}

		if minReplicas%privateSubnetsCount != 0 {
			return fmt.Errorf("Hosted clusters require that the number of compute nodes be a multiple of "+
				"the number of private subnets %d, instead received: %d", privateSubnetsCount, minReplicas)
		}
		return nil
	}

	if multiAZ {
		if minReplicas < 3 {
			return fmt.Errorf("Multi AZ cluster requires at least 3 compute nodes")
		}
		if minReplicas%3 != 0 {
			return fmt.Errorf("Multi AZ clusters require that the number of compute nodes be a multiple of 3")
		}
	} else if minReplicas < 2 {
		return fmt.Errorf("Cluster requires at least 2 compute nodes")
	}
	return nil
}

func MaxReplicasValidator(minReplicas int, maxReplicas int, multiAZ bool, isHostedCP bool, privateSubnetsCount int) error {
	if minReplicas > maxReplicas {
		return fmt.Errorf("max-replicas must be greater or equal to min-replicas")
	}

	if isHostedCP {
		if maxReplicas%privateSubnetsCount != 0 {
			return fmt.Errorf("Hosted clusters require that the number of compute nodes be a multiple of "+
				"the number of private subnets %d, instead received: %d", privateSubnetsCount, maxReplicas)
		}
		return nil
	}

	if multiAZ && maxReplicas%3 != 0 {
		return fmt.Errorf("Multi AZ clusters require that the number of compute nodes be a multiple of 3")
	}
	return nil
}

func ValidateAvailabilityZonesCount(multiAZ bool, availabilityZonesCount int) error {
	if multiAZ && availabilityZonesCount != MultiAZCount {
		return fmt.Errorf("The number of availability zones for a multi AZ cluster should be %d, "+
			"instead received: %d", MultiAZCount, availabilityZonesCount)
	}
	if !multiAZ && availabilityZonesCount != singleAZCount {
		return fmt.Errorf("The number of availability zones for a single AZ cluster should be %d, "+
			"instead received: %d", singleAZCount, availabilityZonesCount)
	}

	return nil
}
