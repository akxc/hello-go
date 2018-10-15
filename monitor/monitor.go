package monitor

import (
	"fmt"
)

func Connect() {
	job := Job{name: "Job in monitor", age: 10}
	fmt.Println("connect to a ssh", job)
}
