package cardtage

import (
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	"math"
	"path/filepath"
)

type Cardtage interface {
	Generate() (err error)
}

type CardtageImpl struct {
	InFileGlob        string
	OutFilename       string
	Density           float64
	Unit              imagick.ResolutionType
	PageWidthU        float64
	PageHeightU       float64
	PageMarginWidthU  float64
	PageMarginHeightU float64
	CardWidthU float64
	CardHeightU float64
	CardBorderU       float64
}

func (c *CardtageImpl) getTileGeo() (string, error) {
	// TODO: do some crazy optimization that reorients cards as needed for optimal count per page.
	printableAreaWidthU := (c.PageWidthU - c.PageMarginWidthU*2)
	printableAreaHeightU := (c.PageHeightU - c.PageMarginHeightU*2)
	maxWidthCount := uint(math.Floor(printableAreaWidthU / (c.CardWidthU + 2*c.CardBorderU)))
	maxHeightCount := uint(math.Floor(printableAreaHeightU / (c.CardHeightU + 2*c.CardBorderU)))
	//fmt.Printf("card dimensions: %.2fx%.2f+%.2f+%.2f\n", c.CardWidthU, c.CardHeightU, c.CardBorderU, c.CardBorderU)
	//fmt.Printf("printable area dimensions: %.2fx%.2f\n", printableAreaWidthU, printableAreaHeightU)

	if maxWidthCount == 0 || maxHeightCount == 0 {
		return "", fmt.Errorf("requested card dimensions are too large for the printable area of the provided page. card is %f.2x%f.2, but printable area only is %f.2x%f.2", c.CardWidthU, c.CardHeightU, printableAreaWidthU, printableAreaHeightU)
	}
	return fmt.Sprintf("%dx%d+0+0", maxWidthCount, maxHeightCount), nil
}

func (c *CardtageImpl) getThumbGeo() string {
	borderPx := uint(c.CardBorderU*c.Density)
	return fmt.Sprintf("+%d+%d", borderPx, borderPx)
}

func (c *CardtageImpl) getExtentOffsetX() (int) {
	return int(-c.PageMarginWidthU *c.Density)
}

func (c *CardtageImpl) getExtentOffsetY() (int) {
	return int(-c.PageMarginWidthU *c.Density)
}

func (c *CardtageImpl) Generate() (err error) {
	mw := imagick.NewMagickWand()
	//fmt.Printf("%+v", c)
	inFilenames, err := filepath.Glob(c.InFileGlob)
	if err != nil {
		return err
	}

	for _, inFilename := range inFilenames {
		// set density of the reader
		if err = mw.SetResolution(c.Density, c.Density); err != nil {
			return err
		}
		//fmt.Println(inFilename)
		if err = mw.ReadImage(inFilename); err != nil {
			return err
		}
		if err = mw.ResizeImage(round2Uint(c.Density*c.CardWidthU), round2Uint(c.Density*c.CardHeightU), imagick.FILTER_LANCZOS, 1); err != nil {
			return err
		}
	}

	collated := imagick.NewMagickWand()

	if err = collated.SetResolution(c.Density, c.Density); err != nil {
		return err
	}

	// this doesn't appear to do anything, but it's nice to dream
	if err = collated.SetGravity(imagick.GRAVITY_CENTER); err != nil {
		return err
	}
	dw := imagick.NewDrawingWand()
	tileGeo, err := c.getTileGeo()
	if err != nil {
		return err
	}

	collated = mw.MontageImage(dw, tileGeo, c.getThumbGeo(), imagick.MONTAGE_MODE_CONCATENATE, "0x0")
	pw := imagick.NewPixelWand()
	pw.SetColor("white")
	if err = collated.SetBackgroundColor(pw); err != nil {
		return err
	}
	if err = collated.ExtentImage(round2Uint(c.Density*c.PageWidthU), round2Uint(c.Density*c.PageHeightU), c.getExtentOffsetX(), c.getExtentOffsetY()); err != nil {
		return err
	}
	if err = collated.SetImageUnits(c.Unit); err != nil {
		return err
	}
	collated.NextImage()
	if err = collated.ExtentImage(round2Uint(c.Density*c.PageWidthU), round2Uint(c.Density*c.PageHeightU), c.getExtentOffsetX(), c.getExtentOffsetY()); err != nil {
		return err
	}

	fmt.Printf("Writing montage to: '%s'...\n", c.OutFilename)
	if err = collated.WriteImages(c.OutFilename, true); err != nil {
		return err
	}
	fmt.Printf("Complete.\n")
	return nil
}
