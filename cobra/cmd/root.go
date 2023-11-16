package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/iqra-shams/cobra-example/pkg"
	"github.com/iqra-shams/cobra-example/utils"
	
)

var rootCmd = &cobra.Command{
	Use:   "file",
	Short: "Process a file",
	Run: func(cmd *cobra.Command, args []string) {
		FilePath, _ := cmd.Flags().GetString("file")
		Routines, _ := cmd.Flags().GetInt("routines")
		ProcessFile(FilePath, Routines)
	},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	// rootCmd.AddCommand(fileCmd)
	rootCmd.Flags().IntP("routines", "r", 3, "Number of routines")
	rootCmd.Flags().String("file", "docs/newFile.txt", "to read file")
	// Add other commands as needed
}
func ProcessFile(filePath string, routines int) {
	channal := make(chan pkg.Summary)
	FileData:=utils.ReadFile(filePath)
	chunk := len(FileData) / routines
	startIndex := 0
	endIndex := chunk
	for iterations := 0; iterations < routines; iterations++ {
		go pkg.Counts(FileData[startIndex:endIndex], channal)
		startIndex = endIndex
		endIndex += chunk
	}
	utils.GetChunksSummary(channal , routines)
	fmt.Printf("total number of lines : %d \n", utils.Lines)
	fmt.Printf("total number of words : %d \n", utils.Words)
	fmt.Printf("total number of vowels: %d \n", utils.Vowels)
	fmt.Printf("total number of puncuations : %d \n", utils.Puncuations)
}
