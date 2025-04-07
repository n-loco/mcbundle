package operations

import (
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
)

func esbuild(modCtx *moduleContext) error {
	recipeModule := modCtx.recipeModule

	mainFile := filepath.Join(modCtx.modSourcePath, "main")
	if stat, err := os.Stat(mainFile + ".ts"); err == nil && !stat.IsDir() {
		mainFile += ".ts"
	} else if stat, err := os.Stat(mainFile + ".js"); err == nil && !stat.IsDir() {
		mainFile += ".js"
	} else {
		return &MainFileNotFoundError{ExpectedFiles: []string{"main.ts", "main.ts"}, ModuleType: recipeModule.Type}
	}

	outFile := filepath.Join(modCtx.packDistDir, "scripts", recipeModule.Type.String()+".js")

	sourcemap := api.SourceMapNone
	if !modCtx.release {
		sourcemap = api.SourceMapLinked
	}

	result := api.Build(api.BuildOptions{
		// General
		Plugins: []api.Plugin{
			mcNativeModResolverPlugin(modCtx),
		},
		Write:  true,
		Bundle: true,

		// Input
		EntryPoints: []string{mainFile},
		MainFields:  []string{"minecraft_server", "minecraft", "module", "main"},

		// Output
		Outfile:           outFile,
		Platform:          api.PlatformNeutral,
		Target:            api.ES2020,
		MinifyWhitespace:  modCtx.release,
		MinifySyntax:      modCtx.release,
		MinifyIdentifiers: modCtx.release,

		// Debug
		Sourcemap:      sourcemap,
		SourceRoot:     filepath.Join(modCtx.packDistDir, "scripts"),
		SourcesContent: api.SourcesContentExclude,
	})

	if len(result.Errors) > 0 {
		return &ESBuildErrorWrapper{Messages: result.Errors}
	}

	return nil
}
