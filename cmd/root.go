package cmd

import (
	"fmt"
	"github.com/bamItsCam/cardtage/internal/cardtage"
	"github.com/spf13/cobra"
	"gopkg.in/gographics/imagick.v2/imagick"
)

var (
	// Used for flags.
	density          float64
	border           float64
	pageDimensions   string
	marginDimensions string
	cardDimensions   string
	rootCmd          = &cobra.Command{
		Use:   "cardtage",
		Short: "Cardtage tiles images you give it onto an output pdf",
		Long: `Cardtage does one thing and it does that one thing really-mediocerly.
Give it files, tell it how far to space them, and tell it the size of page
you'd like to make (like 8.5x11) and tada! A montage of those images.
Useful for printing diy playing cards.`,
		Args: cobra.ExactArgs(2),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			pageWidthU, pageHeightU, err := cardtage.Res2WxH(pageDimensions)
			if err != nil {
				return fmt.Errorf("error: %s", err)
			}
			marginWidthU, marginHeightU, err := cardtage.Res2WxH(marginDimensions)
			if err != nil {
				return fmt.Errorf("error: %s", err)
			}
			cardWidthU, cardHeightU, err := cardtage.Res2WxH(cardDimensions)
			if err != nil {
				return fmt.Errorf("error: %s", err)
			}
			cardtage := cardtage.CardtageImpl{
				InFileGlob:        args[0],
				OutFilename:       args[1],
				Density:           density,
				Unit:              imagick.RESOLUTION_PIXELS_PER_INCH,
				CardWidthU:        cardWidthU,
				CardHeightU:       cardHeightU,
				PageWidthU:        pageWidthU,
				PageHeightU:       pageHeightU,
				PageMarginWidthU:  marginWidthU,
				PageMarginHeightU: marginHeightU,
				CardBorderU:       border,
			}
			if err := cardtage.Generate(); err != nil {
				return err
			}
			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Float64VarP(&density, "density", "d", 100, "PPI/resolution to use when ingesting and exporting")
	rootCmd.Flags().Float64VarP(&border, "border", "b", .01, "in inches, border size around each card")
	rootCmd.Flags().StringVarP(&pageDimensions, "page", "p", "8.5x11", "in inches, size of the page you'd like to output")
	rootCmd.Flags().StringVarP(&marginDimensions, "margin", "m", "0.25x0.25", "in inches, the margin that should be respected")
	rootCmd.Flags().StringVarP(&cardDimensions, "card", "c", "2.5x3.5", "in inches, the size of the card/tile")

}
