package main

import (
	"bufio"
	"fmt"
	"log"
	"nick_moore/spell-checker/checker"
	"os"
	"strconv"
	"strings"
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
		line := strings.TrimSpace(scanner_dict.Text())
		dict_line := strings.Split(line, " ")
		for _, dict_word := range dict_line {
			local_dict = append(local_dict, strings.TrimSpace(dict_word))
		}
	}

	spell_checker := checker.CheckSpelling(local_dict, scanner_doc)
	spell_suggestions := checker.GetSuggestions(local_dict)
	spell_doc_contexts := checker.GetDocContexts()

	fmt.Println("Mispellings...")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i, word := range spell_checker {
		fmt.Println(strconv.Itoa(i+1) + ": " + word)
		fmt.Println(strings.Join(spell_doc_contexts[i], " "))
		fmt.Println("\n--- SUGGESTIONS ---")
		fmt.Println(spell_suggestions[i])
		fmt.Println()
	}

}
