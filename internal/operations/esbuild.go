package operations

import (
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/projctx"
)

func esbuild(modCtx *projctx.ModuleContext) (diagnostic *alert.Diagnostic) {
	recipeModule := modCtx.RecipeModule

	mainFile := filepath.Join(modCtx.ModSourcePath, "main")
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

	outFile := filepath.Join(modCtx.PackDistDir, "scripts", recipeModule.Type.String()+".js")

	sourcemap := api.SourceMapNone
	if !modCtx.Release {
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
		MinifyWhitespace:  modCtx.Release,
		MinifySyntax:      modCtx.Release,
		MinifyIdentifiers: modCtx.Release,

		// Debug
		Sourcemap:      sourcemap,
		SourceRoot:     filepath.Join(modCtx.PackDistDir, "scripts"),
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
