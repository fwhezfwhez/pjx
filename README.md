## pjx
pjx is a tool helps auto-generate server side directories and some go code.The generated directories will look like:

## Start
`go get -u github.com/fwhezfwhez/pjx`

Make sure in cmd type `pjx --version`, output normal.

## Flow
Take helloWorld project for example:

`pjx new helloWorld`

`cd helloWorld`

`pjx module user`

## Doc

```txt
appName
  | -- module
  | -- config
  | -- dependence
  | -- independence
  | -- main.go
```
What are they?

- appName: project name, for example `helloWorld`
- module: all modules about service, for example `user`, `shop`. Each module has inner directories.They're documented below.
- config: some config of the project.
- dependence: packages or files of common util tool. These packages and files might import project's inner package.
- independence: packages or files of common util tool. These pkg and files will not import any of this project.It can be no-harm add, remove, reuse-copy.
- main.go: project entrance.

#### module
Module divides project into modules such as `user`, `shop`.Its generated directories will look like:
```txt
module
  | -- user
  |     | -- userPb
  |     | -- userModel
  |     | -- userRouter
  |     | -- userService
  |     | -- userTestClient
  |     | -- userExport
  | ...
```

What are they?

- user: module name.
- userPb: proto file and generated go file.
- userModel: db model or service model.
- userRouter: http, tcp router.
- userService: http, tcp service codes.
- userTestClient: generate test as client codes.
- userExport: export user module as a single server.

Commands are below:

- pjx new appName
- pjx module moduleName
- pjx test-client functionName [--http] [--tcp] [--grpc]
