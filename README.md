
# gogitlocalstats

**gogitlocalstats** is a command-line tool designed to streamline tracking Git contributions across multiple repositories on your local machine. With a simple interface, it scans specified directories for Git repositories, compiles commit statistics, and displays contribution history in a visually organized grid.

## Key Features

- **Automatic Repository Detection**: Scans directories to identify all Git repositories, storing paths in a local file for future reference.
- **Detailed Contribution Tracking**: Analyzes commit history to display daily contributions for the last six months, highlighting active days.
- **Customizable Filtering**: Filters commits by email address to tailor contribution stats to your specific user profile.
- **Easy Access to Historical Stats**: Outputs a clear visual representation of commits per day over six months, allowing users to track activity at a glance.

## Installation

1. Clone the repository.
2. Compile the Go files and move the executable to a directory in your PATH, like `/usr/local/bin`.

```bash
go build -o gogitlocalstats main.go scan.go stats.go
sudo mv gogitlocalstats /usr/local/bin/
```

## Usage

- Add a new directory to scan for Git repositories:

  ```bash
  gogitlocalstats -add /path/to/directory
  ```

- View contribution statistics for a specific email:

  ```bash
  gogitlocalstats -email your-email@example.com
  ```

Perfect for developers with multiple local repositories, **gogitlocalstats** helps you stay on top of your Git contributions across all your projects, all in one place.
