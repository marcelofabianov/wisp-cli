package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var contentFS embed.FS

func SetContentFS(fs embed.FS) {
	contentFS = fs
}

var rootCmd = &cobra.Command{
	Use:   "wisp-cli [type]",
	Short: "wisp-cli é uma ferramenta para explorar os value objects do pacote wisp.",
	Long:  `Uma ferramenta de linha de comando para visualizar conceitos, exemplos e documentação de cada tipo do pacote wisp.`,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return listAvailableTypes(), cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Nenhum tipo especificado. Tipos disponíveis:")
			for _, t := range listAvailableTypes() {
				fmt.Printf("  - %s\n", t)
			}
			fmt.Println("\nUse 'wisp-cli [type]' para ver os detalhes.")
			return
		}
		typeName := strings.ToLower(args[0])
		displayTypeInfo(typeName)
	},
}

func displayTypeInfo(typeName string) {
	basePath := filepath.Join("content", typeName)

	concept, err := contentFS.ReadFile(filepath.Join(basePath, "concept.md"))
	if err != nil {
		fmt.Printf("Erro: Tipo '%s' não encontrado.\n", typeName)
		fmt.Println("Tente um dos tipos disponíveis:", strings.Join(listAvailableTypes(), ", "))
		os.Exit(1)
	}
	fmt.Println("--- Conceito ---")
	fmt.Println(string(concept))

	example, err := contentFS.ReadFile(filepath.Join(basePath, "example.go.txt"))
	if err == nil {
		fmt.Println("\n--- Exemplo de Uso ---")
		fmt.Println(string(example))
	}

	url, err := contentFS.ReadFile(filepath.Join(basePath, "url.txt"))
	if err == nil {
		fmt.Println("\n--- Documentação ---")
		fmt.Printf("Para mais detalhes, veja: %s\n", strings.TrimSpace(string(url)))
	}
}

func listAvailableTypes() []string {
	var types []string
	entries, err := fs.ReadDir(contentFS, "content")
	if err != nil {
		return nil
	}
	for _, entry := range entries {
		if entry.IsDir() {
			types = append(types, entry.Name())
		}
	}
	return types
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
