package cmd

import (
	"os"

	"github.com/ohzqq/webp"
	"github.com/spf13/cobra"
)

var opts = webp.Options{
	Quality: webp.DefaultQuality,
	Method:  webp.DefaultMethod,
}

var animate bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webp",
	Short: "commands for converting and animated WebPs",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&opts.Exact, "exact", "e", false, "Exact preserve the exact RGB values in transparent area.")
	rootCmd.PersistentFlags().BoolVarP(&opts.Lossless, "lossless", "l", false, "Lossless indicates whether to use the lossless compression. Lossless will ignore quality")
	rootCmd.PersistentFlags().IntVarP(&opts.Quality, "quality", "q", webp.DefaultQuality, "Quality in the range [0,100]. Quality of 100 implies Lossless. Default is 75.")
	rootCmd.PersistentFlags().IntVarP(&opts.Method, "method", "m", webp.DefaultMethod, "Method is quality/speed trade-off (0=fast, 6=slower-better). Default is 4.")

	rootCmd.PersistentFlags().IntVarP(&opts.LoopCount, "loopcount", "c", 0, "Number of times the animation should loop (0 = infinite).")
	rootCmd.PersistentFlags().IntSliceVarP(&opts.Durations, "durations", "d", []int{webp.DefaultDuration},
		"Durations is duration per frame in milliseconds. Default is 50. If only one is provided, it'll be the duration for every frame.",
	)
	rootCmd.PersistentFlags().IntSliceP(&opts.Disposals, "disposals", []int{0},
		"Disposal methods after frame display (keep or clear). If only one is provided, it'll be the duration for every frame.",
	)
}
