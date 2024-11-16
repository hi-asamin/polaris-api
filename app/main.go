package main

import (
	"polaris-api/infrastructure"
	"polaris-api/infrastructure/router"
)

func main() {
    infrastructure.NewDatabase()
    router.Init()
}
