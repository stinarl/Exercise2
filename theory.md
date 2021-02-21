Exercise 2 - Theory questions
-----------------------------

### What is an atomic operation?
> Atomic operations are operations that run completely independent of each other, not beeing aware of each other.

### What is a critical section?
> A critical section is a protected section that can only be operated by one process at a time. It is often shared sections which can lead to flaws if 
executed from concurrent processes at the same time. 


### What is the difference between race conditions and data races?
> "Race condition: A race condition is a situation, in which the result of an operation depends on the interleaving of certain individual operations.
Data race: A data race is a situation, in which at least two threads access a shared variable at the same time. At least on thread tries to modify the variable."

Copied 10.02.2021 from : https://www.modernescpp.com/index.php/race-condition-versus-data-race#:~:text=Race%20condition%3A%20A%20race%20condition,tries%20to%20modify%20the%20variable.


### What are the differences between semaphores, binary semaphores, and mutexes?
> A mutex is a locking mechanism that is used to make sure that a critical section only can be entered by one thread at the time.
A semaphore is a signal used to control access to threads such as critical sections. Counting semaphores does this by a counting integer, while binary semaphores is either zero - no access - og 1 - access.


### What are the differences between channels (in Communicating Sequential Processes, or as used by Go, Rust), mailboxes (in the Actor model, or as used by Erlang, D, Akka), and queues (as used by Python)? 
> CSP are a formal language for describing patterns og communication between concurrent processes. channels (go) serve the same purpose. A mailbox is were threads are qued while waiting to be prosecuted by the actor. In a mailbox it is possible to make priorities between the waiting threads, which meaans it serves the same purpose as a channel  or CSP.


### List some advantages of using message passing over lock-based synchronization primitives.
> In addition to just locking you can share messages.
It is easier to keep track of what you do.
Simpler algorithms (in some cases)
Erase risk for deadlocks etc.
Less chance for bugs.


### List some advantages of using lock-based synchronization primitives over message passing.
> It might in some cases be simpler to implement. And also a bit faster to run.
