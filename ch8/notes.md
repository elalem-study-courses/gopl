# Chapter 8 Notes

* To create a channel that sends events periodically we can use
  ```golang
    time.Tick(1 * time.Second)
  ```
*  The time.After function immediately returns a channel, and starts a new goroutine that sends a single value on that channel after the specified time.
