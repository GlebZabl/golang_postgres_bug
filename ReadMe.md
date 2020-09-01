# Bug in "database/sql" + "github.com/lib/pq"
when i tried to select rows from postgreSQL db with GoLang libs "database/sql" and "github.com/lib/pq"
i found a bug. I used pattern "SELECT column_name FROM table_name WHERE column_name IN($1,$2...$n)", when number of 
arguments was really huge,i`ve got error (" sql: expected 0 arguments, got 65536", numbers in error message depends 
on args number), so i write this project and explore this strange bug.

# how to reproduce
- create a table in your postgreSQL database with script "create_db.sql" 
- clone this repository, modify 21st string in main.go(put your database connection string there)
- run programs few times with different args(change a parameter in 12 string of "main.go")

## environment
i have go 1.14 on my pc and use postgres:11-alpine docker image to deploy db

