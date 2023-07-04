package main

import (
    "bufio"
    "database/sql"
    "errors"
    "fmt"
    "os"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    args := readPromptArgs()

    err := validatePromptArgs(args)
    if err != nil {
        fmt.Println("Error validating arguments:", err)
        return
    }

    db, err := connectToDatabase(args)
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close()

    fmt.Printf("You are connected to the database %s\n", args.Database)
}

type PromptArgs struct {
    Hostname string
    Username string
    Password string
    Database string
}

func validatePromptArgs(args PromptArgs) error {
    if args.Hostname == "" {
        return errors.New("Hostname is required")
    }
    if args.Username == "" {
        return errors.New("Username is required")
    }
    if args.Database == "" {
        return errors.New("Database name is required")
    }
    return nil
}

func readPromptArgs() PromptArgs {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Hostname: ")
    hostname, _ := reader.ReadString('\n')
    hostname = strings.TrimSpace(hostname)

    fmt.Print("Username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("Password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    fmt.Print("Database: ")
    database, _ := reader.ReadString('\n')
    database = strings.TrimSpace(database)

    args := PromptArgs{
        Hostname: hostname,
        Username: username,
        Password: password,
        Database: database,
    }

    return args
}

func connectToDatabase(args PromptArgs) (*sql.DB, error) {
    connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", args.Username, args.Password, args.Hostname, args.Database)

    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}
