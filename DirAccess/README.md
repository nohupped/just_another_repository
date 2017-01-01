
strace output when using readdir in C when compared to Go: Note that this doesn't do an (f/l)stat on all files in that directory unlike the go equivalent of os.File.Readdirnames.

#### C program strace count output

```
➜  ~ sudo strace -c ./a.out 
% time     seconds  usecs/call     calls    errors syscall
------ ----------- ----------- --------- --------- ----------------
  0.00    0.000000           0         1           read
  0.00    0.000000           0       286           write
  0.00    0.000000           0         3           open
  0.00    0.000000           0         3           close
  0.00    0.000000           0         4           fstat
  0.00    0.000000           0         8           mmap
  0.00    0.000000           0         3           mprotect
  0.00    0.000000           0         1           munmap
  0.00    0.000000           0         3           brk
  0.00    0.000000           0         3         3 access
  0.00    0.000000           0         1           execve
  0.00    0.000000           0         2           getdents
  0.00    0.000000           0         1           arch_prctl
------ ----------- ----------- --------- --------- ----------------
100.00    0.000000                   319         3 total

```


#### Go code

```
package main
import (
		"fmt"
	"os"
)
func main() {
	x, err := os.Open("/")
	if err != nil {
		panic(err)
	}
	y, err := x.Readdir(0)
	if err != nil {
		panic(err)
	}
	for _, i := range y {
	fmt.Println(i)
	}

}
```

#### Go program strace count output

```
➜  ~ sudo strace -c ./goprog

% time     seconds  usecs/call     calls    errors syscall
------ ----------- ----------- --------- --------- ----------------
100.00    0.000021           0       114           rt_sigaction
  0.00    0.000000           0        27           write
  0.00    0.000000           0        27           lstat
  0.00    0.000000           0         8           mmap
  0.00    0.000000           0         1           munmap
  0.00    0.000000           0         8           rt_sigprocmask
  0.00    0.000000           0         3           clone
  0.00    0.000000           0         1           execve
  0.00    0.000000           0         2           sigaltstack
  0.00    0.000000           0         1           arch_prctl
  0.00    0.000000           0         1           gettid
  0.00    0.000000           0         1           futex
  0.00    0.000000           0         1           sched_getaffinity
  0.00    0.000000           0         2           getdents64
  0.00    0.000000           0         1           openat
------ ----------- ----------- --------- --------- ----------------
100.00    0.000021                   198           total


```
