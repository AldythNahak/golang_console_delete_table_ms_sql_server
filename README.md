# golang_console_delete_table_ms_sql_server
A console-based Go application to back up tables from a Microsoft SQL Server database. Supports selective optional deletion table.

---

## 🧩 Features

- 🔄 Delete individual or all tables from a specified SQL Server database
- 🧹 Optionally delete tables
- ✅ Supports concurrent delete processing
- 🖥 Simple CLI interface for automation and scripting

---

## 🚀 Requirements

- Go 1.18 or higher
- SQL Server 2012 or later
- Access credentials with permission to read and drop tables

---

## ⚙️ Configuration

Edit the configuration in `config/config.go` or via environment variables if supported.

Key values include:
- Source server & database
- Table selection mode

---

## 📦 Usage

```bash
# Clone the repo
git clone https://github.com/AldythNahak/golang_console_delete_table_ms_sql_server.git
cd golang_console_delete_table_ms_sql_server

# Build or run the app
go run main.go
```
---

## 📦 Dependencies

The following Go packages are used:
```bash
go get github.com/denisenkom/go-mssqldb
go get github.com/jmoiron/sqlx
go get golang.org/x/term
```