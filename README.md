## pjx
pjx is a tool helps auto-generate server side directories and some go code.Supporting linux, windows, mac.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [1. Start](#1-start)
- [2. Module](#2-module)
    - [2.1 Directory generate](#21-directory-generate)
    - [2.2 Package storage and migration](#22-package-storage-and-migration)
- [3. Commands](#3-commands)
    - [3.1 `new`](#31-new)
    - [3.2 `module`](#32-module)
    - [3.3 `add`](#33-add)
    - [3.4 `use`](#34-use)
    - [3.5 `merge`](#35-merge)
    - [3.6 `clone`](#36-clone)
- [4. Optional args](#4-optional-args)
- [5. FAQ](#5-faq)
    - [5.1. pjx command not found?](#51-pjx-command-not-found)
    - [5.2. How to design module directories as wanted?](#52-how-to-design-module-directories-as-wanted)
    - [5.3 permission deny?](#53-permission-deny)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 1. Start

way 1:

- `go get -u github.com/fwhezfwhez/pjx`

way 2:

- `git clone https://github.com/fwhezfwhez/pjx.git`
- `cd pjx`
- `go install`


Make sure in cmd type `pjx --version`, output normal.

## 2. Module
pjx now provides functional flows below:

#### 2.1 Directory generate
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

**module**

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

**Design your own directories refers to [4.2. How to design module directories as wanted?](#42-how-to-design-module-directories-as-wanted)**

`// - pjx test-client functionName [--http] [--tcp] [--grpc]`

#### 2.2 Package storage and migration
**Make sure configure system env `pjx_path`.This will let pjx know where to storage package locally**

For storing and migrating local package, pjx provides commands below:

Take helloworld for example:

`pjx add helloworld` add a package named helloword in current dir into repo

`pjx use helloworld` insert a package from repo to current dir

<p align="left">
    <a href="http://i1.fuimg.com/684630/839f83b7f1f3669a.gif"><img src="http://i1.fuimg.com/684630/839f83b7f1f3669a.gif"></a>
</p>


**Make sure configure system env `pjx_path`.This will let pjx know where to storage package locally.If not well set, it still work and will save packages in user-home/pjx_path**

`add` and `use` keyword ruled below:

(value in '<>' is necessary, '[]' is optional. if `namespace`, `tag` not set, use 'global' and 'master'. ':value' means it's a value not fixed.)

`pjx add <packageName> [:namespace] [:tag] [-f]` it will add a package in current dir into repo `${pjx_path}/:namespace/:package-name@:tag`.Optional args:

- `-f` force add package, if exist, replace old.

`pjx use <packageName> [:namespace] [:tag] [-o :rename]` it will insert a package from repo into current dir.Optional args:

- `-o` add package with another name.

## 3. Commands

#### 3.1 `new`
`pjx new <appName>` new a project
```go
pjx new hello
```

#### 3.2 `module`
`pjx module <moduleName> [-m <key>]` using specific template to generate module.

- To select module template, you can refer to FAQ 5.2.
```go
pjx module user
pjx module shop -m test
```

#### 3.3 `add`
`pjx add <packageName> [namespace] [tag] [-f]` add package into repo(at `${pjx_path}`). If pjx_path not set, they will be stored in user-home/pjx_path.
```go
pjx add hello // package hello will be add into pjx_path/global/hello
pjx add hello fwhezfwhez master// package hello will be add into pjx_path/fwhezfwhez/hello
pjx add hello global tmp // package hello will be add into pjx_path/global/hello@tmp
```
#### 3.4 `use`
`pjx use <packageName> [namespace] [tag] [-o <rename>]` use a package from repo and insert to current dir.If pjx_path not set, use pjx_path default `user-home/pjx_path`.

```go
pjx use hello // use pjx_path/global/hello and insert to current dir.
pjx use hello fwhezfwhez tmp -o hello2 // use pjx_path/fwhezfwhez/hello@tmp and insert into current dir with name hello2
```

#### 3.5 `merge`
`pjx merge <path> <namespace> [-f/-u]` merge all packages in path into namespace
```go
pjx merge /home/web/repo global -f // copy all /home/web/repo's sub dir into pjx_path/global, if exists, replace the old.
pjx merge /Users/web/repo fwhezfwhez -u // copy all /Users/web/repo's sub dir into pjx_path/fwhezfwhez, if exists, jump this.
```

#### 3.6 `clone`
`pjx clone url.git <namespace> [-u/-f]` clone a remote repo,and copy all sub dir in it to pjx_path/<namespace>
```go
pjx clone https://github.com/fwhezfwhez/pjx-repo.git global -u  jump the existing case
pjx clone https://github.com/fwhezfwhez/pjx-repo.git global -f  replace the existing old
```

## 4. Optional args

| value | meaning | example | why | scope |
| ---- | ---- | ---- | --- | ---- |
| -l | open log | pjx add xxx -l | show log | all |
| -f | add package by force | pjx add xxx -f | avoid package exist error | add, merge, clone |
| -o | use package in another name | pjx use xxx -o xxx2 | avoid package exist error| use |
| -m | choose module template,it's at 'pjx/module-template.go' | pjx module user -m test| to design module directories as wanted | module |
| -u | jump existed package with the same name when meet command `merge` and `clone`| PJX merge g:/repo fwhezfwhez -u | avoid package exist error | merge, clone |

## 5. FAQ

#### 5.1. pjx command not found?

`go get -u ...` or `go install` will put `pjx` execute file into ${GOPATH}/bin. Make sure your ${GOPATH}/bin is in your system path.

#### 5.2. How to design module directories as wanted?

If you don't like `xRouter`, `xService`... this kind of directories, you can just modify `module-template.go`.By default, there is two keys `default` and `test` to refer the template.Pjx will choose as below:

- `pjx module user`   using default key,and generate xRouter,xService,xModel,xPb,xTestClient,xExport
- `pjx module user -m test` using test key, and generate xModel,xService,xRouter.

**After you modify module-template.go, don't forget to run `go install` to refresh pjx command**
<p align="left">
    <a href="http://i2.tiimg.com/684630/aac5c27572431f5f.gif"><img src="http://i2.tiimg.com/684630/aac5c27572431f5f.gif"></a>
</p>

#### 5.3 permission deny?
Make sure the spot you execute `pjx` has proper permission. For instance, when you execute `pjx add xx`, you must has read permission to the package xx and write permission to pjx_path.This kind of question should never depend on pjx to fix it.

