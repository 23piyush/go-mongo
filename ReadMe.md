go mod init github.com/23piyush/go-mongo  : similar to npm init
go get github.com/julienschmidt/httprouter  : similar to npm install
httprouter - standard library to handle http routes
go.sum is like package-lock.json
go.mod has dependency list
go.sum has more details like dependency of dependencies in go.mod file
go get gopkg.in/mgo.v2 
mgo.v2 - package that helps to interact with mongodb
go get gopkg.in/mgo.v2/bson - may be you will get no output on command window as it is already installed. But in some packages, even though the mother package is installed you will not get the child pacakges. 


>> go run main.go
Error: panic: no reachable servers

goroutine 1 [running]:
main.getSession(...)
        C:/golang-databases/go-mongo/main.go:26
main.main()
        C:/golang-databases/go-mongo/main.go:14 +0x1c5       
exit status 2