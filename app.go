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
	TargetNaming = "%s.xlsx"
)

var REG *regexp.Regexp

func init() {
	REG = regexp.MustCompile(`[\t,]`)
}

func DoConvert(command Command) {
	fmt.Printf("source file:[%s],Delimiter:[%s],dest:[%s]\n", command.SourceFile, command.Delimiter, command.Dest)
	var file *xlsx.File
	var sheet *xlsx.Sheet

	csv, err := os.Open(command.SourceFile)
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
	var myReg = new(MyReg)
	if command.Delimiter == "" {
		myReg.Reg = REG
	} else {
		myReg.Reg = regexp.MustCompile(command.Delimiter)
	}
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
		myReg.handleLine(sheet, line, command)
	}

	basename := filepath.Base(csv.Name())
	basenameAndExtension := strings.Split(basename, ".")
	targetFileName := fmt.Sprintf(TargetNaming, basenameAndExtension[0])
	fileToSave := buildSaveLocation(command.Dest, targetFileName)
	err = file.Save(fileToSave)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("[2excel] xlsx file OK:", fileToSave)

}

func buildSaveLocation(dir, filename string) string {
	return filepath.Join(dir, filename)
}

type MyReg struct {
	Reg *regexp.Regexp
}

func (m *MyReg) handleLine(sheet *xlsx.Sheet, line string, command Command) {
	items := m.Reg.Split(line, -1)
	row := sheet.AddRow()
	for _, item := range items {
		cell := row.AddCell()
		cell.Value = item
	}
}
