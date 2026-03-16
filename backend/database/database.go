package database

import (
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/sqlserver"
    "gorm.io/gorm"
    "backend/models"
    "database/sql"
)

// DB is the global database connection
var DB *gorm.DB

// SQLDB is the underlying sql.DB for advanced operations (for being able to run queries without ORM)
var SQLDB *sql.DB

func Connect() {
    DB = connectWithEnv()
}

// Special connection function, used for tests. Runs in-memory.
func ConnectTest() *gorm.DB {
    db := sqlite.Open(":memory:") // in-memory for tests
    gormDB, err := gorm.Open(db, &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect test DB:", err)
    }
    if err := gormDB.AutoMigrate(&models.Goal{}, &models.Percentage{}); err != nil {
        panic("failed to migrate database: " + err.Error())
    }

    // Get the underlying sql.DB for advanced operations
    SQLDB, err = gormDB.DB()
    if err != nil {
        panic(err)
    }

    DB = gormDB
    return DB
}

// Connects to the database based on environment variables. Support SQLite and MySQL.
// Connects to SQLite, MySQL, or MSSQL depending on DB_TYPE
func connectWithEnv() *gorm.DB {
    dbType := os.Getenv("DB_TYPE")
    conn := os.Getenv("DB_CONNECTION")

    var dialector gorm.Dialector

    switch dbType {
    case "mysql":
        dialector = mysql.Open(conn)

    case "mssql":
        // Example MSSQL connection string:
        // "sqlserver://sa:Password!123@localhost:1433?database=MyDB"
        if conn == "" {
            log.Fatal("DB_CONNECTION must be set for MSSQL")
        }
        dialector = sqlserver.Open(conn)

    default:
        if conn == "" {
            conn = "app.db"
        }
        dialector = sqlite.Open(conn)
    }

    gormDB, err := gorm.Open(dialector, &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate the models
    if err := gormDB.AutoMigrate(&models.Goal{}, &models.Percentage{}); err != nil {
        panic("failed to migrate database: " + err.Error())
    }

    SQLDB, err = gormDB.DB()
    if err != nil {
        panic(err)
    }

    // If in debug mode, drop all tables first
    if os.Getenv("GIN_MODE") == "debug" {
        log.Println("Debug mode: dropping all tables")
        _ = gormDB.Migrator().DropTable(&models.Goal{}, &models.Percentage{})
        _ = gormDB.AutoMigrate(&models.Goal{}, &models.Percentage{})
    }

    return gormDB
}
