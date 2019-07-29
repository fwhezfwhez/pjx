package  main

var moduleTmplMap = map[string]string{
	"default": moduleTmpl,
    "test": test,
}

var moduleTmpl = `
# Help generate module directories as you want. 'pjx module user'
## generate where. If same dir as execute path, set it empty.
#### package: empty

package: module
dirList:
- '{$object}Service'
- '{$object}Router'
- '{$object}Model'
- '{$object}Pb'
- '{$object}Export'
- '{$object}TestClient'
`

var test = `
package: empty
dirList:
- '{$object}Service'
- '{$object}Router'
- '{$object}Model'
`



