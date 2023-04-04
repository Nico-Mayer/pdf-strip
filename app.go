package main

import (
	"context"
	"fmt"
	"os"
	goruntime "runtime"

	"github.com/unidoc/unipdf/v3/model"
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

func (a *App) Encrypt(path, userPW, ownerPW string) bool {
	return true
}

func (a *App) Decrypt(path string, pw string, owner bool) bool {

	return true
}

type PdfProperties struct {
	IsEncrypted bool
	CanView     bool // Is the document viewable without password?
	NumPages    int
}

func (a *App) GetPDFInfo(path string) (*PdfProperties, error) {
	ret := PdfProperties{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return nil, err
	}
	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return nil, err
	}
	ret.IsEncrypted = isEncrypted
	ret.CanView = true

	// Try decrypting with an empty one.
	if isEncrypted {
		auth, err := pdfReader.Decrypt([]byte(""))
		if err != nil {
			return nil, err
		}
		ret.CanView = auth
		return &ret, nil
	}
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return nil, err
	}
	ret.NumPages = numPages

	return &ret, nil
}

func (a *App) GetOs() string {
	return goruntime.GOOS
}
