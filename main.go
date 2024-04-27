package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(Domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	
	mxRecord, err := net.LookupMX(Domain)
	
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(Domain)

	if err != nil {
		log.Printf("Error: %v\n",err)
	}
	
	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + Domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v", hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord")
	
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)
	}

}
