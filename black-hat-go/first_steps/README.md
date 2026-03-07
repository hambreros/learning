# First steps

write file `main.go`

run it

```sh
go run main.go

Hello, Black Hat Gophers!
```

rename it and build it

```sh
mv main.go hello.go

go build hello.go
./hello
```

compile without debugging info and symbol table

```sh
ls -lrth
total 4880
-rw-r--r--@ 1 michael  staff    91B  7 Mar 15:48 hello.go
-rwxr-xr-x@ 1 michael  staff   2.4M  7 Mar 15:49 hello

go build -ldflags "-w -s" hello.go
ls -lrth
total 3248
-rw-r--r--@ 1 michael  staff    91B  7 Mar 15:48 hello.go
-rwxr-xr-x@ 1 michael  staff   1.6M  7 Mar 15:50 hello
```

cross compile

```sh
go build -ldflags "-w -s" hello.go
file hello
hello: Mach-O 64-bit executable arm64

GOOS="windows" GOARCH="amd64" \
  go build -ldflags "-w -s" -o hello.exe hello.go

file hello.exe
hello.exe: PE32+ executable (console) x86-64, for MS Windows
```
