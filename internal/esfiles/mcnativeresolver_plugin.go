package esfiles

import (
	"encoding/json"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projctx"
)

var mcNativeModRegExpS = `^@minecraft\/(?:common|debug-utilities|diagnostics|server(?:|-admin|-editor|-gametest|-net|-ui))$`

type packageJSON struct {
	Version *jsonst.SemVer `json:"version"`
}

func minecraftNativeModuleResolver(modCtx *projctx.ModuleContext) esbuild.Plugin {
	return esbuild.Plugin{
		Name: "go:github.com/mcbundle/mcbundle/internal/esfiles/minecraftNativeModuleResolver",
		Setup: func(build esbuild.PluginBuild) {
			mcNativeModResolverSetup(modCtx, build)
		},
	}
}

func mcNativeModResolverSetup(modCtx *projctx.ModuleContext, build esbuild.PluginBuild) {
	build.OnResolve(
		esbuild.OnResolveOptions{
			Filter: mcNativeModRegExpS,
		},
		func(args esbuild.OnResolveArgs) (esbuild.OnResolveResult, error) {
			var err = findNativeModuleVersion(&build, &args, modCtx)
			if err != nil {
				return esbuild.OnResolveResult{}, err
			}
			return esbuild.OnResolveResult{External: true}, nil
		},
	)
}

func findNativeModuleVersion(
	build *esbuild.PluginBuild, args *esbuild.OnResolveArgs, modCtx *projctx.ModuleContext,
) error {
	if hasDep := modCtx.HasScriptDependency(args.Path); hasDep {
		return nil
	}

	var resolveErr = &resolveNativeModErr{NativeModule: args.Path}

	var packageJSONImport = args.Path + "/package.json"
	var result = build.Resolve(packageJSONImport, esbuild.ResolveOptions{
		Kind:       args.Kind,
		ResolveDir: args.ResolveDir,
	})

	if len(result.Errors) > 0 {
		return resolveErr
	}

	var packageJSONPath = result.Path

	var packageJSONData, readErr = vreadFile(packageJSONPath)
	if readErr != nil {
		return resolveErr
	}

	var packageJSON packageJSON
	if unmarshErr := json.Unmarshal(packageJSONData, &packageJSON); unmarshErr != nil {
		return resolveErr
	}

	modCtx.AddScriptDependency(args.Path, packageJSON.Version)

	return nil
}
