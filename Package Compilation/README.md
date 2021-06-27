# Protoc
In order to compile your files with .proto extension you need to install protoc cli. On OSX installation is easy and you can do it through Homebrew with the follow command. 

*brew install protobuf*

After installation, run this command

*protoc -I. proto/*.proto --go_out=:addresspb* 

**-I** option indicates where are gonna be our proto files and the **--go_out** option seeks for the folder you specify as the place to put your files compiled. Just in case if you'll face some problems checkout your golang env variables like GOPATH and GOBIN. Don't forget it this command is customized for this example, that's why we use addresspb as the output directory cuz we have created that folder and inside proto folder we have our files to be compiled(* indicates we are gonna compile all files).

 - export GOBIN=$GOPATH/bin    
 - export GOPATH=$HOME/go

Just check that $GOBIN is added to your $GOPATH, and $GOPATH to your $PATH variable.

**NOTE:** If you got stuck with this error ***The import path must contain at least one forward slash ('/') character.***. In some of your proto files, look over your **option go_package** definition. Somehow the protoc compiler has changed the way it works with relative packages. For more information open this [link](https://github.com/techschool/pcbook-go/issues/3#issuecomment-823206034)
