[![GoDoc](https://godoc.org/github.com/federicobaldo/go-kura?status.svg)](https://godoc.org/github.com/federicobaldo/go-kura)
# go-kura
Some utilities to parse Eclipse Kura payloads in Go.

## Kura Datatype
The kuradatatypes package contains the protobuf generated Go code for the Kura payload
as defined in the Kura repository:

https://github.com/eclipse/kura/blob/develop/kura/org.eclipse.kura.core.cloud/src/main/protobuf/kurapayload.proto

In order to use the kuradatatypes package you can import it in your Go program using:

```go
package main

import (
    pb "github.com/federicobaldo/go-kura/kuradatatypes"
    "github.com/golang/protobuf/jsonpb"
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

    fmt.Printf("Here my Kura payload as string: %s\n", message)

    marshaler := &jsonpb.Marshaler{}
    json_string, _ := marshaler.MarshalToString(message)
    
    fmt.Printf("Here my Kura payload as json string: %s\n", json_string)
}
```

## Examples
You can find some example programs in the ./examples folder. 

The mqtt_subscriber example, subscribe to an mqtt topic and convert the KuraPayload to
a json string. Here the usage:

```sh
Usage of examples/mqtt_subscriber:
  -clientid string
    	A clientid for the connection (default "localhost")
  -log.level value
    	Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal, panic]. (default info)
  -password string
    	Password to match username
  -qos int
    	The QoS to subscribe to messages at
  -server string
    	The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883 (default "tcp://127.0.0.1:1883")
  -topic string
    	Topic to subscribe to (default "#")
  -username string
    	A username to authenticate to the MQTT server
```

You can build the examples using ```go build filename.go```

Detailed instructions on how to setup your development environment are here:
https://golang.org/doc/code.html
