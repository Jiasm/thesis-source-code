package main

import (
	"constant"
	"entity"
	"testing"
)

func Test_main(t *testing.T) {
	testGroup := entity.Group{1, "Niko", 0, 2, 1 }

	t.Log("build group struct", testGroup)

	t.Log("get session key", constant.KEY_USER)
}