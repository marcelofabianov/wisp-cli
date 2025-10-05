package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wisp-cli [type]",
	Short: "wisp-cli é uma ferramenta para explorar os value objects do pacote wisp.",
	Long:  `Uma ferramenta de linha de comando para visualizar conceitos, exemplos e documentação de cada tipo do pacote wisp.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		fmt.Printf("Buscando informações para o tipo: %s\n", args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
