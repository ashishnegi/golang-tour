# On a tour of GO in 5 hours..

1. Following https://golang.org/doc/code.html
2. Touring https://tour.golang.org/welcome/1

## Building :
1. Install Go
2. Run following commands :

```
mkdir $HOME/golang
export GOPATH=$HOME/golang
export PATH=$PATH:$GOPATH/bin

cd $HOME/golang
go install github.com/ashishnegi/goblocks
go test github.com/ashishnegi/stringutils
$GOPATH/bin/goblocks
```

3. Run particular program
```
go run github.com/ashishnegi/search-engine/search.go
```
4. Get test dependencies
```
go get -t
```

Presentation: https://docs.google.com/presentation/d/1IWn_Yxi7B7KwH5oyYQElc3INbJY5Rw2R2ZWdVx7_0uk/edit?usp=sharing
Tutorial: https://github.com/ashishnegi/golang-tour/blob/master/tutorial.txt
