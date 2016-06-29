var http = require('http');
var mysql = require('mysql');
var fs = require('fs');
var formidable = require('formidable');
var util = require('util');

var server = http.createServer(function (req, res) {
    if (req.method.toLowerCase() == 'get') {
        displayForm(res);
    } else if (req.method.toLowerCase() == 'post') {
        processFormFieldsIndividual(req, res);
    }

});

function displayForm(res) {
    console.log("trying to display visit.html");
    fs.readFile('visit.html', function (err, data) {
        res.writeHead(200, {
            'Content-Type': 'text/html',
                'Content-Length': data.length
        });
        res.write(data);
        res.end();
    });
}


// TODO rename function
function processFormFieldsIndividual(req, res) {
    //Store the data from the fields in your data store.
    //The data store could be a file or database or any other store based
    //on your application.
    var fields = [];
    var form = new formidable.IncomingForm();
    form.on('field', function (field, value) {

       	//console.log(field);
       	//console.log(value);
        fields[field] = value;
       	console.log(fields);
	var createNewVisit = function(callback) {
	    callback(null, fields)
	};
	createNewVisit(createNewVisitCallback);
    });

    form.on('end', function () {
        res.writeHead(200, {
            'content-type': 'text/plain'
        });
    });
    // [zr] what does this do?
    form.parse(req);
}


var createNewVisitCallback = function(err, fields) {
	if (err) throw err;
	console.log('fields: '+fields);
	
	var bnum = fields["bnum"];
	console.log('Creating new visit with bracelet number: '+bnum);
	var connection	= mysql.createConnection({
		host	: 'banya-mysql',
		user	: 'root',
		password: 'passwd',
		database: 'banya'
	});
	
	connection.connect();
	console.log('connected, inserting row to visit  table.');

	var sql = "INSERT INTO visit (date, unique_id, bracelet_num, entry_time) VALUES (?, ?, ?, ?)";
	var inserts = [getTodaysDate(), getUniqueID(), bnum, getTimeNow()];
	sql = mysql.format(sql, inserts);

	connection.query(sql, function(err, rows, fields) {
		if (err) throw err;
	
		console.log('The solution is: ', rows);
		console.log('The solution is: ', fields);
	});
	connection.end();
}

server.listen(8080);
console.log("server listening on 8080");

// ---------------------------------------
// helpers (TODO move to seperate files)

function getTodaysDate() {

    var date = new Date();

    var year = date.getFullYear();

    var month = date.getMonth() + 1;
    month = (month < 10 ? "0" : "") + month;

    var day  = date.getDate();
    day = (day < 10 ? "0" : "") + day;

    return year + ":" + month + ":" + day;

}

function getTimeNow() {

    var date = new Date();

    var hour = date.getHours();
    hour = (hour < 10 ? "0" : "") + hour;

    var min  = date.getMinutes();
    min = (min < 10 ? "0" : "") + min;

    var sec  = date.getSeconds();
    sec = (sec < 10 ? "0" : "") + sec;

    return hour + ":" + min + ":" + sec;

}

//]TODO GENERATE RANDOM NUMBER
function getUniqueID() {
	
	var min = 10000;
	var max = 99999;	

	return Math.random() * (max - min) + min;
}
