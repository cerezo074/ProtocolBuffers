build: 
	rm -rf userpb
	mkdir userpb
	protoc -I. proto/*.proto --go_out=:userpb
run: 
	go run main.go

.PHONY: build run