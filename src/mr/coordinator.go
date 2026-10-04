package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type TaskStatus string

const (
	NotStarted TaskStatus = "not_started"
	InProgress TaskStatus = "in_progress"
	Done       TaskStatus = "done"
)

type Task struct {
	Status   TaskStatus
	Filename string
}

type Coordinator struct {
	mapTasks []Task
	nReduce  int
}

// Your code here -- RPC handlers for the worker to call.

// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
func (coordinator *Coordinator) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}

// RequestTask handles a worker's request for its next unit of work.
func (coordinator *Coordinator) RequestTask(args *RequestTaskArgs, reply *RequestTaskReply) error {
	reply.TaskType = WaitTask
	return nil
}

// start a thread that listens for RPCs from worker.go
func (coordinator *Coordinator) server(sockname string) {
	rpc.Register(coordinator)
	rpc.HandleHTTP()
	os.Remove(sockname)
	listener, err := net.Listen("unix", sockname)
	if err != nil {
		log.Fatalf("listen error %s: %v", sockname, err)
	}
	go http.Serve(listener, nil)
}

// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
func (coordinator *Coordinator) Done() bool {
	ret := false

	// Your code here.


	return ret
}

// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
func MakeCoordinator(sockname string, files []string, nReduce int) *Coordinator {
	coordinator := Coordinator{
		nReduce: nReduce,
	}

	for _, filename := range files {
		task := Task{
			Status:   NotStarted,
			Filename: filename,
		}

		coordinator.mapTasks = append(coordinator.mapTasks, task)
	}

	coordinator.server(sockname)
	return &coordinator
}
