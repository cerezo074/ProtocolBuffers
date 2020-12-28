build: 
	rm -rf personpb
	mkdir personpb
	protoc -I. proto/*.proto --go_out=:personpb
run: 
	go run main.go

.PHONY: build run