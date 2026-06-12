package highlight

// Syntax highlighting for terminal output using the chroma library
import (
	"io"
	"os"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

// Print writes syntax-highlighted content to stdout using ANSI terminal colors.
// Falls back to plain text if the language is unknown or highlighting fails.
func Print(content string, language string) error {
	return Fprint(os.Stdout, content, language)
}

// Fprint writes syntax-highlighted content to the given writer.
func Fprint(w io.Writer, content string, language string) error {
	// Pick the right lexer for the language
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	// Use monokai — looks great on dark terminals
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}

	// Use 256-color ANSI formatter for broad terminal support
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	// Tokenize and format
	iterator, err := lexer.Tokenise(nil, content)
	if err != nil {
		// Fall back to plain output on tokenize failure
		_, writeErr := io.WriteString(w, content)
		return writeErr
	}

	return formatter.Format(w, style, iterator)
}
