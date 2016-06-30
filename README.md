# go-kura
Some utilities to parse Eclipse Kura payloads in Go.

## Kura Datatype
The kuradatatypes package contains the protobuf generated Go code for the Kura payload
as defined in the Kura repository:

https://github.com/eclipse/kura/blob/develop/kura/org.eclipse.kura.core.cloud/src/main/protobuf/kurapayload.proto

In order to use it you can import it in your Go program using:

```go
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
```