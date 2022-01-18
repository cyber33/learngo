## Where should you save your Go source code?
* Anywhere on my computer
* Under $GOPATH
* Under $GOPATH/src *CORRECT*

## What does $GOPATH mean?
* It's a file for Go runtime
* Stores Go source code files and compiled packages
* It's a path for gophers to follow

## Do you need to set $GOPATH?
* Yes
* No: It's stored on my desktop
* No: It's stored under my user path *CORRECT*

## How can you print $GOPATH?
* Using `ls` command
* Using `go env GOPATH` command *CORRECT*
* Using `go environment` command




C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>go env GOPATH
C:\Program Files\go

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>dir "C:\Program Files\go"
 Volume in drive C is Acer
 Volume Serial Number is AE3D-6D4D

 Directory of C:\Program Files\go

11/20/2021  11:36 AM    <DIR>          .
11/20/2021  11:36 AM    <DIR>          ..
11/20/2021  11:36 AM    <DIR>          api
11/04/2021  02:05 PM            55,782 AUTHORS
11/23/2021  08:03 PM    <DIR>          bin
11/04/2021  02:05 PM                52 codereview.cfg
11/04/2021  02:05 PM             1,339 CONTRIBUTING.md
11/04/2021  02:05 PM           107,225 CONTRIBUTORS
11/20/2021  11:36 AM    <DIR>          doc
11/20/2021  11:35 AM    <DIR>          lib
11/04/2021  02:05 PM             1,479 LICENSE
11/20/2021  11:36 AM    <DIR>          misc
11/04/2021  02:05 PM             1,303 PATENTS
11/20/2021  11:35 AM    <DIR>          pkg
11/04/2021  02:05 PM             1,480 README.md
11/04/2021  02:05 PM               397 SECURITY.md
11/20/2021  11:36 AM    <DIR>          src
11/20/2021  11:36 AM    <DIR>          test
11/04/2021  02:05 PM                 8 VERSION
               9 File(s)        169,065 bytes
              10 Dir(s)   8,575,434,752 bytes free

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>dir "C:\Program Files\go\src
 Volume in drive C is Acer
 Volume Serial Number is AE3D-6D4D

 Directory of C:\Program Files\go\src

11/20/2021  11:36 AM    <DIR>          .
11/20/2021  11:36 AM    <DIR>          ..
11/04/2021  02:05 PM               407 all.bash
11/04/2021  02:05 PM               754 all.bat
11/04/2021  02:05 PM               388 all.rc
11/20/2021  11:35 AM    <DIR>          archive
11/04/2021  02:05 PM             3,790 bootstrap.bash
11/20/2021  11:36 AM    <DIR>          bufio
11/04/2021  02:05 PM             1,872 buildall.bash
11/20/2021  11:35 AM    <DIR>          builtin
11/20/2021  11:36 AM    <DIR>          bytes
11/04/2021  02:05 PM               518 clean.bash
11/04/2021  02:05 PM               598 clean.bat
11/04/2021  02:05 PM               380 clean.rc
11/20/2021  11:36 AM    <DIR>          cmd
11/04/2021  02:05 PM             1,519 cmp.bash
11/20/2021  11:35 AM    <DIR>          compress
11/20/2021  11:35 AM    <DIR>          container
11/20/2021  11:36 AM    <DIR>          context
11/20/2021  11:36 AM    <DIR>          crypto
11/20/2021  11:35 AM    <DIR>          database
11/20/2021  11:35 AM    <DIR>          debug
11/20/2021  11:35 AM    <DIR>          embed
11/20/2021  11:36 AM    <DIR>          encoding
11/20/2021  11:35 AM    <DIR>          errors
11/20/2021  11:35 AM    <DIR>          expvar
11/20/2021  11:36 AM    <DIR>          flag
11/20/2021  11:36 AM    <DIR>          fmt
11/20/2021  11:36 AM    <DIR>          go
11/04/2021  02:05 PM               288 go.mod
11/04/2021  02:05 PM             1,538 go.sum
11/20/2021  11:36 AM    <DIR>          hash
11/20/2021  11:36 AM    <DIR>          html
11/20/2021  11:36 AM    <DIR>          image
11/20/2021  11:35 AM    <DIR>          index
11/20/2021  11:36 AM    <DIR>          internal
11/20/2021  11:36 AM    <DIR>          io
10/10/2021  01:58 PM    <DIR>          k8s.io
11/20/2021  11:36 AM    <DIR>          log
11/04/2021  02:05 PM             7,135 make.bash
11/04/2021  02:05 PM             4,666 make.bat
11/04/2021  02:05 PM               553 Make.dist
11/04/2021  02:05 PM             3,110 make.rc
11/20/2021  11:36 AM    <DIR>          math
11/20/2021  11:36 AM    <DIR>          mime
11/20/2021  11:36 AM    <DIR>          net
11/20/2021  11:36 AM    <DIR>          os
11/20/2021  11:36 AM    <DIR>          path
11/20/2021  11:36 AM    <DIR>          plugin
11/04/2021  02:05 PM             1,106 race.bash
11/04/2021  02:05 PM             1,091 race.bat
11/04/2021  02:05 PM             2,295 README.vendor
11/20/2021  11:36 AM    <DIR>          reflect
11/20/2021  11:36 AM    <DIR>          regexp
11/04/2021  02:05 PM             1,851 run.bash
11/04/2021  02:05 PM             1,034 run.bat
11/04/2021  02:05 PM               436 run.rc
11/20/2021  11:36 AM    <DIR>          runtime
11/20/2021  11:36 AM    <DIR>          sort
11/20/2021  11:36 AM    <DIR>          strconv
11/20/2021  11:36 AM    <DIR>          strings
11/20/2021  11:36 AM    <DIR>          sync
11/20/2021  11:36 AM    <DIR>          syscall
11/20/2021  11:35 AM    <DIR>          testdata
11/20/2021  11:36 AM    <DIR>          testing
11/20/2021  11:35 AM    <DIR>          text
11/20/2021  11:36 AM    <DIR>          time
11/20/2021  11:36 AM    <DIR>          unicode
11/20/2021  11:35 AM    <DIR>          unsafe
11/20/2021  11:35 AM    <DIR>          vendor
              21 File(s)         35,329 bytes
              49 Dir(s)   8,575,434,752 bytes free

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo> dir "C:\Program Files\go\bin"
 Volume in drive C is Acer
 Volume Serial Number is AE3D-6D4D

 Directory of C:\Program Files\go\bin

11/23/2021  08:03 PM    <DIR>          .
11/23/2021  08:03 PM    <DIR>          ..
11/20/2021  10:24 AM        15,753,216 dlv-dap.exe
11/23/2021  08:03 PM         4,041,216 g3nicodes.exe
11/23/2021  08:03 PM         4,045,824 g3nshaders.exe
11/23/2021  08:03 PM         4,344,320 glapi2go.exe
11/20/2021  10:20 AM         3,393,536 go-outline.exe
11/04/2021  02:08 PM        14,077,952 go.exe
11/04/2021  02:08 PM         3,535,872 gofmt.exe
11/20/2021  10:20 AM        24,153,600 gopls.exe
               8 File(s)     73,345,536 bytes
               2 Dir(s)   8,580,771,840 bytes free

C:\Users\mark_\Data\Google\Interview Prep\Golang\learngo>