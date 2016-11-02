package test

import (
	"fmt"
	"time"
	"testing"
)

func TestTime(t *testing.T) {
	time.UTC
	fmt.Printf("\n +%v",time.Now().UTC())
}

