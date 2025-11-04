package main

import (
	"fmt"
	"testing"

	"github.com/taneekpet/pascal-compiler-with-go/models"
)

func Test_parseProgram(t *testing.T) {
	token, err := parseProgram(models.NewSlide("PROGRAM p_name ;  ."))
	fmt.Printf("%v %v\n", token, err)
}
