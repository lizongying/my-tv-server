package statics

import (
	"embed"
	_ "embed"
)

//go:embed mytv
var Mytv embed.FS
