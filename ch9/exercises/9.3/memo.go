package memo

type request struct {
	key      string
	response chan<- result
}

type Memo struct {
	requests chan request
}

type Func func(string, chan<- struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		cls := make(chan struct{})

		if e == nil {
			e := &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, cls)
		}
		go func(req request) {
			select {
			case <-e.ready:
				go e.deliver(req.response)
			case <-cls:
				<-e.ready
				delete(cache, req.key)
			}
		}(req)
	}
}

func (e *entry) call(f Func, key string, cls chan<- struct{}) {
	e.res.value, e.res.err = f(key, cls)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	response <- e.res
}
