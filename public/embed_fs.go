package public

import "embed"

// embeddedFS holds our static content within public, embedded into the binary
// no need to transfer the public folder into the Docker image or with the app binary
//
//go:embed *
var EmbeddedFS embed.FS
