# book-list

Command line tool for CRUDing book JSON data.

## Usage

Clone the repo.
```bash
git clone https://github.com/cameron-yee/book-list.git
```

Delete and create a new file called `books.json` to start with new data.
```bash
rm books.json && touch books.json
```

Write the initial structure to `books.json`.
```
{
  "books": []
}
```

Build the package.
```bash
go build
```

Alias the executable in ~/.zshrc.
```bash
alias book_list=<PATH-TO-CLONE>/book-list
```

Run commands.
```bash
book_list help
```

