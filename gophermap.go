package main

import (
	"fmt"
	"gophermap/parser"
	"io/ioutil"
	"os"
)

func main() {
	usage := `
	For the moment gophermap will just take a scan file and output 
	enumerated services in a format something like :
	| host-ip | port | service version | 

	.\gophermap [FILE-TYPE] [file.ext]
	 file type can be [nessus-csv, nessus-csv-web, nmap, rumble]
		note that while rumble uses the nmap xml format it also seems to be hiding 
		a lot of version data in banners in the JSON blob. Thus the dedicated
		format. 
	`

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(usage)
		os.Exit(3)
	}

	// gophermap currently accepts 3 scan file formats [nmapxml, rumblexml, nessus csv]
	switch args[0] {
	case "nessus-csv":
		// check if file (2nd arg) exists; exit if not
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Scan file does not exist")
			os.Exit(1)
		}

	case "nessus-csv-web":
		// check if file (2nd arg) exists; exit if not
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Scan file does not exist")
			os.Exit(1)
		}

		// parse the csv
		parser.NessusPrettyWeb(args[1])

	case "nessus-xml":
		// check if file (2nd arg) exists; exit if not
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Scan file does not exist")
			os.Exit(1)
		}

		// parse the csv
		parser.NessusPrettyXML(args[1])

	case "nmap":
		// check if file (2nd arg) exists; exit if not
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Scan file does not exist")
			os.Exit(1)
		}

		// open the file with ReadFile;
		fl, err := ioutil.ReadFile(args[1])
		if err != nil {
			fmt.Println("Error opening file!")
		}
		// pass opened file (byte slice)
		parser.NmapPrettyPrint(fl)

	case "rumble":
		// check if file (2nd arg) exists; exit if not
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Println("Scan file does not exist")
			os.Exit(1)
		}

		// ReadFile returns contents as byte slice
		fl, err := ioutil.ReadFile(args[1])
		if err != nil {
			fmt.Println("Error opening file!")
		}

		// parse file
		parser.RumblePrettyPrint(fl)

	default:
		// report and exit if not of the 3 file types above
		fmt.Println("Select  file type that exists")
		os.Exit(3)
	}

}
