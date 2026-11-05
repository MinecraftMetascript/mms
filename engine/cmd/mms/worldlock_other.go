//go:build !unix && !windows

package main

// worldInUse cannot tell on this system, so it says no.
func worldInUse(world string) bool { return false }
