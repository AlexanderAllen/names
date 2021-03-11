package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	name := namesgenerator.GetRandomName(0)
	fmt.Println(strings.Replace(name, "_", "-", 1))
}
