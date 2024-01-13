# Spell Checker - Nick Moore
## My Awesome Cross Platform Spell Checker

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

### TODO
- [x] Take the dictionary and document as CLI inputs
    - [x] error if they don't exist
- [x] Check the spelling for errors
    - [x] Build a list of mispelled words with row and column number
    - [x] Return the list
- [ ] Check the incorrect list for word suggestions
    - [ ] Loop through the incorrect list and build another list for suggestions using the below algorithm.
- [x] Handle proper nouns correctly.

### Research
- (Levenshtein Distance)[https://en.wikipedia.org/wiki/Levenshtein_distance#:~:text=Informally%2C%20the%20Levenshtein%20distance%20between,considered%20this%20distance%20in%201965]
    + The number of single character edits between two words.
        > kitten -> sitting = 3
