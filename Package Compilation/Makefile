build: 
	rm -rf addresspb
	mkdir addresspb
	protoc -I. proto/*.proto --go_out=:addresspb

.PHONY: build