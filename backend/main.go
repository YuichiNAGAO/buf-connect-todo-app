package main

import (
	"context"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	taskv1 "github.com/YuichiNAGAO/buf-connect-todo-app/backend/gen/proto/task/v1"
	taskv1connect "github.com/YuichiNAGAO/buf-connect-todo-app/backend/gen/proto/task/v1/taskv1connect"
)

type Server struct{}

func (s *Server) GetTaskList(
	ctx context.Context,
	req *connect.Request[taskv1.GetTaskListRequest],
) (*connect.Response[taskv1.GetTaskListResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&taskv1.GetTaskListResponse{
		Tasks: []*taskv1.Task{
			{
				Id:   "1",
				Name: "Task 1",
			},
			{
				Id:   "2",
				Name: "Task 2",
			},
		},
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	server := &Server{}
	mux := http.NewServeMux()
	path, handler := taskv1connect.NewTaskServiceHandler(server)
	log.Println("Serving TaskService at", path)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
