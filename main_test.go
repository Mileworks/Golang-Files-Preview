package main

import (
	"github.com/mholt/archiver/v3"
	"testing"
)

func TestUnzip(t *testing.T) {
	archiver.Unarchive("/Users/mac/Downloads/AutoVue配置操作手册.zip", "tmp/decompress")

}