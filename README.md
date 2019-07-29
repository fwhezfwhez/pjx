## pjx
pjx is a tool helps auto-generate server side directories and some go code.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [pjx](#pjx)
- [1. Start](#1-start)
- [2. Module](#2-module)
    - [2.1 directory generate](#21-directory-generate)
    - [module](#module)
    - [2.2 package storage and migration](#22-package-storage-and-migration)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 1. Start
`go get -u github.com/fwhezfwhez/pjx`

Make sure in cmd type `pjx --version`, output normal.

## 2. Module
pjx now provides functional flows below:

#### 2.1 directory generate
For developing golang server project, pjx can auto-generate projects.

Take helloWorld project for example:

`pjx new helloworld`

`cd helloworld`

`pjx module user`

<p align="left">
    <a href="http://i2.tiimg.com/684630/b10b449ce4e75370.gif"><img src="http://i2.tiimg.com/684630/b10b449ce4e75370.gif"></a>
</p>


The generated directories will look like:

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

`// - pjx test-client functionName [--http] [--tcp] [--grpc]`

#### 2.2 package storage and migration
**Make sure configure system env `pjx_path`.This will let pjx know where to storage package locally**

For storing and migrating local package, pjx provides commands below:

Take helloworld for example:

`pjx add helloworld` add a package named helloword in current dir into repo

`pjx use helloworld` insert a package from repo to current dir

<p align="left">
    <a href="http://i1.fuimg.com/684630/839f83b7f1f3669a.gif"><img src="http://i1.fuimg.com/684630/839f83b7f1f3669a.gif"></a>
</p>


**Make sure configure system env `pjx_path`.This will let pjx know where to storage package locally**

`add` and `use` keyword ruled below:

(value in '<>' is necessary, '[]' is optional. if `namespace`, `tag` not set, use 'global' and 'master'. ':value' means it's a value not fixed.)

`pjx add <packageName> [:namespace] [:tag] [-f]` it will add a package in current dir into repo `${pjx_path}/:namespace/:package-name@:tag`.Optional args:

- `-f` force add package, if exist, replace old.

`pjx use <packageName> [:namespace] [:tag] [-o :rename]` it will insert a package from repo into current dir.Optional args:

- `-o` add package with another name.
