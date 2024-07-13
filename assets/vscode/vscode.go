package vscode

import (
	`embed`
	_ "embed"
)

//go:embed index.html
var HTML string

//go:embed static/*
var Static embed.FS

//go:embed monaco-editor/*
var Monaco embed.FS
