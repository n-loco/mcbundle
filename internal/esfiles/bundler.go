package esfiles

import (
	"os"
	"path/filepath"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx"
)

type BundlingContext byte

const (
	BundlingContextGeneric BundlingContext = iota
	BundlingContextServerModule
)

func (bctx BundlingContext) String() string {
	switch bctx {
	case BundlingContextGeneric:
		return "generic"
	case BundlingContextServerModule:
		return "server module"
	}
	return ""
}

type JSBundlerOptions struct {
	BundlingContext    BundlingContext
	StripDebug         bool
	WorkDir            string
	MainFields         []string
	SourceDir          string
	PossibleEntryFiles []string
	OutPutFile         string

	plugins []esbuild.Plugin
}

func (opts *JSBundlerOptions) AddNativeResolverPlugin(modCtx *projctx.ModuleContext) {
	opts.plugins = append(opts.plugins, mcNativeModResolverPlugin(modCtx))
}

func JSBundler(opts *JSBundlerOptions) (diagnostic *alert.Diagnostic) {
	var entryPath string
	var success uint8

	for _, entryName := range opts.PossibleEntryFiles {
		success += getEntryAttempt(opts.SourceDir, entryName, &entryPath)
	}

	switch {
	case success == 0:
		diagnostic = diagnostic.AppendError(&entryFileNotFoundErrAlert{
			ExpectedFiles:   opts.PossibleEntryFiles,
			BundlingContext: opts.BundlingContext,
		})
	case success > 1:
		diagnostic = diagnostic.AppendError(&tooManyEntriesError{
			BundlingContext: opts.BundlingContext,
			ExpectedFiles:   opts.PossibleEntryFiles,
		})
	}

	if diagnostic.HasErrors() {
		return
	}

	var sourcemap = esbuild.SourceMapNone
	if !opts.StripDebug {
		sourcemap = esbuild.SourceMapLinked
	}

	var sourcemapPath, _ = filepath.Rel(opts.WorkDir, filepath.Join(opts.OutPutFile, ".."))

	var result = esbuild.Build(esbuild.BuildOptions{
		// General
		Plugins: opts.plugins,
		Write:   true,
		Bundle:  true,

		// Input
		EntryPoints: []string{entryPath},
		MainFields:  opts.MainFields,

		// Output
		Outfile:           opts.OutPutFile,
		Platform:          esbuild.PlatformNeutral,
		Target:            esbuild.ES2020,
		MinifyWhitespace:  opts.StripDebug,
		MinifySyntax:      opts.StripDebug,
		MinifyIdentifiers: opts.StripDebug,

		// Debug
		Sourcemap:      sourcemap,
		SourceRoot:     sourcemapPath,
		SourcesContent: esbuild.SourcesContentExclude,
	})

	for _, eswarn := range result.Warnings {
		wwarn := esbuildWrapperAlert(eswarn)
		diagnostic = diagnostic.AppendError(&wwarn)
	}

	for _, eserror := range result.Errors {
		werror := esbuildWrapperAlert(eserror)
		diagnostic = diagnostic.AppendError(&werror)
	}

	return
}

func getEntryAttempt(
	sourceDir string,
	entryName string,
	outPut *string,
) uint8 {
	var entryPath = filepath.Join(sourceDir, entryName)

	if entryFile, _ := os.Stat(entryPath); entryFile != nil && !entryFile.IsDir() {
		*outPut = entryPath
		return 1
	}

	return 0
}
