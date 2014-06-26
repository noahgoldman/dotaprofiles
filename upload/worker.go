package upload

import (
	"log"
)

type Worker struct {
	Id          int
	Work        chan *UploadRequest
	WorkerQueue chan chan *UploadRequest
	QuitChan    chan bool
}

func NewWorker(id int, workerQueue chan chan *UploadRequest) *Worker {
	return &Worker{id, make(chan *UploadRequest), workerQueue, make(chan bool)}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				UploadAndNotify(work)
			case <-w.QuitChan:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true // value doesn't matter
	}()
}

func UploadAndNotify(req *UploadRequest) {
	err := Upload_S3(req.file, req.filename)
	if err != nil {
		log.Print(err)
	}

	// 1 denotes a TextMessage
	err = req.conn.WriteMessage(1, []byte("hai"))
	if err != nil {
		log.Print(err)
	}
}
