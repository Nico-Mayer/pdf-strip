package main

import (
	"context"
	"os"
	goruntime "runtime"

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
		return ""
	}
	return dirPath

}

func (a *App) Encrypt(path, userPW, ownerPW string) error {
	conf := model.NewAESConfiguration(userPW, ownerPW, 256)
	err := api.EncryptFile(path, "", conf)

	if err != nil {
		return err
	}
	return nil
}

func (a *App) Decrypt(path, pw string, owner bool) error {
	var conf *model.Configuration

	if owner {
		conf = model.NewAESConfiguration("", pw, 256)
	} else {
		conf = model.NewAESConfiguration(pw, "", 256)
	}

	err := api.DecryptFile(path, "", conf)
	if err != nil {
		return err
	}
	return nil
}

type PDFInfo struct {
	Exists    bool
	Encrypted bool
	Info      []string
}

func (a *App) GetPDFInfo(path string) (PDFInfo, error) {
	ret := PDFInfo{
		Exists:    true,
		Encrypted: false,
		Info:      nil,
	}
	conf := model.NewDefaultConfiguration()

	//Check if file exists
	f, err := os.Open(path)
	if err != nil {
		ret.Exists = false
	}
	defer f.Close()
	// Validate file
	err = api.ValidateFile(path, conf)
	if err != nil {
		ret.Encrypted = true
	}
	// Get file info
	info, err := api.InfoFile(path, nil, conf)
	if err != nil {
		ret.Info = nil
	}
	ret.Info = info
	return ret, nil
}

func (a *App) GetOs() string {
	return goruntime.GOOS
}
