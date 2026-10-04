package mr

import "testing"

// These are small tests for the worker/coordinator RPC contract.
func TestRequestTaskReply(t *testing.T) {
	reply := RequestTaskReply{
		TaskType: MapTask,
		TaskID:   1,
		Filename: "pg-test.txt",
		NReduce:  3,
	}

	if reply.TaskType != MapTask {
		t.Errorf("TaskType = %q, want %q", reply.TaskType, MapTask)
	}
	if reply.TaskID != 1 {
		t.Errorf("TaskID = %d, want %d", reply.TaskID, 1)
	}
	if reply.Filename != "pg-test.txt" {
		t.Errorf("Filename = %q, want %q", reply.Filename, "pg-test.txt")
	}
	if reply.NReduce != 3 {
		t.Errorf("NReduce = %d, want %d", reply.NReduce, 3)
	}
}

func TestRequestTaskReturnsWait(t *testing.T) {
	c := Coordinator{}
	args := RequestTaskArgs{}
	reply := RequestTaskReply{}

	err := c.RequestTask(&args, &reply)

	if err != nil {
		t.Fatalf("RequestTask returned error: %v", err)
	}

	if reply.TaskType != WaitTask {
		t.Errorf("TaskType = %q, want %q", reply.TaskType, WaitTask)
	}
}
