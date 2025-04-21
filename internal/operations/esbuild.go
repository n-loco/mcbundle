package operations

import (
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/n-loco/bpbuild/internal/alert"
)

func esbuild(modCtx *moduleContext) (diagnostic *alert.Diagnostic) {
	recipeModule := modCtx.recipeModule

	mainFile := filepath.Join(modCtx.modSourcePath, "main")
	if stat, err := os.Stat(mainFile + ".ts"); err == nil && !stat.IsDir() {
		mainFile += ".ts"
	} else if stat, err := os.Stat(mainFile + ".js"); err == nil && !stat.IsDir() {
		mainFile += ".js"
	} else {
		diagnostic = diagnostic.AppendError(&MainFileNotFoundErrAlert{
			ExpectedFiles: []string{"main.ts", "main.js"},
			ModuleType:    recipeModule.Type,
		})
		return
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

	for _, eswarn := range result.Warnings {
		wwarn := ESBuildWrapperAlert(eswarn)
		diagnostic = diagnostic.AppendError(&wwarn)
	}

	for _, eserror := range result.Errors {
		werror := ESBuildWrapperAlert(eserror)
		diagnostic = diagnostic.AppendError(&werror)
	}

	return
}
