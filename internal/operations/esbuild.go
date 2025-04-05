package operations

import (
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/n-loco/mcbuild/internal/rcontext"
)

func esbuild(ctx *rcontext.Context, buildModCtx *buildModuleContext) error {
	mainFile := filepath.Join(buildModCtx.sourcePath, "main")
	if stat, err := os.Stat(mainFile + ".ts"); err == nil && !stat.IsDir() {
		mainFile += ".ts"
	} else if stat, err := os.Stat(mainFile + ".js"); err == nil && !stat.IsDir() {
		mainFile += ".js"
	} else {
		return &MainFileNotFoundError{ExpectedFiles: []string{"main.ts", "main.ts"}, ModuleType: buildModCtx.recipeModule.Type}
	}

	outFile := filepath.Join(buildModCtx.buildPath, "scripts", buildModCtx.recipeModule.Type.String()+".js")

	sourcemap := api.SourceMapNone
	if !ctx.Release {
		sourcemap = api.SourceMapLinked
	}

	result := api.Build(api.BuildOptions{
		// General
		Plugins: []api.Plugin{
			mcNativeModResolverPlugin(buildModCtx),
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
		Sourcemap:         sourcemap,
		MinifyWhitespace:  ctx.Release,
		MinifySyntax:      ctx.Release,
		MinifyIdentifiers: ctx.Release,
	})

	if len(result.Errors) > 0 {
		return &ESBuildErrorWrapper{Messages: result.Errors}
	}

	return nil
}
