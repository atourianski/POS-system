package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"

	"fmt"
	"net/http"
	"time"
)

var db *sql.DB
var err error

func main() {
	// main shit
	db, err = sql.Open("mysql", "root:secure@tcp(172.17.0.2:3306)/banya")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// new user arrives. we call it a session & send in the bracelet ID
	http.HandleFunc("/initializeSession", initializeSession)

	// if adding new items, list them (keep track internally of the ID)
	// TODO
	// http.HandleFunc("/listAvailableItems", listAvailableItems)

	// select some items from above and add them
	http.HandleFunc("/addItemsToASession", addItemsToASession)

	http.HandleFunc("/displayBill", displayBill)
	http.HandleFunc("/closeBill", closeBill)

	// file serving endpoints
	http.HandleFunc("/", homePage)
	http.HandleFunc("/newSession", newSession)
	http.HandleFunc("/addItems", addItems)
	http.HandleFunc("/closeSession", closeSession)

	// serve it all up
	http.ListenAndServe(":8080", nil)
}

var defaultHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
`

// TODO use some fucking templates
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func newSession(w http.ResponseWriter, r *http.Request) {
	htmlToWrite := defaultHTML + `<body>
	<h1>New Session</h1>
	Enter an available locker number:
	<form method="POST" action="/initializeSession">
	<input type="number" name="bracelet_id">
	<input type="submit" value="Submit">
</body>
</html>`
	w.Write([]byte(htmlToWrite))
	//http.ServeFile(w, r, "new_session.html")
}

func addItems(w http.ResponseWriter, r *http.Request) {
	htmlToWrite := defaultHTML + `<body>
	<h1>Add Some Items</h1>
	<form method="POST" action="/addItemsToASession">
		Locker Number:<br>
		<input type="number" name="bracelet_id"><br>

		Soup:<br>
		<input type="number" name="soup"><br>

		Beer:<br>
		<input type="number" name="beer"><br>

		Scrub:<br>
		<input type="number" name="scrub"><br>

	<input type="submit" value="Submit">
	</form><br>
	<a href="/">Home</a>
</body>
</html>`

	w.Write([]byte(htmlToWrite))
	//http.ServeFile(w, r, "add_items.html")
}

func closeSession(w http.ResponseWriter, r *http.Request) {
	htmlToWrite := defaultHTML + `<body>
	<h1>Close A Session</h1>
	<form method="POST" action="/displayBill">
	<input type="number" name="bracelet_id">
	<input type="submit" value="Submit">
	</form><br>
	<a href="/">Home</a>
