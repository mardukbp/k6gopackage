# Adding a Go package to k6

1. `go get github.com/loadimpact/k6`
2. `cd $GOPATH/src/github.com/loadimpact/k6`
3. `go get -u github.com/mardukbp/k6gopackage@version`
4. In `js/modules/index.go`:
   - add `"github.com/mardukbp/k6gopackage"` to the list of imported packages
   - add an entry to the `Index` map that begins with **"k6/"** (this is hardcoded in k6)
     (e.g. `"k6/k6gopackage": k6gopackage.New(),`)
5. Add the module to the `vendor` directory:
    `go mod vendor`
6. Build k6: `go build`
7. Try it: `./k6 run vendor/github.com/mardukbp/k6gopackage/k6gopackage.js`
