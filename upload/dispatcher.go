package upload

var WorkerQueue chan chan *UploadRequest

var WorkQueue chan *UploadRequest

func StartDispatcher(n int) {
	WorkerQueue = make(chan chan *UploadRequest, n)

	for i := 0; i < n; i++ {
		worker := NewWorker(i, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <- WorkQueue:
				go func() {
					worker := <- WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}
