# snowflake
A tiny module to generate time based 64-bit unique id, inspired by Twitter id (snowflake).

# Installation
```
go get -u github.com/zhangpanyi/snowflake
```

# Usage
### initialization
```go
import (
    "github.com/zhangpanyi/snowflake"
)

worker, err := snowflake.NewWorker(1, 0)
if err != nil {
	panic(err)
}
```
Create a instance of snowflake as shown above which will be used to generate snowflake ids afterward.

### Snowflake ID generation
```go
id1 := worker.Generate()
id2 := worker.Generate()
fmt.Println(id1, id2)
```
Output:
```
6613627584607424512 6613627584607424513
```
