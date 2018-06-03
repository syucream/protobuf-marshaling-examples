package main

import (
	"fmt"

	"github.com/golang/protobuf/ptypes"
	loggingv1 "github.com/syucream/protobuf-sandbox/src/logging/v1"
)

func main() {
	ts := ptypes.TimestampNow()

	v1event := loggingv1.Event{
		Id:        1,
		CreatedAt: ts,
		EventType: loggingv1.Event_TWEET,
		UserId:    1,
		Value:     "test tweet!",
	}
	fmt.Println(v1event)
}
