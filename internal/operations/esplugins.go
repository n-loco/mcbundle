package operations

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/n-loco/mcbuild/internal/jsonst"
	"github.com/n-loco/mcbuild/internal/mcmanifest"
)

type packageJSON struct {
	Version *jsonst.SemVer `json:"version"`
}

func mcNativeModResolverPlugin(buildModCtx *buildModuleContext) api.Plugin {
	return api.Plugin{
		Name: "Minecraft Native Module Resolver",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(
				api.OnResolveOptions{
					Filter: `^@minecraft\/(?:common|debug-utilities|diagnostics|server(?:|-admin|-editor|-gametest|-net|-ui))$`,
				},
				func(args api.OnResolveArgs) (api.OnResolveResult, error) {
					err := findNativeModuleVersion(&build, &args, buildModCtx)
					if err != nil {
						return api.OnResolveResult{}, err
					}
					return api.OnResolveResult{External: true}, nil
				},
			)
		},
	}
}

func findNativeModuleVersion(
	build *api.PluginBuild,
	args *api.OnResolveArgs,
	buildModCtx *buildModuleContext,
) error {
	if _, ok := buildModCtx.scriptDependencyMap[args.Path]; ok {
		return nil
	}

	packageJSONPath := filepath.Join(args.Path, "package.json")
	result := build.Resolve(packageJSONPath, api.ResolveOptions{
		Kind:       args.Kind,
		ResolveDir: args.ResolveDir,
	})

	if len(result.Errors) > 0 {
		return &NativeModuleError{NativeModule: args.Path}
	}

	var packageJSONData []byte
	var readErr error
	if packageJSONData, readErr = os.ReadFile(result.Path); readErr != nil {
		return &NativeModuleError{NativeModule: args.Path}
	}

	var packageJSON packageJSON
	if unmarshErr := json.Unmarshal(packageJSONData, &packageJSON); unmarshErr != nil {
		return &NativeModuleError{NativeModule: args.Path}
	}

	buildModCtx.scriptDependencyMap[args.Path] = mcmanifest.Dependency{
		ModuleName: args.Path,
		Version:    packageJSON.Version,
	}

	return nil
}
