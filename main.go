package main

import (
	"bufio"
	"fmt"
	"log"
	"nick_moore/spell-checker/checker"
	"os"
	"strconv"
)

func main() {
	local_dict := []string{}
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Please provide a dictionary file and a document file")
		os.Exit(1)
	}

	dictionary, err_dict := os.Open(args[0])
	document, err_doc := os.Open(args[1])

	if err_dict != nil {
		log.Fatal(err_dict)
	}

	if err_doc != nil {
		log.Fatal(err_doc)
	}

	defer func() {
		if err_dict = dictionary.Close(); err_dict != nil {
			log.Fatal(err_dict)
		}

		if err_doc = document.Close(); err_doc != nil {
			log.Fatal(err_doc)
		}
	}()

	scanner_dict := bufio.NewScanner(dictionary)
	scanner_doc := bufio.NewScanner(document)

	for scanner_dict.Scan() {
		local_dict = append(local_dict, scanner_dict.Text())
	}

	spell_checker := checker.CheckSpelling(local_dict, scanner_doc)
	spell_suggestions := checker.GetSuggestions()

	fmt.Println("Mispellings...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i, word := range spell_checker {
		fmt.Println(strconv.Itoa(i+1) + ": " + word)
	}

	fmt.Println(spell_suggestions)

}
