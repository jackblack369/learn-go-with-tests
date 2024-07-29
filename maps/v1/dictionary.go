package main

// Dictionary store definitions to words.
type Dictionary map[string]string

type AsianCountry map[int]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) string {
	return d[word]
}

// Add a word and definition to the dictionary.
func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}

// Update a word's definition in the dictionary.
func (d Dictionary) Update(word string, definition string) {
	d[word] = definition
}

// Delete a word from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// Search find country by number.
func (ac AsianCountry) Search(id int) string {
	return ac[id]
}

// Add a country to the dictionary.
func (ac AsianCountry) Add(id int, country string) {
	ac[id] = country
}
