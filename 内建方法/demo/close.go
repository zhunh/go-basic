package demo

//close -> chan

func CloseChan()  {
	mChan := make(chan int,1)
	defer close(mChan)
	mChan <- 1
}
