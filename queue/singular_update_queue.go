package queue

type UpdateResponse struct {
	Err error
	Val interface{}
}

type Updater func(request interface{}) UpdateResponse

type updateRequest struct {
	request         interface{}
	responseChannel chan UpdateResponse
}

type SingularUpdateQueue struct {
	requestChannel chan updateRequest
	update Updater
}

func NewSingularUpdateQueue(size int, updater Updater) *SingularUpdateQueue {
	queue := &SingularUpdateQueue{
		requestChannel: make(chan updateRequest, size),
		update: updater,
	}
	go queue.singularUpdate()
	return queue
}

func (queue *SingularUpdateQueue) Push(request interface{}) {
	queue.requestChannel <- updateRequest{
		request:         request,
		responseChannel: nil,
	}
}

func (queue *SingularUpdateQueue) PushForResponse(request interface{}) UpdateResponse {
	req := updateRequest{
		request:         request,
		responseChannel: make(chan UpdateResponse),
	}
	go queue.push(req)
	resp := <- req.responseChannel
	return resp
}

func (queue *SingularUpdateQueue) push(req updateRequest) {
	queue.requestChannel <- req
}

func (queue *SingularUpdateQueue) singularUpdate() {
	for {
		select {
		case req := <-queue.requestChannel:
			resp := queue.update(req.request)
			if req.responseChannel != nil {
				req.responseChannel <- resp
			}
		}
	}
}
