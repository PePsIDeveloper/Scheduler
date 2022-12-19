# Scheduler

**Usage**

    func main() {
		s := NewScheduler()
	
		s.Schedule(func() {
			fmt.Println("Hello, world!")
		}, time.Second)
		s.Run()
	}

***
