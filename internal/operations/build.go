package operations

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/esfiles"
	"github.com/mcbundle/mcbundle/internal/mcfiles"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func BuildProject(projCtx *projctx.ProjectContext, release bool) {
	var projType = projCtx.Recipe.Type

	if projType == projfiles.RecipeTypeAddOn {
		var bpCtx, rpCtx = projCtx.AddonContext(release)

		buildPack(&bpCtx)
		buildPack(&rpCtx)
	} else {
		var packCtx = projCtx.PackContext(release)
		buildPack(&packCtx)
	}
}

func buildPack(packCtx *projctx.PackContext) {
	var diagnostic = packCtx.Diagnostic

	var projRecipe = packCtx.Recipe
	var packType = packCtx.PackType

	var buildPath = packCtx.PackDistDir
	if _, err := os.Stat(buildPath); err == nil {
		os.RemoveAll(buildPath)
	}

	var foundDeps []mcfiles.Dependency
	var builtModules []mcfiles.Module

	for _, recipeModule := range projRecipe.Modules {
		if recipeModule.BelongsTo() != packType {
			continue
		}

		var modCtx = packCtx.ModuleContext(&recipeModule)

		var mod = buildModule(&modCtx, diagnostic)

		builtModules = append(builtModules, mod)
	}

	if diagnostic.HasErrors() {
		return
	}

	foundDeps = packCtx.ScriptDependencies()

	var packIconPath, _ = packCtx.Rel(
		filepath.Join(packCtx.PackDistDir, "pack_icon.png"),
	)

	var packIconFile, packIconErr = os.Create(packIconPath)
	if packIconErr != nil {
		diagnostic.AppendError(alert.WrappGoError(packIconErr))
	} else {
		packIconFile.Write(packCtx.PackIcon)
		packIconFile.Close()
	}

	writeManifest(packCtx, builtModules, foundDeps)
}

func buildModule(modCtx *projctx.ModuleContext, diagnostic alert.Diagnostic) (mod mcfiles.Module) {
	var recipeModule = modCtx.RecipeModule

	switch recipeModule.Type {
	case projfiles.ModuleTypeData:
		fallthrough
	case projfiles.ModuleTypeResources:
		{
			copyDataToBuild(modCtx.ModSourcePath, modCtx.PackDistDir, diagnostic)
		}
	case projfiles.ModuleTypeServer:
		{
			var bundleOpts = esfiles.JSBundlerOptions{
				Diagnostic:      diagnostic,
				BundlingContext: esfiles.BundlingContextServerModule,
				StripDebug:      modCtx.Release,
				WorkDir:         modCtx.WorkDir,
				MainFields:      []string{"minecraft_server", "minecraft", "module", "main"},
				SourceDir:       modCtx.ModSourcePath,
				PossibleEntryFiles: []string{
					"main.ts", "main.mts", "main.cts",
					"main.js", "main.mjs", "main.cjs",
				},
				OutPutFile: filepath.Join(modCtx.PackDistDir, "scripts", "server.js"),
			}

			bundleOpts.AddNativeResolverPlugin(modCtx)

			esfiles.JSBundler(&bundleOpts)

			if !diagnostic.HasErrors() {
				mod.Entry = "scripts/server.js"
				mod.Language = "javascript"
			}
		}
	default:
		panic("invalid module")
	}

	if diagnostic.HasErrors() {
		return
	}

	mod.UUID = recipeModule.UUID
	mod.Version = recipeModule.Version
	mod.Type = mcfiles.ModuleTypeFromRecipeModuleType(recipeModule.Type)

	return
}

func copyDataToBuild(from string, to string, diagnostic alert.Diagnostic) {
	diagnostic.AppendError(alert.WrappGoError(os.CopyFS(to, os.DirFS(from))))
}
