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

func TestRequestTaskAssignments(t *testing.T) {
    c := Coordinator{
        mapTasks: []Task{
            {Status: NotStarted, Filename: "file1.txt"},
            {Status: NotStarted, Filename: "file2.txt"},
            {Status: NotStarted, Filename: "file3.txt"},
        },
        nReduce: 3,
    }

    args := RequestTaskArgs{}

    for i := range 4 {
        reply := RequestTaskReply{}

        err := c.RequestTask(&args, &reply)
        if err != nil {
            t.Fatalf("RequestTask returned error: %v", err)
        }

        if i == 3 {
            // All three tasks have been assigned.
            if reply.TaskType != WaitTask {
                t.Errorf("Request %d: got %q, want %q",
                    i, reply.TaskType, WaitTask)
            }
            continue
        }

        if reply.TaskType != MapTask {
            t.Errorf("Request %d: got %q, want %q",
                i, reply.TaskType, MapTask)
        }

        if reply.TaskID != i {
            t.Errorf("Request %d: got task ID %d, want %d",
                i, reply.TaskID, i)
        }

        if c.mapTasks[i].Status != InProgress {
            t.Errorf("Task %d: status is %q, want %q",
                i, c.mapTasks[i].Status, InProgress)
        }
    }
}
