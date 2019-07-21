# Chapter 8: Goroutines and Channels

* FTP client/server exercise is not complete.
* When sending on an unbuffered channel it will wait until the other end receives.
* When sending on an buffered channel of size 1 the sending goroutine will not block it will continue as normal. But will block on the next send.
* time.After and time.Tick are generating channels to be used as timers.
* time.NewTicker() is more convenient than time.Tick()
