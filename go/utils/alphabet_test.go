package utils_test

import (
	"github.com/mparaiso-snippets/snippets/go/utils"
	"reflect"
	"testing"
)

type any = interface{}

func TestAlphabet(t *testing.T) {
	alphabet := utils.NewAlphabet("ABC")
	expect(t, alphabet.Contains('A'), true)
	expect(t, alphabet.Contains('a'), false)
	expect(t, alphabet.ToChar(1), 'B')
	expect(t, alphabet.ToIndex('A'), 0)

}
func expect(t *testing.T, actual, expected any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, wanted %v", actual, expected)
	}
}
