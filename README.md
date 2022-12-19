# Scheduler
A scheduler is a software component that is in charge of scheduling tasks to be completed at a predetermined time or interval. Backups, system upgrades, and other sorts of maintenance chores are often managed and executed by schedulers in computer systems. Schedulers may be written in a number of programming languages and generally include functionality such as the ability to plan jobs to run at certain times, perform tasks on a repeating schedule, and handle task dependencies. Schedulers are a crucial component of many software systems and may be utilised in a wide range of applications, from basic command-line programmes to sophisticated distributed systems.


**Usage**

    func main() {
		s := NewScheduler()
	
		s.Schedule(func() {
			fmt.Println("Hello, world!")
		}, time.Second)
		s.Run()
	}

***
