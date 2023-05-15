package main

import (
	"log-reader-go/internal/config"
	"log-reader-go/internal/utils"
	"log-reader-go/pkg/color"
	"os"
	"time"
)

var (
	startTimeExec time.Time
	cLog          config.LogFile
)

func main() {
	defer end()

	/** Lendo os argumentos passados */
	err := utils.ReadArgs(&cLog)

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	/** Abrindo o arquivo de log */
	f, err := os.Open(cLog.Filename)

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	defer f.Close()

	filestat, err := f.Stat()

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	cLog.Name = filestat.Name()
	cLog.Size = filestat.Size()

	/** */

	/**
	if cLog.LogStartTime != nil {
		// checa o primeiro registro do arquivo (Data)
	}
	*/

	/**
	if cLog.LogEndTime != nil {
		// checa o último registro do arquivo (Data)
	}
	*/

	offset := cLog.Size - 1

	lastLine, err := utils.ReadLine(f, uint64(offset), true)

	if err != nil {
		color.PrintError(err.Error())
		return
	}

	utils.ParsingLineLog(string(lastLine))

	/** Ler a primeira linha e verificar se o inicio de tempo está depois */
	//
	/** Ler a última linha e verificar se o final de tempo está antes. */

	// fmt.Println(cLog)

}

func init() {
	startTimeExec = time.Now()
	color.PrintCyan("Started: ", startTimeExec.String())
}

func end() {
	diff := time.Since(startTimeExec)

	color.PrintCyan("Ended: ", time.Now().String())
	color.PrintYellow("Elapsed Time: ", diff.Truncate(time.Second).String())
}
