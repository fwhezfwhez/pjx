## versions

#### v1.1.0
1. provide commands:
   - `pjx version` show its version
   - `pjx help` show help list
   - `pjx new helloword` generate a directory for app helloword
   - `cd helloword` -> `pjx module user`  generate module user for helloword

#### v1.2.0
1. add command:
   - `pjx add <pkg-name> [namespace] [tag]`, `pjx add db` -> `pjx add db global master` add a sub-dir from current dir into pjx local repo.Local repo can be set env value `pjx_path`.
   - `pjx use <pkg-name> [namespace] [tag]`, `pjx use db` -> `pjx use db global master` insert a pkg to current dir from pjx local repo.
