Test commands:

go test .
go test -v .
go test -coverprofile=cover.out
go tool cover -html=cover.out
go test -cover .

#To run single test 
go test -run Test_isPrime
go test -v -run Test_isPrime

#To run group of tests
    funct Test_<groupname>_isPrime()
    go -v -run Test_<groupname>