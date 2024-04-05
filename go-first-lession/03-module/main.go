package main

// start: go mod init github.com/bigwhite/module-mode
// Semantic Import Versioning
// Minimal Version Selection
// v1: import "github.com/sirupsen/logrus"
// v2: import "github.com/sirupsen/logrus/v2"
// v0.y.z v1.y.z donot need v0 v1 to import
// 所以 Go 会在该项目依赖项的所有版本中，选出符合项目整体要求的“最小版本”。

// module op 1:
// add "github.com/google/uuid"
// add logrus.Println(uuid.NewString())
// execute  go build
// according advice of IDE : go get github.com/google/uuid
// better: go mod tidy

// module op 2: downgrade version or upgrade version
// execute: go list -m -versions github.com/sirupsen/logrus
// execute: go get package with version
//    go get github.com/sirupsen/logrus@v1.7.0
// or using go mod tidy to dowmgrade version or upgrade version
// example:
// execute: go mod edit --require=github.com/sirupsen/logrus@v1.6.0
// execute: go mod tidy

// module op 3: import package with major version bigger than 1
// execute: go get github.com/go-redis/redis/v7

// module op 4: upgrade to incompatible major version
// execute: go get github.com/go-redis/redis/v8

// module op 5: remove an dependency
// delete import from source
// check all dependencies: go list -m all
// using "go mod tidy" to delete import from go.mod
// execute: go mod tidy

// module op 6: special scenarios using vendor
// After go1.14, if in Go project top level, there is a vendor directory,
// then go build will using vendor firstly, unless you use: go build -mod=mod

import (
	//_ "github.com/go-redis/redis/v8" // “_” void import
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello, go module mode")
	logrus.Println(uuid.NewString())
}
