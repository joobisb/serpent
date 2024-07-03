package completion

import (
	"fmt"
	"io"
	"text/template"
)

const bashCompletionTemplate = `
_generate_{{.Name}}_completions() {
    # Capture the full command line as an array, excluding the first element (the command itself)
    local args=("${COMP_WORDS[@]:1}")

    # Set COMPLETION_MODE and call the command with the arguments, capturing the output
    local completions=$(COMPLETION_MODE=1 "{{.Name}}" "${args[@]}")

    # Use the command's output to generate completions for the current word
    COMPREPLY=($(compgen -W "$completions" -- "${COMP_WORDS[COMP_CWORD]}"))

    # Ensure no files are shown, even if there are no matches
    if [ ${#COMPREPLY[@]} -eq 0 ]; then
        COMPREPLY=()
    fi
}

# Setup Bash to use the function for completions for '{{.Name}}'
complete -F _generate_{{.Name}}_completions {{.Name}}
`

func GenerateBashCompletion(
	w io.Writer,
	rootCmdName string,
) error {
	tmpl, err := template.New("bash").Parse(bashCompletionTemplate)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	err = tmpl.Execute(
		w,
		map[string]string{
			"Name": rootCmdName,
		},
	)
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}
