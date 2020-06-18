//+build mage

package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/magefile/mage/sh"
)

var Default = Local

// NewPost creates a new post with the passed name
// run with 'HUGOSUMMERAME=name mage newpost'
func NewPost() error {
	files, err := ioutil.ReadDir("./content/posts")
	if err != nil {
		return err
	}
	return sh.RunV("hugo", "new", "posts/Day"+strconv.Itoa(len(files)+1)+".md")
}

// Build builds the site.
func Build() error {
	if err := Clean(); err != nil {
		return err
	}
	return sh.RunV("hugo")
}

// Clean cleans up the public folder.
func Clean() error {
	return os.RemoveAll("public/")
}

// Local runs the server at 8080
func Local() error {
	return sh.RunV("hugo", "serve", "-p", "8080")
}