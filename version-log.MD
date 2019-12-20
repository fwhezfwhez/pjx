## versions

#### v1.0.0
1. provide commands:
   - `pjx version` show its version
   - `pjx help` show help list
   - `pjx new helloword` generate a directory for app helloword
   - `cd helloword` -> `pjx module user`  generate module user for helloword

#### v2.0.0
1. add commands:
   - `pjx add <pkg-name> [namespace] [tag]` 
   
        `pjx add db` -> `pjx add db global master` add a sub-dir from current dir into pjx local repo.Local repo can be set env value `pjx_path`.
   - `pjx use <pkg-name> [namespace] [tag]` 
   
        `pjx use db` -> `pjx use db global master` insert a pkg to current dir from pjx local repo.
#### v2.2.0
1. add commands:
   - `pjx merge <dir-path> <namespace> [-f] [-u]`
   
        `pjx merge g://repo2 fwhezfwhez -f` add packages from g://repo2 into `${pjx_path}/fwhezfwhez`, if existing the same folder,replace the olds.
   - `pjx clone <git-repo> <namespace> [-f] [-u]`
   
        `pjx clone https://github.com/fwhezfwhez/pjx-repo.git global -f`  add packages from https://github.com/fwhezfwhez/pjx-repo.git into `${pjx_path}/global`, if existing the same folder,replace the olds.

#### v2.3.0
1. add commands:
    - `pjx encrypt [-file] *.json [-d <dir-path>] [-secret <secret>]` encrypt files endwith xx.json to xx.json.crt
    - `pjx decrypt [-file] *.crt [-d <dir-path>] [-secret <secret>]`  decrypt xxx.crt to xxx
