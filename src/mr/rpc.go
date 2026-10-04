package mr

//
// RPC definitions.
//
// RPC definitions used for communication between workers and the coordinator.
// Workers request tasks from the coordinator and receive task assignments in response.
// These types define the shared request/response contract for MapReduce coordination.
// Additional RPC messages can be added here as the worker-coordinator protocol grows.
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

