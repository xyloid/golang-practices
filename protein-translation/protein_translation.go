/*
Package protein implements functions for protein translation
*/
package protein

import (
	"errors"
)

// ErrStop for stop codon
var ErrStop error = errors.New("Stop")

// ErrInvalidBase for invalid base
var ErrInvalidBase error = errors.New("Invalid Base")

// FromCodon maps a codon to a protein
func FromCodon(codon string) (string, error) {
	if protein, ok := table[codon]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

// FromRNA maps a rna to proteins
func FromRNA(rna string) ([]string, error) {
	runes := []rune(rna)
	var record = map[string]bool{}
	result := []string{}

	for i := 0; i < len(runes); i += 3 {
		codon := string(runes[i : i+3])
		protein, err := FromCodon(codon)
		switch err {
		case nil:
			if _, ok := record[protein]; !ok {
				record[protein] = true
				result = append(result, protein)
			}
		case ErrStop:
			return result, nil
		case ErrInvalidBase:
			return result, ErrInvalidBase
		}
	}

	return result, nil
}

var table = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}
