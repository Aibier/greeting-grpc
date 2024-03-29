# gRPC Go

### Build

#### `Linux/MacOS`

```shell
make all
```
***all is a Makefile rule** - check the other [rules](#makefile)

#### `Windows - Chocolatey`
```shell
choco install make
make all
```
***all is a Makefile rule** - check the other [rules](#makefile)

#### `Windows - Without Chocolatey`

```shell
protoc -Igreet/proto --go_opt=module=github.com/Aibier/greeting-grpc --go_out=. --go-grpc_opt=module=github.com/Aibier/greeting-grpc --go-grpc_out=. greet/proto/*.proto

protoc -Icalculator/proto --go_opt=module=github.com/Aibier/greeting-grpc --go_out=. --go-grpc_opt=module=github.com/Aibier/greeting-grpc --go-grpc_out=. calculator/proto/*.proto

protoc -Iblog/proto --go_opt=module=github.com/Aibier/greeting-grpc --go_out=. --go-grpc_opt=module=github.com/Aibier/greeting-grpc --go-grpc_out=. blog/proto/*.proto

go build -o bin/greet/client.exe ./greet/client
go build -o bin/greet/server.exe ./greet/server

go build -o bin/calculator/client.exe ./calculator/client
go build -o bin/calculator/server.exe ./calculator/server

go build -o bin/blog/client.exe ./blog/client
go build -o bin/blog/server.exe ./blog/server
```

<a name="makefile"></a>
## Makefile

For more information about what are the rules defined in the Makefile, please type:

```shell
make help
```

## Reporting a bug

As I need to know a little bit more information about your environment to help you, when filling an issue, please provide the output of:

```shell
make about
```
## Go env setup in the curren directory
export PATH=$PATH:$HOME/go/bin, in order to add the GOPATH and
export PATH=$PATH:/usr/local/go/bin, in order to add GOROOT