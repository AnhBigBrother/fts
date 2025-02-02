# fts

A CLI application that simulates folder-tree structure.

## Installation

**Prerequisites:**

- Go installed
- A command line interface

**Install:**

```bash
go install github.com/AnhBigBrother/fts
```

## Usage

`fts [command]`

**Available Commands:**

`add`     Add a new file or folder  
`del`     Remove file or folder  
`get`     Retrieve file/folder  
`search`  Search file/folder by name  
`show`    Show folder-tree structure  
`help`    Help about any command  

**Flags:**

`-f`, `--file`   <bool> set is file or not  
`-d`, `--dir`    <string> file/folder's parent directory (default "/")  
`-n`, `--name`   <string> file/folder name (default "NAME")  
`-t`, `--type`   <string> file type (default "txt")  
`-h`, `--help`            help for fts  

**Note:** Use `fts [command] --help` for more information about a command.

**Example:**

```bash
$ fts add -n folder1
added folder folder1 to directory /

$ fts add -d /folder1 -n folder2
added folder folder2 to directory /folder1

$ fts add -f -d folder1 -n file1 -t txt Lorem ipsum sit amet...
added file file1.txt to directory folder1

$ fts add -f -d /folder1/folder2 -n file2 -t txt Lorem ipsum sit amet...
added file file2.txt to directory /folder1/folder2

$ fts get -f -d /folder1/folder2 -n file2 -t txt
{
  "directory": "/folder1/folder2",
  "file_name": "file2",
  "content": "Lorem ipsum sit amet...",
  "file_type": "txt"
}

$ fts get -d /folder1
{
  "directory": "/",
  "folder_name": "folder1",
  "files": [
    "file1.txt"
  ],
  "sub_folders": [
    "folder2"
  ]
}

$ fts show
|__ /
      |__ folder1
          |__ folder2

$ fts del -f -d /folder1/folder2 -n file2 -t txt
Deleted file file2.txt from /folder1/folder2 directory

$ fts del -d /folder1 -n folder2
Deleted folder folder2from /folder1 directory

$ fts search -f file1
found 1 files named 'file1':
{
  "directory": "/folder1",
  "file_name": "file1",
  "content": "Lorem ipsum sit amet...",
  "file_type": "txt"
}

$ fts search folder1
found 1 folders named 'folder1':
{
  "directory": "/",
  "folder_name": "folder1",
  "files": [
    "file1.txt"
  ],
  "sub_folders": []
}
```
