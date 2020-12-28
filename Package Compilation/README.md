# Protoc
In order to compile your files with .proto extension you need to install protoc cli. On OSX installation is easy and you can do it through Homebrew with the follow command. 

*brew install protobuf*

After installation, run this command.

*protoc .proto -I. --go_out=:addresspb* 

Remember you need to be in the same folder(**-I** option) in which your proto files are located and the **--go_out** option seeks for the folder you specify as the place to put your files compiled. Just in case you'll face some problems checkout your golang env variables like GOPATH and GOBIN.

 - export GOBIN=$GOPATH/bin    
 - export GOPATH=$HOME/go

Just check that $GOBIN is added to your $GOPATH, and $GOPATH to your $PATH variable.
