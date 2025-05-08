package main

import (
	"bufio"
	"delete_table_console/config"
	"fmt"
	"os"
	"sync"
	"time"
)

var serverConnection config.ConnectionParams
var targetTableName string

func processDelete(listBackup []config.DeleteList) {
	startProcess := time.Now()
	totalBackup := len(listBackup)
	fmt.Printf("\n\n👻  : Please wait while I help backup your %v table", totalBackup)
	fmt.Printf("\n\n👻  : Start at 🕑 %s \n\n", startProcess.Format("2006-01-02 15:04:05"))

	maxProccess := make(chan struct{}, 8)
	var wg sync.WaitGroup

	for _, tbl := range listBackup {
		maxProccess <- struct{}{} // acquire a slot
		wg.Add(1)

		go func(tbl config.DeleteList) {
			defer wg.Done()
			defer func() { <-maxProccess }() // release the slot

			config.DropTable(serverConnection, tbl.TableName)
		}(tbl)
	}

	wg.Wait() // Wait for all to finish

	elapsed := time.Since(startProcess) // Calculate duration
	fmt.Printf("\n\n👻  : ✅ Done with execution time: 🕑 %s", elapsed)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("👻 : Welcome to Delete {MS SQL SERVER} table program 🎃🎃🎃...\n")
	serverConnection = config.GetInputSourceServer(reader)
	listDatabaseServer := config.GetListDatabaseInServer(serverConnection)
	serverConnection.Database = config.GetInputSourceDB(reader, listDatabaseServer, true)

	fmt.Printf("\n 🌐  You're selected Source: %s - %s \n\n", serverConnection.Server, serverConnection.Database)

	inputTableName, listTableForBackup := config.GetInputTableName(reader, serverConnection)
	targetTableName = inputTableName

	if !config.AskingContinueBackupProcess(reader) {
		fmt.Println("\n👻 : Thank you for interact with me ... ")
		fmt.Println("👻 : Have A Nice Day ... ")
		return
	}

	processDelete(listTableForBackup)
}
