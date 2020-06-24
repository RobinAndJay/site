//+build mage

package main

import (
	"io/ioutil"
	"strconv"
	"time"

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
	if err := sh.RunV("hugo"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "--all"); err != nil {
		return err
	}
	if err := sh.RunV("git", "commit", "-m", "rebuilding site: "+time.Now().String()); err != nil {
		return err
	}
	if err := sh.RunV("git", "push"); err != nil {
		return err
	}
	if err := sh.RunV("cd", "public"); err != nil {
		return err
	}
	if err := sh.RunV("git", "add", "--all"); err != nil {
		return err
	}
	if err := sh.RunV("git", "commit", "-m", "rebuilding site: "+time.Now().String()); err != nil {
		return err
	}
	if err := sh.RunV("git", "push"); err != nil {
		return err
	}
	return sh.RunV("cd", "..")
}

// Local runs the server at 8080
func Local() error {
	return sh.RunV("hugo", "serve", "-p", "8080")
}
