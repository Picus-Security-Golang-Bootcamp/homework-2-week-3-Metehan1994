package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID                int    `json:"ID"`
	BookName          string `json:"bookName"`
	NumOfPages        int    `json:"numOfPages"`
	NumOfBooksinStock int    `json:"numOfBooksinStock"`
	Price             int    `json:"price"`
	StockCode         string `json:"stockCode"`
	ISBN              string `json:"ISBN"`
	Writer            struct {
		WriterID   int    `json:"WriterID"`
		WriterName string `json:"WriterName"`
	} `json:"Writer"`
	IsDeleted bool `json:"isDeleted"`
}

type BookList struct {
	AllBooks []Book
}

var usage = `Usage: 5 command line arguments are valid, which are "list", "search", "get", "delete" and "buy".
Options:
	-"list" is used without any arguments.
	-"search" is used with a string to check if the book is available or not.
	-"get" is used with an argument specifying its ID
	-"delete" is used with an argument specifying its ID
	-"buy" is used with two arguments specifying its ID and number of books ordered
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}

	booksDetails := ReadData()

	ListofBooks := BookList{}
	for _, OneBookDetails := range booksDetails {
		newBook := Book{}
		newBook = newBook.setAllProperties(OneBookDetails)
		ListofBooks.AllBooks = ListofBooks.AddBook(newBook)
	}

	if len(os.Args) == 1 { //condition for running the program without any arguments
		fmt.Println("You did not write any arguments to check the list of books.")
	} else if os.Args[1] == "list" && len(os.Args) == 2 { //condition for list as an only argument
		list(ListofBooks.AllBooks)
	} else if len(os.Args) > 2 && os.Args[1] == "search" { //condition includes search and additional arguments accounted for book name.
		nameOfBook := ""
		for i := 2; i < len(os.Args); i++ {
			nameOfBook += os.Args[i]
			if i != len(os.Args)-1 { //condition to eliminate adding space after final word
				nameOfBook += " "
			}
		}
		search(nameOfBook, ListofBooks.AllBooks)
	} else if len(os.Args) == 3 && os.Args[1] == "get" { //condition allowing only for "get" with ID
		bookID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("It cannot be converted into a int")
		}
		get(bookID, ListofBooks.AllBooks)
	} else if len(os.Args) == 3 && os.Args[1] == "delete" { //condition allowing only for "delete" with ID
		bookID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("It cannot be converted into a int")
		}
		ListofBooks.AllBooks = delete(bookID, ListofBooks.AllBooks)
	} else if len(os.Args) == 4 && os.Args[1] == "buy" { //condition allowing only for "buy" with ID and number of order
		bookID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("String for book ID cannot be converted into a int")
		}
		numOfBooksOrdered, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("String for order number cannot be converted into a int")
		}
		ListofBooks.AllBooks = buy(bookID, numOfBooksOrdered, ListofBooks.AllBooks)
	} else {
		usageAndExit("It is not used, properly.")
	}

	WriteData(ListofBooks.AllBooks)

}

//ReadData converts the json formatted data into the map
func ReadData() []map[string]interface{} {
	booksDetails := []map[string]interface{}{}
	contents, err := ioutil.ReadFile("Books.json")
	if err != nil {
		fmt.Println("The data cannot be read from the json file")
	}

	if err := json.Unmarshal(contents, &booksDetails); err != nil {
		fmt.Println("The data cannot be read from the json file")
	}
	return booksDetails
}

//WriteData converts the struct into json formatted data after changes
func WriteData(List []Book) {
	file, _ := json.MarshalIndent(List, "", " ")

	_ = ioutil.WriteFile("Books.json", file, 0644)
}

//setAllProperties initializes structs with their attributes
func (b *Book) setAllProperties(Details map[string]interface{}) Book {
	b.ID = int(Details["ID"].(float64))
	b.BookName = Details["bookName"].(string)
	b.NumOfPages = int(Details["numOfPages"].(float64))
	b.NumOfBooksinStock = int(Details["numOfBooksinStock"].(float64))
	b.Price = int(Details["price"].(float64))
	b.StockCode = Details["stockCode"].(string)
	b.ISBN = Details["ISBN"].(string)
	b.Writer.WriterID = int((Details["Writer"].(map[string]interface{}))["WriterID"].(float64))
	b.Writer.WriterName = (Details["Writer"].(map[string]interface{}))["WriterName"].(string)
	b.IsDeleted = Details["isDeleted"].(bool)
	return *b
}

//AddBook adds created book structs
func (b *BookList) AddBook(newBook Book) []Book {
	b.AllBooks = append(b.AllBooks, newBook)
	return b.AllBooks
}

//list shows the whole booklist.
func list(List []Book) {
	for i := 0; i < len(List); i++ {
		if !List[i].IsDeleted {
			fmt.Println(List[i].BookName)
		}
	}
}

//search determines whether the string written corresponds any books or not.
func search(nameOfBook string, List []Book) {
	bookFound := false
	nameOfBook = strings.ToLower(nameOfBook)
	for i := 0; i < len(List); i++ {
		book := List[i].BookName
		lowerCaseBook := strings.ToLower(book)
		if strings.Contains(lowerCaseBook, nameOfBook) && !List[i].IsDeleted {
			bookFound = true
			fmt.Println(List[i].BookName)
		}
	}
	if !bookFound {
		fmt.Println("There is no such a book in the booklist.")
	}
}

//PrintBookInfo prints all info defined in struct
func PrintBookInfo(b Book) {
	fmt.Println("Book ID:", b.ID)
	fmt.Println("Name:", b.BookName)
	fmt.Println("Pages:", b.NumOfPages)
	fmt.Println("Number of books in stock:", b.NumOfBooksinStock)
	fmt.Println("price:", b.Price, "TL")
	fmt.Println("Stock Code:", b.StockCode)
	fmt.Println("ISBN:", b.ISBN)
	fmt.Println("Stock Code:", b.StockCode)
	fmt.Println("Author ID:", b.Writer.WriterID)
	fmt.Println("Author Name:", b.Writer.WriterName)
}

//get prints the book info for demanded ID
func get(bookID int, List []Book) {

	for i := 0; i < len(List); i++ {
		if bookID == List[i].ID {
			if List[i].IsDeleted {
				fmt.Println(List[i].BookName)
			} else {
				PrintBookInfo(List[i])
			}
		}
	}
}

//delete states the book deleted
func delete(bookID int, List []Book) []Book {

	for i := 0; i < len(List); i++ {
		if bookID == List[i].ID {
			if List[i].IsDeleted {
				fmt.Println("It has been already deleted.")
				return List
			}
			fmt.Printf("%s is deleted.", List[i].BookName)
			List[i].BookName += " (It has been removed from the list)"
			List[i].NumOfBooksinStock = 0
			List[i].IsDeleted = true
		}
	}

	return List
}

//buy function reduces the number of books in a stock according to ID and order number
func buy(bookID int, numOfOrders int, List []Book) []Book {
	bookIDFound := false
	for i := 0; i < len(List); i++ {
		if bookID == List[i].ID {
			if List[i].IsDeleted {
				fmt.Println("The book has been deleted from the list.")
				return List
			}
			bookIDFound = true
			if numOfOrders < List[i].NumOfBooksinStock {
				List[i].NumOfBooksinStock = List[i].NumOfBooksinStock - numOfOrders
			} else {
				fmt.Println("There is no sufficient amount of book!!!")
			}
			PrintBookInfo(List[i])
		}
	}
	if !bookIDFound {
		fmt.Println("The book ID is wrong.")
	}
	return List
}

func usageAndExit(msg string) {

	fmt.Fprintf(os.Stderr, msg)
	fmt.Fprintf(os.Stderr, "\n\n")
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
