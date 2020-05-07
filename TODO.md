# TODO
  -- **CLI**: previous field editing (":back")
  -- add flags to allow fields to be added from command line without input
    * booklist add book -t "test" -a "C.S. Lewis" -g Fantasy -rb "Cameron Yee, Jordan Yee" -o false -eo "Cameron Yee"
  -- add flag to sort alphabetically
  -- add print range

## WIP
  -- **CLI**: print defaults (owned (true/false):)
  -- **CLI**: more options for boolean fields (true/false/t/f/yes/no/y/n)

## DONE
  -- add flag for true/false (booklist search readby "Cameron Yee" -f --false
  -- add default messages for no search results found
  -- add deleteReadBy update function
  -- change read field to readBy field ("Cameron Yee, Jordan Yee")
  -- limit search results (flag -l 10 --limit 10)
  -- GitHub pull check to make sure merge conflicts don't ruin everything
  -- clean up printing
  -- verbose printing for search/filter
  -- clean up CLI so it makes since
  -- change panic calls to fmt.Print
  -- Add CLI functions for users and reading-lists
  -- update reading list members
  -- update user reading lists
  -- update reading list books

## NOPE
  -- **CLI**: booklist list series "Wheel of Time"
    * Can be done with search **CLI**
-----------------------------------------------------------------

## Feature requests
  -- change genre to a list of predefined fields (genre.json)
  -- JSON file imports
  -- *__books.json append to books.json
  -- more advanced search/filter (composite search/filter)
  -- more verbose options
  -- list book fields (book_list list authors) 
    * number of books for author 
    * allow filter (author="Bonhoeffer", read=false)

## New fields
- Book length (short, medium, long)
- Series order
- Tags
- Note

