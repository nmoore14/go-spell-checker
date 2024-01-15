# Spell Checker - Nick Moore
## My Awesome Cross Platform Spell Checker

### How to Run
- macOS:
    + `cd` into the root of this project.
    + You should see a `spell-checker` and a `spell-checker.exe` (windows) and an assets folder
    + In the root folder run `./spell-checker assets/dictionary.txt assets/test_0.txt`
        - `assets/dictionary.txt` is the dictionary file provided for this assessment.
        - `assets/test_0.txt` is a file that created for testing
- Windows:
    + `cd` into the root of this project.
    + You should see a `spell-checker` and a `spell-checker.exe` (windows) and an assets folder
    + In the root folder run `.\spell-checker.exe .\assets\dictionary.txt .\assets\test_0.txt`
        - `assets/dictionary.txt` is the dictionary file provided for this assessment.
        - `assets/test_0.txt` is a file that created for testing

### Motivation
- Easily check your document for mispellings against our dictionary. Once checked, provide a list of mispellings along with suggestions to help the user correct their document.
- Cross Platform CLI tool should be easily installed and ran on macOS, Windows, and Linux...this is why I decided to go with the Go programming language.

### Pseudo
- Require a document to passed into the CLI in `.txt` form
    + Through an error with clear reason and exit the program
- Check the document against the dictionary.
    + If a dictionary is not supplied use the one stored in my assets folder (provided via the code test)
- Store the mispelled words and their locations (Line number and line position)
- Use the Levenshtein Distance to build a list of suggestions for close matches (5 closest matches)
    + Decided against top 5 due to not knowing context of the text...spell-checker isn't smart enought yet.

### TODO
- [x] Take the dictionary and document as CLI inputs
    - [x] error if they don't exist
- [x] Check the spelling for errors
    - [x] Build a list of mispelled words with row and column number
    - [x] Return the list
- [x] Check the incorrect list for word suggestions
    - [x] Loop through the incorrect list and build another list for suggestions using the below algorithm.
- [x] Handle proper nouns correctly.
- [x] Print some surrounding context to help the user pick a suggestion.

### Research
- (Levenshtein Distance)[https://en.wikipedia.org/wiki/Levenshtein_distance#:~:text=Informally%2C%20the%20Levenshtein%20distance%20between,considered%20this%20distance%20in%201965]
    + The number of single character edits between two words.
        > kitten -> sitting = 3
