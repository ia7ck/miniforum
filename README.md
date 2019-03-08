## development

0. Start PostgreSQL
1. `export DATA_SOURCE_NAME="dbname=your_db_name password=your_password"`
    + or in `db.go` replace `os.Getenv("DATA_SOURCE_NAME")` to above string 
2. `psql -f model/setup.sql -d your_db_name`
3. `go run main.go`
