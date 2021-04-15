package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

func setLicenseKey() error {
	apiKey := os.Getenv("UNICLOUD_METERED_KEY")
	if apiKey == "" {
		return fmt.Errorf("Missing UNICLOUD_METERED_KEY environment variable")
	}

	err := license.SetMeteredKey(apiKey)
	if err != nil {
		return err
	}

	return nil
}

func extractText(inputPath string) error {
	doc, err := document.Open(inputPath)
	if err != nil {
		return err
	}

	extracted := doc.ExtractText()
	fmt.Println(extracted.Text())
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <inputfile.docx>\n", os.Args[0])
		os.Exit(1)
	}

	err := setLicenseKey()
	if err != nil {
		fmt.Printf("Failed to load license: %v\n", err)
		os.Exit(1)
	}

	inputPath := os.Args[1]

	err = extractText(inputPath)
	if err != nil {
		fmt.Printf("Error extracting text from %s : %v\n", inputPath, err)
		os.Exit(1)
	}
}
