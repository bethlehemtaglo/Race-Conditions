Results from Task 3:
The output differs from each time meaning we have a case of race conditions.
Both functions increment() and decrement() are using the same global variable, meaning its a 'race' to whichever finishes first. Making the output different each time

Results from Task 4:
Inserting a delay by locking whos turn it is to adress the global variable, actually ensures the correct result of 0. Because in theory increasing a number, than decresing the same number will always equal 0. 

Resulst from Task 5:
With Go i made a channel for the buffer to ensure the synchronization between goroutines. The output is as expected with the producer producing the numbers and adding them to the channel, and with blocking the consumer can be synchronized with the proudcer, therefore only consuming when there is a value. It also outputs a deadlock when all the goroutines are a sleep, knowing we have reached the end of the cycle. 

