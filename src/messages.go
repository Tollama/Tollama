package main

import (
	"github.com/ollama/ollama/api"
)

type errMsg struct{ err error }
type modelsMsg struct{ models []api.ListModelResponse }
