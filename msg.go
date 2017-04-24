package paxi

import "fmt"

/**************************
 *    Interface Related   *
 **************************/

type Message interface{}

/***************************
 * Client-Replica Messages *
 ***************************/

type CommandID int

type Request struct {
	CommandID CommandID
	Command   Command
	ClientID  ID
	Timestamp int64
}

func (p Request) String() string {
	return fmt.Sprintf("Request {cid=%d, cmd=%v, id=%s}", p.CommandID, p.Command, p.ClientID)
}

type Reply struct {
	OK        bool
	CommandID CommandID
	LeaderID  ID
	ClientID  ID
	Command   Command
	Timestamp int64
}

func (r Reply) String() string {
	return fmt.Sprintf("Reply {ok=%t, cid=%d, lid=%s, id=%v, cmd=%v}", r.OK, r.CommandID, r.LeaderID, r.ClientID, r.Command)
}

type Read struct {
	CommandID CommandID
	Key       Key
}

type ReadReply struct {
	CommandID CommandID
	Value     Value
}

/**************************
 *     Config Related     *
 **************************/

type EndpointType uint8

const (
	CLIENT EndpointType = iota
	NODE
)

type Register struct {
	EndpointType EndpointType
	ID           ID
	Addr         string
}
