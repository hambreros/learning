# Learning

## Tue 2nd March

- isntall msys2 as it seems to work on windows without need for admin
- still can't seem to find `gtypsit`
- can install the follwoing
```
pacman -S mingw-w64-ucrt-x86_64-gcc
pacman -S vim
pacman -S git
pacman -S python
pacman -S mingw-w64-ucrt-x86_64-go
```

```
# can also
set -o vi
```

- python seems to work
```
python --version
Python 3.12.12
```

- go still needs some setup
```
go version
go: cannot find GOROOT directory: 'go' binary is trimmed and GOROOT is not set
```

Needed to copy some files from a Mac - tried to connect via smb://192.168.68.1
but didn't seem to work. Instead fired up a local python server on the mac

```
python -m http.server --directory ./ 4000

# then on the windows machine MSYS2
wget http://192.168.68.1:4000
```

read through the index and fetched the files I needed
