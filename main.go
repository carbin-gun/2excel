package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io"

	"regexp"

	"path/filepath"

	"github.com/tealeg/xlsx"
)

const (
	CSV          = "/Users/fanfu/query_result.txt"
	TargetNaming = "%s.xlsx"
)

var REG *regexp.Regexp

func init() {
	REG = regexp.MustCompile(`[\t,]`)
}
func main() {
	var file *xlsx.File
	var sheet *xlsx.Sheet

	csv, err := os.Open(CSV)
	if err != nil {
		panic(err)
	}
	defer csv.Close()

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("[2excel] begin convert...")
	buf := bufio.NewReader(csv)
	for {
		line, err := buf.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		//file read finished.
		if err != nil && err == io.EOF {
			break
		}
		// read going on.
		line = strings.TrimSpace(line)
		handleLine(sheet, line)
	}

	basename := filepath.Base(csv.Name())
	basenameAndExtension := strings.Split(basename, ".")
	target := fmt.Sprintf(TargetNaming, basenameAndExtension[0])
	err = file.Save(target)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("[2excel] xlsx file OK:", target)
}

func handleLine(sheet *xlsx.Sheet, line string) {
	items := REG.Split(line, -1)
	//items := strings.Split(line, "\t")
	row := sheet.AddRow()
	for _, item := range items {
		cell := row.AddCell()
		cell.Value = item
	}
}
