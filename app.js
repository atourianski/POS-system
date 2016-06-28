var mysql	= require('mysql');
var connection	= mysql.createConnection({
	host	: 'banya-mysql',
	user	: 'root',
	password: 'passwd',
	database: 'banya'
});

connection.connect();

connection.query('SELECT 1 + 1 AS solution', function(err, rows, fields) {
	if (err) throw err;

	console.log('The solution is: ', rows [0].solution);
});
connection.end();
