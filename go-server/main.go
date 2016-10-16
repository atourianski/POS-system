package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	//"reflect"
	"strconv"
	"strings"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

	http.ListenAndServe(":8080", nil)
}

var defaultHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>`

type ActiveLockers struct {
	Bracelet_num string
	Entry_time   string
}

type FoodAvailable struct {
	Name  string
	Price string // convert to int...or marshal as such?
}

type DrinksAvailable struct {
	Name  string
	Price string // convert to int...or marshal as such?
}

var openSessionsHTML = `<h2>Open Sessions</h2>
Entry Time | Locker Number<br>
{{range $index, $element := .}}
<b>{{ .Entry_time}}</b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{ .Bracelet_num}}<br>
{{end}}`

// TODO use some fucking templates
func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func newSession(w http.ResponseWriter, r *http.Request) {
	var tpl = defaultHTML + `<body>
	<h1>New Session</h1>
	Enter an available locker number:
	<form method="POST" action="/initializeSession">
	<input type="number" name="bracelet_id">
	<input type="submit" value="Submit"><br>
	<br>
	<a href="/">Home</a>` + openSessionsHTML + `
</body>
</html>`

	t := template.New("t")
	t, err = t.Parse(tpl)
	if err != nil {
		panic(err)
	}

	Lockers := getActiveSession()

	var doc bytes.Buffer
	if err := t.Execute(&doc, Lockers); err != nil {
		panic(err)
	}
	w.Write(doc.Bytes())
}

func addItems(w http.ResponseWriter, r *http.Request) {

	foodTemplate := `<b>Food:</b><br>
		{{range $index, $element := .}}
		{{ .Name}} $ {{ .Price}}<br>
		<input type="number" name="{{ .Name}}"><br>
		{{end}}
		<br>
		`

	tFood := template.New("tFood")
	tFood, err = tFood.Parse(foodTemplate)
	if err != nil {
		panic(err)
	}

	Foods := getActiveFoodstuffs()

	var doc bytes.Buffer
	if err := tFood.Execute(&doc, Foods); err != nil {
		panic(err)
	}
	foodHTML := doc.String()

	// ------------ TODO deduplicate w/ above ------

	var drinksTemplate = `<b>Drinks:</b><br>
		{{range $index, $element := .}}
		{{ .Name}} $ {{ .Price}}<br>
		<input type="number" name="{{ .Name}}"><br>
		{{end}}
		<br>
		`

	tDrinks := template.New("tDrinks")
	tDrinks, err = tDrinks.Parse(drinksTemplate)
	if err != nil {
		panic(err)
	}

	Drinks := getActiveDrinks()

	var doc1 bytes.Buffer
	if err := tDrinks.Execute(&doc1, Drinks); err != nil {
		panic(err)
	}
	drinksHTML := doc1.String()

	// the rendered display
	var htmlToWrite = defaultHTML + `<body>
	<h1>Add Some Items</h1>
	<form method="POST" action="/addItemsToASession">
		Locker Number:<br>
		<input type="number" name="bracelet_id"><br>
		<br>
		` + foodHTML + `<br>
		<br>
		` + drinksHTML + `<br>

		Misc: <b>must set dollar amount</b><br>
		<input type="number" name="misc"><br>

	<input type="submit" value="Submit">
	</form><br>
	<a href="/">Home</a>
