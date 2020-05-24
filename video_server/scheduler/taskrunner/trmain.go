package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}
//定时任务
func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for{
		select {
		case <- w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start()  {
	//开始视频文件清理
	r := NewRunner(3, true,VideoClearDispatcher,VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()

	//something else
	//r1 :=
}