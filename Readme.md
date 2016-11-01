# Learning GO..

but also reading paper https://www.microsoft.com/en-us/research/wp-content/uploads/2016/02/Orleans-MSR-TR-2014-41.pdf

1. Following https://golang.org/doc/code.html

## Building :
1. Install Go
2. Run following commands :

```
mkdir $HOME/golang
export GOPATH=$HOME/golang
export PATH=$PATH:$GOPATH/bin

cd $HOME/golang
go install github.com/ashishnegi/goorleans
go test github.com/ashishnegi/goorleans
$GOPATH/goorleans
```
