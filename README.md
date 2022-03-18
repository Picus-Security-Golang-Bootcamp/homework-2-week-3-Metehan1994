# Homework-2 | Week 3 | Booklist App

## Overview

The program interprets a number of books formatted in json file, list them, searches books among them, gets info for books, create an order for sales and remove a specific book from the list.

## How to Use the App ?

The program works with 5 command line arguments which are:

* "list"
* "search"
* "buy"
* "get"  
* "delete"

***

### list command

```
go run main.go list
```

"list" command shows the booklist.
***

### search command

```
go run main.go search <bookName>
```

"search" command with a string determines whether the string written corresponds any books or not.
***

### get command

```
go run main.go get <bookID>
```

"get" command prints the book info for demanded ID.
***

### delete command

```
go run main.go delete <bookID>
```

"delete" command with book ID removes the book from the list.
***

### buy command

```
go run main.go buy <bookID> <quantity>
```

"buy" command creates an order for the book according to ID in a given quantity.
***

### Some Notes for Usage

1. Program produces error messages when it is executed without considering its usage.

2. After deleting a book, it can be seen with "get" command.

3. The program is insensitive to uppercase & lowercase letters for booknames which makes easier to search the book.

## Package Used

* The program is created with **GO main package**.

* Its import modules like **"flag", "fmt", "strconv", "strings", "os", "io/ioutil"** and **"encoding/json"** are used.
