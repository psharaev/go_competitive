package args_parser

import (
	"fmt"
	"github.com/psharaev/go_competitive/internal/models"
	"os"
	"path/filepath"
	"strconv"
)

func ParseArgs(args []string) (models.Args, error) {
	if len(args) != 3 {
		return models.Args{}, fmt.Errorf("usage: program <output_folder> <count_problems> (1-26)")
	}

	outputFolder := os.Args[1]
	countProblems, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return models.Args{}, fmt.Errorf("fail parse count_problems: %w", err)
	}

	if countProblems < 1 {
		return models.Args{}, fmt.Errorf("number of problems must be positive, actual: %d", countProblems)
	}

	// Получаем пути к шаблонам
	execPath, err := os.Executable()
	if err != nil {
		return models.Args{}, fmt.Errorf("fail get exec path: %w", err)
	}
	execDir := filepath.Dir(execPath)

	res := models.Args{
		OutputFolder:     outputFolder,
		CountProblems:    countProblems,
		TemplatePath:     filepath.Join(execDir, "template"),
		TemplateTestPath: filepath.Join(execDir, "template_test"),
	}
	return res, nil
}
