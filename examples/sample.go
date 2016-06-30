package main

import (
    pb "github.com/federicobaldo/go-kura/kuradatatypes"
    "fmt"
    )

func main(){

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

    fmt.Printf("Here my Kura payload: %s\n", message)
}
