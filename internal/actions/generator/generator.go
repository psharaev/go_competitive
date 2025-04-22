package generator

import (
	"fmt"
	"github.com/psharaev/go_competitive/internal/models"
	"log"
	"os"
	"path/filepath"
)

func Generate(args models.Args) error {
	templateContent, err := os.ReadFile(args.TemplatePath)
	if err != nil {
		return fmt.Errorf("fail read template (%s): %w", args.TemplatePath, err)
	}

	testTemplateContent, err := os.ReadFile(args.TemplateTestPath)
	if err != nil {
		return fmt.Errorf("fail read template_test (%s): %w", args.TemplateTestPath, err)
	}

	err = os.MkdirAll(args.OutputFolder, 0755)
	if err != nil {
		return fmt.Errorf("fail create output directory (%s): %w", args.OutputFolder, err)
	}

	for i := 1; i <= args.CountProblems; i++ {
		char := 'a' + i - 1
		dirName := fmt.Sprintf("%c%d", char, i)
		fullDirPath := filepath.Join(args.OutputFolder, dirName)

		err = os.Mkdir(fullDirPath, 0755)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}

		err = createFile(fullDirPath, dirName, ".go", templateContent)
		if err != nil {
			return err
		}
		err = createFile(fullDirPath, dirName, "_test.go", testTemplateContent)
		if err != nil {
			return err
		}
	}

	return nil
}

func createFile(dirPath, baseName, suffix string, content []byte) error {
	fileName := filepath.Join(dirPath, baseName+suffix)
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		return fmt.Errorf("error writing file (%s): %w", fileName, err)
	}
	return nil
}
