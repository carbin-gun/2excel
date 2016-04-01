package main

type Command struct {
	//the file that needs to convert
	SourceFile string
	//the delimiter between different fields of one row.
	Delimiter string
	//where to put the generated excel file
	Dest string
}
