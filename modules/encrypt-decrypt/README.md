## encrypt-decrypt

<p align="left">
    <a href="https://user-images.githubusercontent.com/36189053/71233242-f63f4c80-232f-11ea-9ce1-4dc76dcadea3.gif"><img src="https://user-images.githubusercontent.com/36189053/71233242-f63f4c80-232f-11ea-9ce1-4dc76dcadea3.gif"></a>
</p>

#### encrypt
```
# all
pjx encrypt -r -d G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config -file *.json -secret helloworld

# short
cd G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config
pjx encrypt -file *.json -secret helloworld

# shortest
cd G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config

# After input below, will require you to input secret, input helloworld
# equal to `pjx encrypt 2.json,3.json`
pjx encrypt *.json
```
This command will locate ``.../config/2.json` and `3.json` and build `3.json.crt`, `2.json.crt`

#### decrypt

```
# all
pjx decrypt -r -d G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config -file *.crt -secret helloworld

# short
cd G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config
pjx decrypt -file *.crt -secret helloworld

# shortest
cd G:\go_workspace\GOPATH\src\pjx\modules\encrypt-decrypt\config

# After input below, will require you to input secret, input helloworld
pjx decrypt -file *.crt
```

This command will locate `.../config/2.json.crt` and `3.json.crt` and build `3.json`, `2.json`
