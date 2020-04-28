# book-list

Command line tool for CRUDing book JSON data.

## Usage

Clone the repo.
```bash
git clone https://github.com/cameron-yee/book-list.git
```

Set up GitHub user.
```bash
touch .env
```

Add `GITHUB_USER` and `GITHUB_PASSWORD` environment variables to `.env`.

Delete and create new files called `books.json`, `users.json`, and `reading-lists.json` to start with new data.
The files need an empty array to be valid JSON.
```bash
rm books.json && echo "[]" >>  books.json
rm users.json && echo "[]" >> users.json
rm reading-lists.json && echo "[]" >> reading-lists.json
```

Build the package.
```bash
go build
```

Add a symlink for the executable.
```bash
sudo ln -s <PATH-TO-CLONE>/book-list /usr/local/bin/booklist
```

Run commands.
```bash
booklist help
```

