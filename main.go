// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"github.com/bamItsCam/cardtage/cmd"
	"gopkg.in/gographics/imagick.v2/imagick"
	"os"
)

func init() {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
}

func main() {
	// Todo: add support for reading a config value that gives quantities per image
	// Todo: add mm support
	if err := cmd.Execute(); err != nil {
		//fmt.Println(err.Error())
		os.Exit(1)
	}

}
