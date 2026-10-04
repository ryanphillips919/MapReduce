package mr

//
// RPC definitions.
//
// remember to capitalize all names.
//

type TaskType string

const (
	MapTask    TaskType = "map"
	ReduceTask TaskType = "reduce"
	WaitTask   TaskType = "wait"
	ExitTask   TaskType = "exit"
)

type RequestTaskArgs struct {
}

type RequestTaskReply struct {
	TaskType TaskType
	TaskID   int
	Filename string
	NReduce  int
}

//
// example to show how to declare the arguments
// and reply for an RPC.
//

type ExampleArgs struct {
	X int
}

type ExampleReply struct {
	Y int
}

// Add your RPC definitions here.

