package cli

import (
	"github.com/mcbundle/mcbundle/assets"
	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/sysi"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

func versionCmd(*argvIterator, alert.Diagnostic) {
	txtui.Printf(txtui.UIPartOut, "%s\n", assets.ProgramVersion)
}

const wikiURL = "https://github.com/n-loco/mcbundle/wiki"

func helpCmd(argv *argvIterator, diagnostic alert.Diagnostic) {
	if argv.hasNext() {
		var option = argv.consume()
		if option == "--browser" {
			if sysi.HasOpenSupport() {
				sysi.Open(wikiURL)
			} else {
				diagnostic.AppendError(alert.AlertF("unsupported option: %s", option))
			}
			return
		} else {
			diagnostic.AppendWarning(alert.AlertF("unknown option: %s", option))
		}
	}

	txtui.Print(txtui.UIPartOut, "Usage: mcbundle [command] <options>\n")
	txtui.Print(txtui.UIPartOut, `
Commands:
    build, bundle              Build the project.
    dev, local-deploy          Copy the built project to com.mojang.
    pack, dist                 Package the project as a .mcaddon or .mcpack.
    help, --help, -h           Show this help message.
    version, --version, -v     Show the current version.

`,
	)
	txtui.Printf(txtui.UIPartOut, "For more information, visit: %s\n", wikiURL)
	if sysi.HasOpenSupport() {
		txtui.Print(txtui.UIPartOut, "Or run: mcbundle help --browser\n")
	}
}
