Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency ensures that you can handle different task at the same time (or it looks like it, even though it actually is going back and forth really fast). Parallelism is the execution in different CPUs or Cores making the execution of the task actually at the same time. 

What is the difference between a *race condition* and a *data race*? 
> A race condition is when the output of the exection depends on the timing of the threads, its essentially just a logical failure. This includes any type of operation. A data race is a type of race condition where different threads try to acsess the same shared data memory. Making the output unreliable because it can vary each time only depending on the finishing time and the synchronization between the threads. This only includes the operation of memory acsessing. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler ensures that the synchronization between threads are in tact. It coordinates the execution of multiple tasks making it the essence of a OS. It basically decides when what is happening in what order. What gets to use the CPU, in what order and context switching.  


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> The point of using multiple threads in programming is to achieve better performance, responsiveness, and efficient utilization of system resources by enabling a program to execute multiple tasks concurrently. 

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> They are managed by the language runtime rather than the OS, that includes the context switching. We have to use cooperative multitasking for them to work. A coroutine explicitly "yields" execution back to the scheduler, allowing another coroutine to run. Unlike threads (which are preemptive), coroutines don't require synchronization primitives like mutexes because only one coroutine executes at a time within a single thread. Fibers are light weight, taking up less memory. Efficient context switching because it doesnt interact with the OS on that.


Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Concurrency can make a programmer's life both easier and harder. It simplifies handling multiple tasks simultaneously, improving responsiveness, resource utilization, and scalability, especially for I/O-bound or naturally concurrent problems. However, it adds complexity due to race conditions, deadlocks, debugging challenges, and non-deterministic behavior.

What do you think is best - *shared variables* or *message passing*?
> Message passing is generally better than shared variables for most concurrent programming tasks because it avoids race conditions, simplifies synchronization, and scales well across distributed systems. However, shared variables can be better for simple, performance-critical tasks on a single machine where synchronization is easy to manage. 



