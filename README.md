# buf-connect-todo-app

#### 参考
- https://zenn.dev/7oh/articles/3b944f9b744932
- https://connectrpc.com/docs/go/getting-started/
- https://takanamito.hateblo.jp/entry/2024/12/02/095349
- https://www.awarefy.dev/blog/better-grpc-connect-go/ 


#### 動作確認

```shell
curl \
    --header "Content-Type: application/json" \
    --data '{}' \              
    http://localhost:8080/proto.task.v1.TaskService/GetTaskList
```