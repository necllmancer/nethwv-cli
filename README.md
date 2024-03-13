# nethwv-cli

![nethwv-cli logo](https://github.com/necllmancer/nethwv-cli/assets/96694331/6f51a6d9-357e-467d-af08-0ab402833983)

## Overview
Nether Weave CLI (`nethwv`) is a versatile command-line interface tool designed for aggregating files from GitHub repositories or local directories into a consolidated PDF document. It supports fetching files from specific branches, tags, or directories within GitHub repositories, as well as compiling files from a local directory. This tool is ideal for creating documentation, archives, or project reviews in an accessible PDF format.

## Features
- Clone files from GitHub repositories, specifying branches, tags, or directories.
- Compile files from local directories into a PDF.
- Streamline documentation and review processes by consolidating multiple files into a single document.

## Installation
To install `nethwv`, ensure you have Go installed on your machine and follow these steps:

1. Clone the repository to your local machine.
2. Navigate to the root directory of the project.
3. Run `make install` to install the CLI tool.

## Usage

### Git Subcommand

The `git` subcommand allows you to generate a PDF from files in a GitHub repository.

```
Usage: nethwv git [options] <user/repo> <output.pdf>

Options:
  -b string
        Branch to clone
  -d string
        Specific directory to retrieve files from
  -t string
        Tag to clone
```

Example usage:
```
nethwv git -b main -d docs octocat/Hello-World output.pdf
```

This command will clone the `main` branch of the `octocat/Hello-World` repository, retrieve files from the `docs` directory, and generate a PDF named `output.pdf`.

### Local Subcommand

The `local` subcommand allows you to generate a PDF from files in a local directory.

```
Usage: nethwv local -p <path> <output.pdf>

Options:
  -p string
        Path to local directory
```

Example usage:
```
nethwv local -p /path/to/local/directory output.pdf
```

This command will retrieve files from the specified local directory (`/path/to/local/directory`) and generate a PDF named `output.pdf`.

## Examples

Here are a few more examples of using nethwv-cli:

1. Generate a PDF from a specific tag of a GitHub repository:
   ```
   nethwv git -t v1.0.0 octocat/Spoon-Knife spoon-knife.pdf
   ```

2. Generate a PDF from a specific directory in a local path:
   ```
   nethwv local -p /home/user/documents/project project.pdf
   ```

3. Generate a PDF from the default branch of a GitHub repository:
   ```
   nethwv git octocat/Spoon-Knife spoon-knife.pdf
   ```

## License

This project is licensed under the [MIT License](LICENSE).