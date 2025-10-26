# Solution Steps

1. Create a new Go module for your project (optional: use 'go mod init').

2. Import the necessary packages: fmt, sync, and time.

3. Define a producer function that sends integers 1-8 into a channel and closes the channel after sending all tasks.

4. Define a consumer function that accepts a consumer ID, the tasks channel, and a pointer to a WaitGroup. Inside the function, use a defer statement to call wg.Done(). Loop over 'tasks' using 'for task := range tasks', printing the processed task and consumer ID, then sleep for 1 second to simulate work.

5. In the main function, create an unbuffered channel of type int for tasks.

6. Initialize a sync.WaitGroup and set its counter to 2 (for two consumers).

7. Start one producer goroutine to send tasks into the channel.

8. Start two consumer goroutines, passing a unique ID (1 and 2), the tasks channel, and the WaitGroup to each.

9. Call wg.Wait() in main to block until both consumers finish processing and exit.

10. Run the program to verify that all 8 tasks are processed and both consumers exit cleanly without deadlock.