</body>
</html>`
	w.Write([]byte(htmlToWrite))
}

func closeSession(w http.ResponseWriter, r *http.Request) {
	var tpl = defaultHTML + `<body>
		<h1>Close A Session</h1>
		<form method="POST" action="/displayBill">
		<input type="number" name="bracelet_id">
		<input type="submit" value="Submit">
		</form><br>
		<a href="/">Home</a><br>` + openSessionsHTML + `
	</body>
	</html>`

	Lockers := getActiveSession()

	t := template.New("t")
	t, err = t.Parse(tpl)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	if err := t.Execute(&doc, Lockers); err != nil {
		panic(err)
	}
	w.Write(doc.Bytes())
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

		orders := map[string]int{}

		// MISC
		miscellaneous, _ := strconv.Atoi(r.FormValue("misc"))
		orders["misc"] = miscellaneous

		// FOOD
		Foods := getActiveFoodstuffs()
		totalsFood := 0
		// iterate through map, match "set" items from user input w/ "active" items from db
		for _, activeFoods := range Foods {
			numberOrdered := r.FormValue(activeFoods.Name)
			if numberOrdered != "" && numberOrdered != "0" {
				num0, _ := strconv.Atoi(numberOrdered)
				price, _ := strconv.Atoi(activeFoods.Price)
				totalsFood = totalsFood + (num0 * price)
			}
		}
		orders["food"] = totalsFood

		// DRINK
		Drinks := getActiveDrinks()
		totalsDrink := 0
		// iterate through map, match "set" items from user input w/ "active" items from db
		for _, activeDrinks := range Drinks {
			numberOrdered := r.FormValue(activeDrinks.Name)
			if numberOrdered != "" && numberOrdered != "0" {
				num0, _ := strconv.Atoi(numberOrdered)
				price, _ := strconv.Atoi(activeDrinks.Price)
				totalsDrink = totalsDrink + (num0 * price)
			}
		}
		orders["drink"] = totalsDrink

		// get current invoice
		var curFood string
		var curDrink string
		var curMisc string

		// here we read in some totals & add to them with a new purchase
		err := db.QueryRow("SELECT food, drink, misc FROM invoices WHERE bracelet_id=? AND invoice_id=?", braceletID, invoiceID).Scan(&curFood, &curDrink, &curMisc)
		if err != nil {
			panic(err)
		}
		// convert to int
		curF, _ := strconv.Atoi(curFood)
		curD, _ := strconv.Atoi(curDrink)
		curM, _ := strconv.Atoi(curMisc)

		// add to what was just ordered
		food := orders["food"] + curF
		drink := orders["drink"] + curD
		misc := orders["misc"] + curM

		// update invoice with food, drink, misc
		_, err = db.Exec("UPDATE invoices SET food=?,drink=?,misc=? WHERE invoice_id=?", food, drink, misc, invoiceID)
		if err != nil {
			http.Error(w, fmt.Sprintf("error updating invoice: %v", err), http.StatusInternalServerError)
			return
		}

		time.Sleep(time.Second * 1)
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
		// XXX this is hacky & can probably be cleaned up
		// query invoice ID to get bill info & display
		var banyaCost string
		var foodCost string
		var drinkCost string
		var miscCost string
		err = db.QueryRow("SELECT banya, food, drink, misc FROM invoices WHERE bracelet_id=? AND invoice_id=?", braceletID, invoiceID).Scan(&banyaCost, &foodCost, &drinkCost, &miscCost)
		if err != nil {
			panic(err)
		}

		bCost, _ := strconv.Atoi(banyaCost)
		fCost, _ := strconv.Atoi(foodCost)
		dCost, _ := strconv.Atoi(drinkCost)
		mCost, _ := strconv.Atoi(miscCost)

		totalCost := bCost + fCost + dCost + mCost

		someHTML := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	Bracelet #: ` + braceletID + `<br>
	Invoice #: ` + invoiceID + `<br>
	Entry Time: ` + entryTime + `<br>

	<br>
	<b>Bill:</b><br>
	<br>

	Banya: ` + banyaCost + `<br>
	Food: ` + foodCost + `<br>
	Drink: ` + drinkCost + `<br>
	Misc: ` + miscCost + `<br>

	<br>
	<b>Total: $ ` + strconv.Itoa(totalCost) + `</b><br>
	<br>
	
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
		// hacky
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

		time.Sleep(time.Second * 1)
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

func getActiveSession() []*ActiveLockers {
	rows, err := db.Query("SELECT bracelet_num, entry_time FROM visit WHERE active=1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Lockers []*ActiveLockers

	for rows.Next() {
		l := new(ActiveLockers)
		err = rows.Scan(&l.Bracelet_num, &l.Entry_time)
		if err != nil {
			panic(err)
		}
		Lockers = append(Lockers, l)
	}
	return Lockers
}

func getActiveFoodstuffs() []*FoodAvailable {
	rows, err := db.Query("SELECT name, price FROM foodstuffs WHERE active=1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Foods []*FoodAvailable

	for rows.Next() {
		f := new(FoodAvailable)
		err = rows.Scan(&f.Name, &f.Price)
		if err != nil {
			panic(err)
		}
		Foods = append(Foods, f)
	}
	return Foods
}

func getActiveDrinks() []*DrinksAvailable {
	rows, err := db.Query("SELECT name, price FROM drinks WHERE active=1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var Drinks []*DrinksAvailable

	for rows.Next() {
		d := new(DrinksAvailable)
		err = rows.Scan(&d.Name, &d.Price)
		if err != nil {
			panic(err)
		}
		Drinks = append(Drinks, d)
	}
	return Drinks
}

func getFormattedDate() string {
	now := time.Now()
	return fmt.Sprintf("%s-%s-%s", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()))
}
