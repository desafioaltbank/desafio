// handlers/prime_handler_test.go
package handlers

import "testing"

//Se false falha o teste
func TestPrimeNumbers(t *testing.T) {
	num := 104743
	if !numberPrime(num) {
		t.Errorf("Failed test for number %d, expected true", num)
	}
}

//Se true falha o teste
func TestNomPrimeNumbers(t *testing.T) {
	num := 4
	if numberPrime(num) {
		t.Errorf("Failed test for number %d, expected false", num)
	}
}
