package db

const (
	createTable = `CREATE TABLE IF NOT EXISTS temp (id SERIAL PRIMARY KEY, message VARCHAR NOT NULL)`
	insertTable = `INSERT INTO temp (message) VALUES ($1)`
	selectById  = `SELECT * FROM temp WHERE id = $1`
)
