package cmd

import (
	"fmt"
	"log"

	"github.com/cybermind-nick/modulenator/project/golang"
	"github.com/spf13/cobra"
)

var Path string
var GitHub string

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "'go' creates a go project or module",
	Long:  `'go' creates a go project or module`,
	Run: func(cmd *cobra.Command, args []string) {

		if Path == "" {
			fmt.Println("generating go project")

			fmt.Println(GitHub)
			// err := golang.CheckForProj()
			// if err != nil {
			// 	log.Fatal(err)
			// }

			if err := golang.NewProject(GitHub); err != nil {
				log.Fatal(err)
			}

			if err := golang.GenerateMain(); err != nil {
				log.Fatal(err)
			}
			fmt.Println("Successfully Generated Go Project...")

		} else {
			fmt.Println("generating go project at specified path")

			fmt.Println("\n", &Path)

			if err := golang.NewProjectWithPath(GitHub, Path); err != nil {
				log.Fatal(err)
			}
			if err := golang.GenerateMain(); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	goCmd.Flags().StringVarP(&GitHub, "github", "g", "", "Set the github account to set the module to")
	goCmd.Flags().StringVarP(&Path, "path", "p", "", "Add alternative path for the project, must be an absolute path")

	goCmd.MarkFlagRequired("github")

	rootCmd.AddCommand(goCmd)

}
