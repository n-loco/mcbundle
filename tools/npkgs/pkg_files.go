package main

import _ "embed"

//go:embed package.template.json
var templatePackage string

//go:embed package.template.publish.json
var templateReleasePackage string

//go:embed THIRD_PARTY.md
var thirdParty string
