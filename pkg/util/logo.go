//go:build !windows
// +build !windows

package util

// Logo is the ASCII logo shown on startup.
var Logo = `
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;125m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;88m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;88m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;125m \x1b[0;00m\x1b[38;5;125m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;161m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m0\x1b[0;00m\x1b[38;5;196m:\x1b[0;00m\x1b[38;5;196m,\x1b[0;00m\x1b[38;5;22m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;196mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;225m@\x1b[0;00m\x1b[38;5;182m0\x1b[0;00m\x1b[38;5;175mf\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162mt\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;204m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;224m.\x1b[0;00m\x1b[38;5;224m:\x1b[0;00m\x1b[38;5;224m1\x1b[0;00m\x1b[38;5;224mL\x1b[0;00m\x1b[38;5;224m0\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;188m@\x1b[0;00m\x1b[38;5;188m8\x1b[0;00m\x1b[38;5;181mG\x1b[0;00m\x1b[38;5;175mC\x1b[0;00m\x1b[38;5;168mf\x1b[0;00m\x1b[38;5;162mt\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;168mi\x1b[0;00m\x1b[38;5;175m1\x1b[0;00m\x1b[38;5;182m;\x1b[0;00m\x1b[38;5;30m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;132m \x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;175m \x1b[0;00m\x1b[38;5;218m \x1b[0;00m\x1b[38;5;219m \x1b[0;00m\x1b[38;5;217m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;217m \x1b[0;00m\x1b[38;5;217m \x1b[0;00m\x1b[38;5;218m \x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;175m \x1b[0;00m\x1b[38;5;218m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;224m:\x1b[0;00m\x1b[38;5;188mf\x1b[0;00m\x1b[38;5;182m0\x1b[0;00m\x1b[38;5;175mC\x1b[0;00m\x1b[38;5;168mL\x1b[0;00m\x1b[38;5;198mf\x1b[0;00m\x1b[38;5;161mt\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161mi\x1b[0;00m\x1b[38;5;161m;\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;175m1\x1b[0;00m\x1b[38;5;188mt\x1b[0;00m\x1b[38;5;225mL\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;224mC\x1b[0;00m\x1b[38;5;224mt\x1b[0;00m\x1b[38;5;188mi\x1b[0;00m\x1b[38;5;224m:\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161m1\x1b[0;00m\x1b[38;5;161mi\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;168mi\x1b[0;00m\x1b[38;5;182mt\x1b[0;00m\x1b[38;5;224mC\x1b[0;00m\x1b[38;5;196m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;225m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162m1\x1b[0;00m\x1b[38;5;162mi\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;33m \x1b[0;00m\x1b[38;5;196m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;224mG\x1b[0;00m\x1b[38;5;224m8\x1b[0;00m\x1b[38;5;224m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;125m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;161m.\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m;\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m:\x1b[0;00m\x1b[38;5;162m,\x1b[0;00m\x1b[38;5;162m.\x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;162m \x1b[0;00m\x1b[38;5;161m \x1b[0;00m\x1b[38;5;161m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;197m \x1b[0;00m\x1b[38;5;198m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;196m \x1b[0;00m\x1b[38;5;203m \x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;211m \x1b[0;00m\x1b[38;5;175m \x1b[0;00m\x1b[38;5;218m \x1b[0;00m\x1b[38;5;217m \x1b[0;00m\x1b[38;5;145m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;217m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231mi\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231mC\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231m0\x1b[0;00m\x1b[38;5;231m@\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231mL\x1b[0;00m\x1b[38;5;231m8\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mt\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231m1\x1b[0;00m\x1b[38;5;231mf\x1b[0;00m\x1b[38;5;231mG\x1b[0;00m\x1b[38;5;231m;\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m.\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m:\x1b[0;00m\x1b[38;5;231m,\x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;231m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m\x1b[38;5;16m \x1b[0;00m
`
