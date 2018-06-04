package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	loggingv1 "github.com/syucream/protobuf-sandbox/src/logging/v1"
	loggingv2 "github.com/syucream/protobuf-sandbox/src/logging/v2"
)

var (
	ts       = ptypes.TimestampNow()
	v1struct = &loggingv1.Event{
		Id:        1,
		CreatedAt: ts,
		EventType: loggingv1.Event_TWEET,
		UserId:    1,
		Value:     "test tweet!",
	}
	v1_5struct = &loggingv1.RichEvent{
		Id:        1,
		CreatedAt: ts,
		EventType: loggingv1.RichEvent_TWEET,
		UserId:    1,
		Value:     "test tweet!",
		Url:       "https://www.example.com/",
		ReplyTo:   1,
	}
	v2struct = &loggingv2.Event{
		Id:          1,
		EventAt:     ts,
		ProcessedAt: ts,
		EventType:   loggingv2.Event_TYPE_TWEET,
		EventSource: loggingv2.Event_SOURCE_PUBLIC_TIMELINE,
		UserId:      1,
		UserAgent:   loggingv2.Event_UA_IOS,
		Value:       "test tweet!",
	}
)

func checkSize() {
	// v1 binary
	fmt.Printf("len(v1bin) : %d\n", proto.Size(v1struct))

	// v1 binary wrapped as Any
	v1any, err := ptypes.MarshalAny(v1struct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("len(v1anyBin) : %d\n", proto.Size(v1any))
}

func handleBetweenV1AndV1_5() {
	// Sender marshales message as v1
	bin, err := proto.Marshal(v1struct)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver unmarshales message as v1.5
	var unmarshaled loggingv1.RichEvent
	if err := proto.Unmarshal(bin, &unmarshaled); err != nil {
		fmt.Println("Couldn't to unmarshale")
	}
	fmt.Println(&unmarshaled)
}

func handleBetweenV1_5AndV1() {
	// Sender marshales message as v1.5
	bin, err := proto.Marshal(v1_5struct)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver unmarshales message as v1
	var unmarshaled loggingv1.Event
	if err := proto.Unmarshal(bin, &unmarshaled); err != nil {
		fmt.Println("Couldn't to unmarshale")
	}
	fmt.Println(&unmarshaled)
}

func handleBetweenV1AndV2() {
	// Sender marshales message as v1
	bin, err := proto.Marshal(v1struct)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver unmarshales message as v2
	var unmarshaled loggingv2.Event
	if err := proto.Unmarshal(bin, &unmarshaled); err != nil {
		fmt.Println("Couldn't to unmarshale")
	}
	fmt.Println(&unmarshaled)
}

func handleBetweenV2AndV1() {
	// Sender marshales message as v2
	bin, err := proto.Marshal(v2struct)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver unmarshales message as v2
	var unmarshaled loggingv1.Event
	if err := proto.Unmarshal(bin, &unmarshaled); err != nil {
		fmt.Println("Couldn't to unmarshale")
	}
	fmt.Println(&unmarshaled)
}

func handleAny() {
	// Sender marshales message as Any
	v2any, err := ptypes.MarshalAny(v2struct)
	if err != nil {
		log.Fatal(err)
	}
	bin, err := proto.Marshal(v2any)
	if err != nil {
		log.Fatal(err)
	}

	// Receiver unmarshales message as Any
	var unmarshaled any.Any
	if err := proto.Unmarshal(bin, &unmarshaled); err != nil {
		fmt.Println("Couldn't to unmarshale")
	}
	fmt.Println(&unmarshaled)

	// Receiver unmarshales message as v1
	var v1event loggingv1.Event
	if ptypes.Is(&unmarshaled, &v1event) {
		if err := ptypes.UnmarshalAny(&unmarshaled, &v1event); err != nil {
			fmt.Println("Couldn't to unmarshale")
		}
		fmt.Println(&v1event)
	} else {
		fmt.Println("Any is not loggingv1.Event")
	}

	// Receiver unmarshales message as v1.5
	var v1_5event loggingv1.RichEvent
	if ptypes.Is(&unmarshaled, &v1_5event) {
		if err := ptypes.UnmarshalAny(&unmarshaled, &v1_5event); err != nil {
			fmt.Println("Couldn't to unmarshale")
		}
		fmt.Println(&v1_5event)
	} else {
		fmt.Println("Any is not loggingv1.RichEvent")
	}

	// Receiver unmarshales message as v2
	var v2event loggingv2.Event
	if ptypes.Is(&unmarshaled, &v2event) {
		if err := ptypes.UnmarshalAny(&unmarshaled, &v2event); err != nil {
			fmt.Println("Couldn't to unmarshale")
		}
		fmt.Println(&v2event)
	} else {
		fmt.Println("Any is not loggingv2.Event")
	}
}

func main() {
	fmt.Println("Check handling between messages might be compatible : ")
	fmt.Println("v1 -> v1.5")
	handleBetweenV1AndV1_5()
	fmt.Println("v1.5 -> v1")
	handleBetweenV1_5AndV1()
	fmt.Println("")

	fmt.Println("Check handling between messages might be NOT compatible : ")
	fmt.Println("v1 -> v2")
	handleBetweenV1AndV2()
	fmt.Println("v2 -> v1")
	handleBetweenV2AndV1()
	fmt.Println("")

	fmt.Println("Check handling Any : ")
	handleAny()
	fmt.Println("")

	fmt.Println("Check marshaled size w/ Any : ")
	checkSize()
}
