package weather

import _ "embed"

// Keep all documentation in this file as embeds in order to easily
// support compiles to other target languages by simply changing the
// language identifier before compilation.

//go:embed text/en/weather.md
var _weather string

//go:embed text/en/get.md
var _get string
