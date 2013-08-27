package luhnmodn

import "testing"

func TestChecksummed(t *testing.T) {
  var tests = []struct {
    s, want string
    base int
  }{
    {"7992739871","79927398713",10},
    {"d4kfd51a3","d4kfd51a39",36},
  }
  for _, c := range tests {
    got := Checksummed(c.s,c.base)
    if got != c.want {
      t.Errorf("Checksummed(%q,%d) == %q, want %q",c.s,c.base,got,c.want)
    }
  }
}

func TestValid(t *testing.T) {
  var tests = []struct {
    s string
    want bool
    base int
  }{
    {"79927398713",true,10},
    {"79917398713",false,10},
    {"d4kfd51a39",true,36},
    {"d4fkd51a39",false,36},
    {"d4fd51a39",false,36},
    {"4dkfd51a39",false,36},
  }
  for _, c := range tests {
    got, _ := Valid(c.s,c.base)
    if got != c.want {
      t.Errorf("Valid(%q,%d) == %t, want %t",c.s,c.base,got,c.want)
    }
  }
}