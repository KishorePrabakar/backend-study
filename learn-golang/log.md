## 1. What is a Go-Routine?

a. Concurrency vs. Parallelism:

Concurrency(Design) - The ability to break a program down into independent units of work (called Goroutines in Go) that can be executed out of order without affecting the final outcome.
It is a design structure.
On a single-core CPU, Go handles this efficiently by multiplexing thousands of lightweight Goroutines onto a small number of OS threads.
The Go runtime scheduler actively monitors which Goroutines are running, waiting, or blocked (e.g., waiting for I/O), and switches between them in nanoseconds.
This gives the illusion that many tasks are running at once without needing a separate OS thread for every single task.
While highly cost-efficient on a single core, a concurrent design is not restricted to one core—it is simply built so that it can handle running on one or many.

Parallelism(Execution) -The simultaneous execution of multiple tasks at the exact same literal millisecond.
Instead of rapidly switching between tasks on one CPU core, tasks are physically distributed across multiple CPU cores to compute data faster.
This is an execution property that depends entirely on having multi-core hardware.

(Note on Go, Java, and Python): Go automatically scales your concurrent design into true parallelism when run on a multi-core machine.
Java also utilizes true parallelism via OS threads and Virtual Threads.
Python can achieve parallelism, but standard Python (CPython) is uniquely limited by its Global Interpreter Lock (GIL), meaning it must use separate OS processes rather than threads to achieve true multi-core CPU parallelism.

    * Go Runtime Scheduler - Responsible to concurrency of mutliple GoRoutines. It used M:N scheduling model, where N GoRoutines are scheduled across M OS Threads.

b. Channels:

Channels are conduits (pipes) in which GoRoutines pass their data, in order to communicate with each other.
A Channel is initialized with the "chan" keyword.
They are strictly typed, meaning - only one data type can pass through the channel.

##### <u>Random task - 1:</u> Building a function that fetches from 3 URLs concurrently and merges results [`Go-routines and channels`]
