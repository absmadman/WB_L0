package db

const (
	createTable = `CREATE TABLE IF NOT EXISTS temp (id SERIAL PRIMARY KEY, message VARCHAR NOT NULL)`
	insertTable = "INSERT INTO temp (message) VALUES ('%s') RETURNING id"
	selectById  = `SELECT * FROM temp WHERE id = $1`
	getTotal    = `SELECT COUNT(*) FROM temp`
)
