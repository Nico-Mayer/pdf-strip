package main

import (
	"context"
	"fmt"
	"io/fs"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
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
	filePath, err := runtime.OpenMultipleFilesDialog(a.ctx, runtime.OpenDialogOptions{
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
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return filePath
}

func (a *App) OpenDir() string {
	dirPath, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		//DefaultDirectory:           string,
		//DefaultFilename  :          string,
		Title: "Choose Directory",
		/* Filters: []runtime.FileFilter{
			{
				DisplayName: "PDF (*.pdf)",
				Pattern:     "*.pdf",
			},
		}, */
		ShowHiddenFiles:            false,
		CanCreateDirectories:       true,
		ResolvesAliases:            true,
		TreatPackagesAsDirectories: false,
	})

	if err != nil {
		fmt.Println(err)
		return ""
	}
	return dirPath

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

func (a *App) Encrypt(path string, userPW string, ownerPW string) bool {
	conf := model.NewAESConfiguration(userPW, ownerPW, 256)
	err := api.EncryptFile(path, "", conf)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (a *App) Decrypt(path string, pw string, owner bool) bool {
	var conf *model.Configuration

	if owner {
		conf = model.NewAESConfiguration("", pw, 256)
	} else {
		conf = model.NewAESConfiguration(pw, "", 256)
	}

	err := api.DecryptFile(path, "", conf)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func (a *App) IsEncrypted(path string) bool {
	config := model.NewDefaultConfiguration()
	err := api.ValidateFile(path, config)

	if err != nil {
		fmt.Println("Is Encrypted", err)
		return true
	}
	return false

}
