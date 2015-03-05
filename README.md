# levellog
leveled log wrapper for golang log package

currently only exports static functions 
```
Fatal, Fatalf, Fatalln, Output, Panic, Panicf, Panicln, Print, Printf, Println
```
usage:
```
levellog.Printf(levellog.DEBUG, "Are you %s today?, "happy")
```
