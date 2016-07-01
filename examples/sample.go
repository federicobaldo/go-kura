package main

import (
	"fmt"
	pb "github.com/federicobaldo/go-kura/kuradatatypes"
	"github.com/golang/protobuf/jsonpb"
)

func main() {

	metric_name := "mymetric"

	metric_type := pb.KuraPayload_KuraMetric_STRING

	metric := &pb.KuraPayload_KuraMetric{
		Name: &metric_name,
		Type: &metric_type,
	}

	metrics := []*pb.KuraPayload_KuraMetric{metric}

	message := &pb.KuraPayload{
		Metric: metrics,
	}

	fmt.Printf("Here my Kura payload as string: %s\n", message)

	marshaler := &jsonpb.Marshaler{}
	json_string, _ := marshaler.MarshalToString(message)

	fmt.Printf("Here my Kura payload as json string: %s\n", json_string)
}
