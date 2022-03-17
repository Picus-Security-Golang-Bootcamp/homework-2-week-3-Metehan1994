package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Writer struct {
	writerID   int
	writerName string
}
type Book struct {
	ID                int
	bookName          string
	numOfPages        int
	numOfBooksinStock int
	price             int
	stockCode         string
	ISBN              string
	Writer
}
type BookList struct {
	AllBooks []Book
}

func main() {

	booksDetails := map[int][]interface{}{
		1:  {"Crime and Punishment", 705, 30, 30, "BOOK1111", "9780393055344", 1, "Fyodor Dostoyevski"},
		2:  {"War and Peace", 1225, 10, 50, "BOOK1112", "9780060798871", 2, "Lev Tolstoy"},
		3:  {"Anna Karenina", 864, 3, 40, "BOOK1113", "9780670894789", 2, "Lev Tolstoy"},
		4:  {"A Tale of Two Cities", 500, 15, 25, "BOOK1114", "9780721407104", 3, "Charles Dickens"},
		5:  {"Meditations", 304, 40, 20, "BOOK1115", "9780915145782", 4, "Marcus Aurelius"},
		6:  {"Robinson Crusoe", 210, 50, 18, "BOOK1116", "9780393044072", 5, "Daniel Defoe"},
		7:  {"The Gambler", 191, 13, 17, "BOOK1117", "9788426120113", 1, "Fyodor Dostoyevski"},
		8:  {"Fathers and Sons", 320, 34, 21, "BOOK1118", "9788423919307", 6, "Ivan Turgenyev"},
		9:  {"Don Quixote", 289, 12, 20, "BOOK1119", "9788497403573", 7, "Miguel de Cervantes"},
		10: {"Romeo and Juliet", 133, 15, 15, "BOOK1120", "9780003252453", 8, "William Shakespeare"},
	}

	ListofBooks := BookList{}
	for ID, OneBookDetails := range booksDetails {
		newBook := Book{}
		newBook = newBook.setAllProperties(ID, OneBookDetails)
		ListofBooks.AllBooks = ListofBooks.AddBook(newBook)
	}

	// wordPtr := flag.String("command", " ", "Listing all books")
	// flag.Parse()

	// if *wordPtr == "list" {
	// 	for _, book := range ListofBooks.AllBooks {
	// 		fmt.Println(book.bookName)
	// 	}
	// }

	if len(os.Args) == 1 { //condition for running the program without any arguments
		fmt.Println("You did not write any arguments to check the list of books.")
	} else if os.Args[1] == "list" && len(os.Args) == 2 { //condition for list as an only argument
		//Additional argument together with "list" will be not accepted
		list(ListofBooks.AllBooks)
	} else if len(os.Args) > 2 && os.Args[1] == "search" { //condition includes search and additional arguments accounted for book name.
		//Only search argument will be not accepted for searching a book.
		nameOfBook := ""
		for i := 2; i < len(os.Args); i++ {
			nameOfBook += os.Args[i]
			if i != len(os.Args)-1 { //condition to eliminate adding space after final word
				nameOfBook += " "
			}
		}
		search(nameOfBook, ListofBooks.AllBooks)
	} else if len(os.Args) == 3 && os.Args[1] == "get" {
		bookID, _ := strconv.Atoi(os.Args[2])
		get(bookID, ListofBooks.AllBooks)
	} else {
		fmt.Println("You have written an invalid argument.")
	}
}

func (b *Book) setAllProperties(ID int, Details []interface{}) Book {
	b.ID = ID
	b.bookName = Details[0].(string)
	b.numOfPages = Details[1].(int)
	b.numOfBooksinStock = Details[2].(int)
	b.price = Details[3].(int)
	b.stockCode = Details[4].(string)
	b.ISBN = Details[5].(string)
	b.Writer.writerID = Details[6].(int)
	b.Writer.writerName = Details[7].(string)
	return *b
}

func (b *BookList) AddBook(newBook Book) []Book {
	b.AllBooks = append(b.AllBooks, newBook)
	return b.AllBooks
}

//list shows the whole booklist.
func list(List []Book) {
	for i := 0; i < len(List); i++ {
		fmt.Println(List[i].bookName)
	}
}

//search takes a specific book name and the booklist, and searches whether the book which is asked is available in the booklist or not.
func search(nameOfBook string, List []Book) {
	bookFound := false
	nameOfBook = strings.ToLower(nameOfBook)
	for i := 0; i < len(List); i++ {
		book := List[i].bookName
		lowerCaseBook := strings.ToLower(book)
		if strings.Contains(lowerCaseBook, nameOfBook) {
			bookFound = true
			fmt.Println(List[i].bookName)
		}
	}
	if !bookFound {
		fmt.Println("There is no such a book in the booklist.")
	}
}

func get(bookID int, List []Book) {

	for i := 0; i < len(List); i++ {
		if bookID == List[i].ID {
			fmt.Println(List[i].bookName)
		}
	}
}
