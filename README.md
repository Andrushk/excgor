# Something like Semaphores for Go
--

Limits number of concurrent goroutines.
All goroutines above limit will be skiped. This will occur until one of working goroutines will finish work. After that the empty place may take any goroutine (but only one).

## Installation

    $ go get github.com/Andrushk/excgor
