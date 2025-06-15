package esfiles

import (
	"encoding/json"
	"os"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projctx"
)

var mcNativeModRegExpS = `^@minecraft\/(?:common|debug-utilities|diagnostics|server(?:|-admin|-editor|-gametest|-net|-ui))$`

type packageJSON struct {
	Version *jsonst.SemVer `json:"version"`
}

func mcNativeModResolverPlugin(modCtx *projctx.ModuleContext) esbuild.Plugin {
	return esbuild.Plugin{
		Name: "Minecraft Native Module Resolver",
		Setup: func(build esbuild.PluginBuild) {
			build.OnResolve(
				esbuild.OnResolveOptions{
					Filter: mcNativeModRegExpS,
				},
				func(args esbuild.OnResolveArgs) (esbuild.OnResolveResult, error) {
					err := findNativeModuleVersion(&build, &args, modCtx)
					if err != nil {
						return esbuild.OnResolveResult{}, err
					}
					return esbuild.OnResolveResult{External: true}, nil
				},
			)
		},
	}
}

func findNativeModuleVersion(
	build *esbuild.PluginBuild,
	args *esbuild.OnResolveArgs,
	modCtx *projctx.ModuleContext,
) error {
	if hasDep := modCtx.HasScriptDependency(args.Path); hasDep {
		return nil
	}

	packageJSONImport := strings.Join([]string{args.Path, "package.json"}, "/")
	result := build.Resolve(packageJSONImport, esbuild.ResolveOptions{
		Kind:       args.Kind,
		ResolveDir: args.ResolveDir,
	})

	if len(result.Errors) > 0 {
		return &resolveNativeModErr{NativeModule: args.Path}
	}

	var packageJSONData []byte
	var readErr error
	if packageJSONData, readErr = os.ReadFile(result.Path); readErr != nil {
		return &resolveNativeModErr{NativeModule: args.Path}
	}

	var packageJSON packageJSON
	if unmarshErr := json.Unmarshal(packageJSONData, &packageJSON); unmarshErr != nil {
		return &resolveNativeModErr{NativeModule: args.Path}
	}

	modCtx.AddScriptDependency(args.Path, packageJSON.Version)

	return nil
}
