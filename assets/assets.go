package assets

import _ "embed"

//go:embed program_version.txt
var ProgramVersion string

//go:embed default_pack_icon.png
var DefaultPackIcon []byte
