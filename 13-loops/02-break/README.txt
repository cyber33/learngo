C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>type main.go
// Copyright ┬⌐ 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
        "fmt"
)

func main() {
        var (
                sum int
                i   = 1
        )

        for {
                if i > 5 {
                        break
                }

                sum += i

                fmt.Println(i, "-->", sum)

                i++
        }

        fmt.Println(sum)
}

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>go run main.go
1 --> 1
2 --> 3
3 --> 6
4 --> 10
5 --> 15
15

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>copu main.go TryIt.go
'copu' is not recognized as an internal or external command,
operable program or batch file.

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>copy main.go TryIt.go
        1 file(s) copied.

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>go run TryIt.go
1 --> 1
2 --> 3
3 --> 6
4 --> 10
5 --> 15
15

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>notepad TryIt.go

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>go run TryIt.go
# command-line-arguments
.\TryIt.go:26:7: invalid operation: sum += i (mismatched types int and float64)

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>go run TryIt.go
1 --> 1
2 --> 3
3 --> 6
4 --> 10
5 --> 15
15

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>go run TryIt.go
1.2 --> 1.2
2.2 --> 3.4000000000000004
3.2 --> 6.6000000000000005
4.2 --> 10.8
10.8

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>type TryIt.go
// Copyright ┬⌐ 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
        "fmt"
)

func main() {
        var (
                sum float64
                i   = 1.2
        )

        for {
                if i > 5 {
                        break
                }

                sum += i

                fmt.Println(i, "-->", sum)

                i++
        }

        fmt.Println(sum)
}

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>pushd ..\..\

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>git add
Nothing specified, nothing added.
hint: Maybe you wanted to say 'git add .'?
hint: Turn this message off by running
hint: "git config advice.addEmptyPathspec false"

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>git add .

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>git commit -m "Try adding float in the loop"
[master d1f15f5] Try adding float in the loop
 1 file changed, 34 insertions(+)
 create mode 100644 13-loops/02-break/TryIt.go

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>popd

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo\13-loops\02-break>