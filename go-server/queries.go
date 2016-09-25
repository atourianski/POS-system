package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// probably want something more secure...
const (
	user     = "root"
	password = "default"
	database = "banya"
)

// for summary of connecting to mysql server, see:
// http://stackoverflow.com/questions/11353679/whats-the-recommended-way-to-connect-to-mysql-from-go

func CheckThatTheDayIsReady() error {
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil {
		return err
	}
	defer con.Close()

	// TODO some checks in here that all is well

	return nil
}

// braceletID should be int and not string!
func doesSessionExist(braceletID string) (bool, error) {
	// open/close & query current day
	return true, nil
}

func addNewSession(braceletID string) error { // if not error, success!
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil {
		return err
	}
	defer con.Close()

	// first, check that this braceletID is not already registered, then

	// insert a new row with a bunch of information
	_, err = con.Exec("insert into table (bracelet_id, time, somethingElse) values (?, ?, ?)", braceletID, time.Now(), "someString")

	return nil
}

type TheStructToFill struct{}

// will also need to return a struct populated with all the relevant info:
// time spent at banya, items ordered, etc, total spent
// ... this func can get re-used a bunch
func getSessionInformation(braceletID string) error {
	con, err := sql.Open("mysql", user+":"+password+"@/"+database)
	if err != nil {
		return err
	}
	defer con.Close()

	// depending on how we setup the database, either one row:
	row := con.QueryRow("select itemsOrdered from sometable where id=?", braceletID)
	items := new(TheStructToFill)            // struct gets rendered as JSON for the js to consume
	if err := row.Scan(&items); err != nil { // will require tweaking, see example in link
		return err
	}

	// or multiple rows (
	rows, err := con.Query("select a, b from item where p1=? and p2=?", "someVariable", "anotherVariable")
	if err != nil {
		return err
	}

	// this should be a slice ... see example...
	items = new(TheStructToFill)
	var someFood string
	for rows.Next() {
		err = rows.Scan(someFood)
		if err != nil {
			return err
		}

		//items = append(items, someFood) // TODO fix!
	}
	return nil // and some populated struct
}

type AllTheFood struct {
	// put arrays of:
	// drink
	// food
	// other
	// scrubs?
}

func whatIsOnTheMenu() (AllTheFood, error) {
	yum := letsJustHardcodeAllTheItemsForNow()
	return yum, nil
}

func letsJustHardcodeAllTheItemsForNow() AllTheFood {
	return AllTheFood{}
}