</body>
</html>`
	w.Write([]byte(htmlToWrite))
	//http.ServeFile(w, r, "close_session.html")
}

func initializeSession(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		braceletID := r.FormValue("bracelet_id")

		activeSession, _ := doesSessionAlreadyExist(braceletID) // throw out invoiceID, it shouldn't exist yet...
		if activeSession {
			http.Error(w, "session already exists. plz close existing # before continuing", http.StatusInternalServerError)
			return // re-direct ... ?
		}

		// to be modified to allow scrubs/misc to be purchased on init
		_, err = db.Exec("INSERT INTO invoices(bracelet_id, banya, food, drink, misc) values (?, ?, 0, 0, 0)", braceletID, "35")
		if err != nil {
			http.Error(w, fmt.Sprintf("error creating new invoice: %v", err), http.StatusInternalServerError)
		}

		var invoiceID string

		err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&invoiceID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting last_insert_id: %v", err), http.StatusInternalServerError)
		}

		date := getFormattedDate()
		// active => 1 => true
		_, err = db.Exec("INSERT INTO visit(date, bracelet_num, entry_time, invoice_id, active) values (?, ?, ?, ?, ?)", date, braceletID, time.Now(), invoiceID, "1")
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting new session: %v", err), http.StatusInternalServerError)
		}

		time.Sleep(time.Second * 1)
		http.Redirect(w, r, "/", 301)
	}
}

// once the iterms are selected, they need to be written to the cumulative invoice
// invoice being the "thing" that is derived from a session at the end of it, basically
func addItemsToASession(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		braceletID := r.FormValue("bracelet_id")

		activeSession, invoiceID := doesSessionAlreadyExist(braceletID)
		if !activeSession {
			http.Error(w, "locker does not exist, plz use an active bracelet #", http.StatusInternalServerError)
			return // re-direct ... ?
		}

		foodPrice := 8
		drinkPrice := 5
		miscPrice := 3

		f, _ := strconv.Atoi(r.FormValue("soup"))
		d, _ := strconv.Atoi(r.FormValue("beer"))
		m, _ := strconv.Atoi(r.FormValue("scrub"))

		food := f * foodPrice
		drink := d * drinkPrice
		misc := m * miscPrice

		// XXX need to read from this row first & add to total, rather than overwrite!!!!
		_, err = db.Exec("UPDATE invoices SET food=?,drink=?,misc=? WHERE invoice_id=?", food, drink, misc, invoiceID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error updating invoice: %v", err), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/addItems", 301)
	}
}

func displayBill(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		braceletID := r.FormValue("bracelet_id")

		activeSession, invoiceID := doesSessionAlreadyExist(braceletID)
		if !activeSession {
			http.Error(w, "cannot close a session that is not active", http.StatusInternalServerError)
			return // re-direct ... ?
		}

		// query active session for entry time info - everything else is basically accessible
		var entryTime string
		// 1 is true is active
		err := db.QueryRow("SELECT entry_time FROM visit WHERE bracelet_num=? AND active=1", braceletID).Scan(&entryTime)
		if err != nil {
			panic(err)
		}

		// query invoice ID to get bill info & display
		var banyaCost string
		var foodCost string
		var drinkCost string
		var miscCost string
		err = db.QueryRow("SELECT banya, food, drink, misc FROM invoices WHERE bracelet_id=? AND invoice_id=?", braceletID, invoiceID).Scan(&banyaCost, &foodCost, &drinkCost, &miscCost)
		if err != nil {
			panic(err)
		}

		someHTML := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	<h1>Bracelet #: ` + braceletID + `</h1>
	<h1>Invoice #: ` + invoiceID + `</h1>
	<h2>Entry Time: ` + entryTime + `</h2>
	<h1>Bill:</h1>
	<h2>banya: ` + banyaCost + `</h2>
	<h2>food: ` + foodCost + `</h2>
	<h2>drink: ` + drinkCost + `</h2>
	<h2>misc: ` + miscCost + `</h2>

	<form method="POST" action="/closeBill?` + braceletID + `&` + invoiceID + `">
	<input type="submit" value="Close Bill">
	</form><br>

	<a href="/">Home</a>
</body>
</html>`
		w.Write([]byte(someHTML))
	}
}

func closeBill(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		theURL := r.URL.String()
		getFields := strings.Split(theURL, "?")[1]
		braceletID := strings.Split(getFields, "&")[0]
		invoiceID := strings.Split(getFields, "&")[1]

		// 0 is bool for false
		_, err = db.Exec("UPDATE visit SET exit_time=?,active=0 WHERE bracelet_num=? AND invoice_id=?", time.Now(), braceletID, invoiceID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error closing invoice: %v", err), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", 301)
	}
}

// ------- helpers -----------------

// braceletID should be int and not string!
func doesSessionAlreadyExist(braceletID string) (bool, string) {
	var invoiceID string
	// this SELECT has duplication throughout
	err := db.QueryRow("SELECT invoice_id FROM visit WHERE bracelet_num=? AND active=1", braceletID).Scan(&invoiceID)

	if err == sql.ErrNoRows {
		// bracelet ID does not currently exist in open visits
		return false, ""
	} else {
		// bracelet ID does exist - do not create new session
		return true, invoiceID
	}
}

func getFormattedDate() string {
	now := time.Now()
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()))
}
