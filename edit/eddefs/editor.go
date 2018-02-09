package eddefs

import (
	"github.com/elves/elvish/daemon"
	"github.com/elves/elvish/eval"
)

type Editor interface {
	// ReadLine reads a line interactively.
	ReadLine() (string, error)
	// Close releases resources used by the editor.
	Close()

	// Evaler returns the Evaler associated with the Editor.
	Evaler() *eval.Evaler
	// Daemon returns the daemon client associated with the Editor.
	Daemon() *daemon.Client

	// Buffer returns the current content and dot position of the buffer.
	Buffer() (string, int)
	// SetBuffer sets the current content and dot position of the buffer.
	SetBuffer(buffer string, dot int)
	// InsertAtDot inserts text at the dot and moves the dot after it.
	InsertAtDot(text string)

	// SetMode sets the current mode of the Editor.
	SetMode(m Mode)
	// SetModeInsert sets the current mode of the Editor to insert mode.
	SetModeInsert()
	// SetModeListing sets the current mode of the Editor to listing mode with
	// the supplied binding and provider.
	SetModeListing(b BindingMap, p ListingProvider)

	// AddTip adds a message to the tip area.
	AddTip(format string, args ...interface{})
	// Notify writes out a message in a way that does not interrupt the editor
	// display. When the editor is not active, it simply writes the message to
	// the terminal. When the editor is active, it appends the message to the
	// notification queue, which will be written out during the update cycle. It
	// can be safely used concurrently.
	Notify(format string, args ...interface{})
}
