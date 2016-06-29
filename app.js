var mysql	= require('mysql');
var connection	= mysql.createConnection({
	host	: 'banya-mysql',
	user	: 'root',
	password: 'passwd',
	database: 'banya'
});

connection.connect();

connection.query('SELECT name, price FROM foodstuffs', function(err, rows, fields) {
	if (err) throw err;

	console.log('The solution is: ', rows);
	console.log('The solution is: ', fields);
});
connection.end();
