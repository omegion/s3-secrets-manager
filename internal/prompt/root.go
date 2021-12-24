package prompt

// Prompt is a struct for prompt control.
type Prompt struct{}

// Options is options for Prompt.
type Options struct{}

// NewPrompt inits new Prompt.
func NewPrompt(options Options) *Prompt {
	return &Prompt{}
}
