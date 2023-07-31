hello: hello/main.go # only build hello again if main.go has changed
	echo "building hello"
	go build -o hello/hello hello/main.go

.PHONY: run-hello
run-hello: # no dependency to trigger the target. a phony target always runs
	go run hello/main.go
