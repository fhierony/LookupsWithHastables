package main

import (
	"fmt"
)

func main() {
	// Make some names.
	employees := []Employee{
		{"Ann Archer", "202-555-0101"},
		{"Bob Baker", "202-555-0102"},
		{"Cindy Cant", "202-555-0103"},
		{"Dan Deever", "202-555-0104"},
		{"Edwina Eager", "202-555-0105"},
		{"Fred Franklin", "202-555-0106"},
		{"Gina Gable", "202-555-0107"},
		{"Herb Henshaw", "202-555-0108"},
		{"Ida Iverson", "202-555-0109"},
		{"Jeb Jacobs", "202-555-0110"},
	}

	hashTable := NewChainingHashTable(10)
	for _, employee := range employees {
		hashTable.set(employee.name, employee.phone)
	}
	hashTable.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hashTable.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hashTable.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hashTable.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hashTable.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hashTable.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hashTable.get("Fred Franklin"))
}

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

type Employee struct {
	name  string
	phone string
}

type ChainingHashTable struct {
	numBuckets int
	buckets    [][]*Employee
}

// NewChainingHashTable Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(numBuckets int) *ChainingHashTable {
	return &ChainingHashTable{numBuckets, make([][]*Employee, numBuckets)}
}

// Display the hash table's contents.
func (hashTable *ChainingHashTable) dump() {
	for key, value := range hashTable.buckets {
		fmt.Printf("Bucket %d :\n", key)
		for _, employee := range value {
			fmt.Printf("\t%s: %s\n", employee.name, employee.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hashTable *ChainingHashTable) find(name string) (int, int) {
	index := hash(name) % hashTable.numBuckets
	for key, employee := range hashTable.buckets[index] {
		if employee.name == name {
			return index, key
		}
	}
	return index, -1
}

// Add an item to the hash table.
func (hashTable *ChainingHashTable) set(name string, phone string) {
	bucket, index := hashTable.find(name)

	// Check if key is already present, if yes, just update
	if index >= 0 {
		hashTable.buckets[bucket][index].phone = phone
	} else {
		// Add a new employee
		hashTable.buckets[bucket] = append(hashTable.buckets[bucket], &Employee{name, phone})
	}
}

// Return an item from the hash table.
func (hashTable *ChainingHashTable) get(name string) string {
	bucket, index := hashTable.find(name)

	// Check if key is present, if yes, return phone
	if index >= 0 {
		return hashTable.buckets[bucket][index].phone
	}

	return ""
}

// Return true if the person is in the hash table.
func (hashTable *ChainingHashTable) contains(name string) bool {
	_, index := hashTable.find(name)

	if index == -1 {
		return false
	}

	return true
}

// Delete this key's entry.
func (hashTable *ChainingHashTable) delete(name string) {
	bucket, index := hashTable.find(name)

	// Check if key is present, if yes, return phone
	if index >= 0 {
		currentBucket := hashTable.buckets[bucket]
		currentBucket[index] = currentBucket[len(currentBucket)-1]
		currentBucket = currentBucket[:len(currentBucket)-1]
		hashTable.buckets[bucket] = currentBucket
	}
}
