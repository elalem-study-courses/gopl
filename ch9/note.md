# Chapter 9 Concurrency with shared variables

* First way to avoid race condition is making variables **immutable** which is hard if we want to make updates.
* Second way to avoid race condition is to avoid accessing variables from multiple goroutines. This variables are **confined** to a single goroutine. We can communicate these variables using channels this is what *Go mantra* says **Do not communicate by sharing memory; instead, share memory by communicating**.
* Serial confinment is that multiple goroutines share some variables by communicating them and these variables are confined within each of these goroutines (imagine it as a number of stages) each stage is confined and variables sent to the next stage are never accessed again with previous stages.
* The third way of synchronization is **mutual exclusion** that is each shared variable is accessed only once at a time.
* By convention the variables guarded by a mutex is defined below this mutex.
