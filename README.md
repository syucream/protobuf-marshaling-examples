# protobuf-marshaling-examples

```
$ go run main.go
Check handling between messages might be compatible :
v1 -> v1.5
id:1 created_at:<seconds:1528128844 nanos:914007623 > event_type:TWEET user_id:1 value:"test tweet!"
v1.5 -> v1
id:1 created_at:<seconds:1528128844 nanos:914007623 > event_type:TWEET user_id:1 value:"test tweet!"

Check handling between messages might be NOT compatible :
v1 -> v2
id:1 event_at:<seconds:1528128844 nanos:914007623 > event_type:TYPE_TWEET
v2 -> v1
id:1 created_at:<seconds:1528128844 nanos:914007623 > user_id:1

Check handling Any :
type_url:"type.googleapis.com/logging.v2.Event" value:"\010\001\022\014\010\314\312\325\330\005\020\307\314\352\263\003\032\014\010\314\312\325\330\005\020\307\314\352\263\003 \001(\0020\0018\002B\013test tweet!"
Any is not loggingv1.Event
Any is not loggingv1.RichEvent
id:1 event_at:<seconds:1528128844 nanos:914007623 > processed_at:<seconds:1528128844 nanos:914007623 > event_type:TYPE_TWEET event_source:SOURCE_PUBLIC_TIMELINE user_id:1 user_agent:UA_IOS value:"test tweet!"

Check marshaled size w/ Any :
len(v1bin) : 33
len(v1anyBin) : 73<Paste>
```
