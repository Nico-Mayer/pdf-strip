package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OpenFiles() []string {
	filePath, _ := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
		//DefaultDirectory:           string,
		//DefaultFilename  :          string,
		Title: "Choose PDF file",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "PDF (*.pdf)",
				Pattern:     "*.pdf",
			},
		},
		ShowHiddenFiles:            false,
		CanCreateDirectories:       false,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: false,
	})
	return filePath
}

func (a *App) Metadata(path string) fs.FileInfo {

	fileInfo, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println("File metadata:", fileInfo)
	return fileInfo
}
