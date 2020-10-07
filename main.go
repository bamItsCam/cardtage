// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"github.com/bamItsCam/cardtage/cmd"
	"gopkg.in/gographics/imagick.v2/imagick"
)


func init() {
	imagick.Initialize()
	// Schedule cleanup
	defer imagick.Terminate()
}

func main() {
	//args := cardtage.CardtageImpl{
	//	InFileGlob:        "input/*",
	//	OutFilename:       "out.pdf",
	//	Density:           100,
	//	Unit:              imagick.RESOLUTION_PIXELS_PER_INCH,
	//	CardWidthU:        2.64,
	//	CardHeightU:       1.73,
	//	PageWidthU:        8.5,
	//	PageHeightU:       11,
	//	PageMarginWidthU:  0.2,
	//	PageMarginHeightU: 0.2,
	//	CardBorderU:       .02,
	//}
	//if err := args.Generate(); err != nil {
	//	panic(err)
	//}
	if err := cmd.Execute(); err != nil {
		panic(err)
	}

}