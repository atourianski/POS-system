package main

import (
	"fmt"
	"net/http"
)

// these endpoints can be called by js, see
// https://www.kirupa.com/html5/making_http_requests_js.htm

func main() {
	// start the day.
	http.HandleFunc("/initializeDay", initializeDay)

	// new user arrives. we call it a session & send in the bracelet ID
	http.HandleFunc("/initializeSession", initializeSession)

	// everytime they want to order, enter ID and serve up invoice
	// with option to add new items/close|merge sessions
	http.HandleFunc("/displaySessionSummary", displaySessionSummary)

	// if adding new items, list them (keep track internally of the ID)
	http.HandleFunc("/listAvailableItems", listAvailableItems)

	// select some items from above and add them
	http.HandleFunc("/addItemsToASession", addItemsToASession)

	// this option can be triggered by same screen rendered from
	// displaySessionSummary
	http.HandleFunc("/closeSession", closeSession)

	// combine the total of two sessions at closing
	http.HandleFunc("/mergeSessions", mergeSessions)

	// create a user friendly (pdf) & printable invoice
	http.HandleFunc("/createInvoice", createInvoice)

	// serve it all up
	http.ListenAndServe(":8080", nil)
}

func initializeDay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "basically a test or pre-check that the day will run smooth. eventually want to enter in the foods that will be available for that day. if, say, a bracelet is lost or locker not working, can enter invalid #'s")

	if r.Method == "GET" {
		if err := CheckThatTheDayIsReady(); err != nil {
			http.Error(w, "errrors", http.StatusInternalServerError)
		}
	}

}

func initializeSession(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this endpoint should be a POST that checks to ensure the bracelet # isn't already being used within 'a day' then initialize a session for that user.")

	if r.Method == "POST" {
		// TODO read braceletID from post
		braceletID := "51"
		if err := addNewSession(braceletID); err != nil {
			http.Error(w, "errors", http.StatusInternalServerError)
		}
	}

	// the user is now active ... they can order food.
}

// the ui for back of the house should, IMO, start with entering the number
// this will pull up all the user info & *can* serve as a way to prevent mistakes
// a post Request with bracelet ID (like initSession) gets this function started
// this func can probably be re-used for creating invoices
func displaySessionSummary(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "invoice info; options: ADD ITEMS, CLOSE")
	if r.Method == "POST" { /* code in here */
	}
}

// ^ a big button called "Add Items" will transfer the user to a list of available items.

// before adding food to an order, we must list it
func listAvailableItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "query the db for all the items wanted")

	if r.Method == "GET" {
		_, err := whatIsOnTheMenu() // much simpler to hardcode for now
		if err != nil {
			http.Error(w, "werror", http.StatusInternalServerError)
		}
		theFood := "all the food printed out nicely"
		fmt.Fprintf(w, "Please select from: %s", theFood)

		// send theFood back to the js function that called it in a JSON payload
	}
}

// ^^ the html form posting to js will have checkboxes (along with the braclet #),
// packaged up as JSON, and sent to the following function.

// once the iterms are selected, they need to be written to the cumulative invoice
// invoice being the "thing" that is derived from a session at the end of it, basically
func addItemsToASession(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { /* code in here */
	}

	// ensure to throw an error if the items *was not* added to the session
	// well always throw errors, but checking here is especially important
}

// after adding a bunch of items, we close the session. this displays the invoice
// with total owed, presents the option to merge 'sessions',
func closeSession(w http.ResponseWriter, r *http.Request) {
	// under the hood, this should move the "active" bracelet & the session's uniqueID,
	// alongside all invoice & items info to "inactive", freeing up the bracelet_id
	if r.Method == "POST" { /* code in here */
	}
}

func mergeSessions(w http.ResponseWriter, r *http.Request) {} // can be used internally when invoice is created

// optionally create a printable pdf or whatever.
func createInvoice(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { /* code in here */
	}
}
