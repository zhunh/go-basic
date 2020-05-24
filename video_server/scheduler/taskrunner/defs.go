package taskrunner

const(
	READY_TO_DIPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "./videos/"
)

type controlChan chan string
//数据
type dataChan chan interface{}
//dispatcher 和 excuter
type fn func(dc dataChan) error
