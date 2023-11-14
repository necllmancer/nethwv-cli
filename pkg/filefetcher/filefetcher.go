package filefetcher

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FetchFileContent(filePath string) (string, error) {
	if !isTextFile(filePath) {
		return "", errors.New("not unsupported file")
	}
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func isTextFile(url string) bool {
	textExtensions := []string{
		".txt", ".md", ".xml", ".json", ".yaml", ".yml", ".csv", ".tsv", ".ini", ".cfg", ".conf", ".properties", ".log", ".rtf", ".tex", ".html", ".htm", ".xhtml", ".css", ".scss", ".sass", ".less", ".js", ".mjs", ".jsx", ".ts", ".tsx", ".sql", ".graphql", ".gql", ".sh", ".bash", ".zsh", ".bat", ".cmd", ".ps1", ".vbs", ".rb", ".erb", ".py", ".pyw", ".r", ".pl", ".php", ".phtml", ".jsp", ".asp", ".aspx", ".cgi", ".lua", ".swift", ".go", ".java", ".class", ".jar", ".c", ".cpp", ".cc", ".cxx", ".h", ".hpp", ".hxx", ".cs", ".vb", ".f", ".f90", ".f95", ".rkt", ".clj", ".cljs", ".cljc", ".groovy", ".gvy", ".kt", ".kts", ".m", ".pas", ".d", ".scala", ".sbt", ".rs", ".elm", ".ex", ".exs", ".eex", ".hx", ".nim", ".b", ".bf", ".v", ".sv", ".svh", ".vhd", ".vhdl", ".ucf", ".qsf", ".asm", ".s", ".agc", ".ags", ".aea", ".sed", ".awk", ".ps", ".eps", ".jsonld", ".webmanifest", ".toml", ".bowerrc", ".npmrc", ".dockerignore", ".gitignore", ".editorconfig", ".prettierrc", ".sol", ".au3", ".nsh", ".nsl", ".lua", ".mak", ".sln", ".vcxproj", ".csproj", ".fsproj", ".xproj", ".dproj", ".lpr", ".gpr", ".uproject", ".pbxproj", ".tcl", ".exp", ".brs", ".wren",
		".c", ".cpp", ".h", ".hpp", ".cs", ".java", ".py", ".pyw", ".js", ".jsx", ".ts", ".tsx", ".rb", ".php", ".swift", ".go", ".rs", ".lua", ".r", ".sh", ".bash", ".zsh", ".bat", ".cmd", ".pl", ".scala", ".kt", ".groovy", ".dart", ".f", ".f90", ".f95", ".asm", ".s", ".vhdl", ".vhd", ".verilog", ".v", ".sv", ".tcl", ".awk", ".sed", ".ps", ".rkt", ".ex", ".exs", ".eex", ".erl", ".hrl", ".hs", ".lhs", ".ml", ".mli", ".sml", ".thy", ".vbs", ".cls", ".vb", ".bas", ".frm", ".ctl", ".vba", ".pas", ".dpr", ".dfm", ".lpr", ".pp", ".fs", ".fsx", ".fsi", ".fsscript", ".lsx", ".lsp", ".scm", ".ss", ".st", ".cob", ".cpy", ".cbl", ".asmx", ".aspx", ".ascx", ".asm", ".nasm", ".s", ".agc", ".ags", ".aea", ".lua", ".p", ".pas", ".d", ".m", ".mm", ".cshtml", ".vbhtml", ".jsp", ".jspx", ".asp", ".aspx", ".asax", ".ashx", ".asmx", ".ascx", ".axd", ".php3", ".php4", ".php5", ".p",
	}

	for _, ext := range textExtensions {
		if strings.HasSuffix(url, ext) {
			return true
		}
	}
	return false
}
