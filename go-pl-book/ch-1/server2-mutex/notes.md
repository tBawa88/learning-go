## Mutex locks
- Mutual exclusion locks are used in concurrent or multithreaded programming to control access of a shared resource. It ensures that at one time only one goroutine or one thread can access the resource, preventing "race conditions".
- **Locks** : When a goroutine wants to access the shared resource, it acquires lock and untill that resource is released, no other go routine can access that resource
- **Unlocks**: Once the goroutine is done using the resource, it releases the mutex, allowing other routines to acquire the lock on that resource
- **Mutual Exclusion** : At any given time, only a single goroutine can acquire the mutex lock, which guarantees that critical sections of code are executed **atomically** (think of a transaction acquiring the row lock in postgresql, making sure that parallel transaction are not executed on same set of rows. )
