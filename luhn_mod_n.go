package luhnmodn

import (
    "bytes"
    "fmt"
    "strconv"
)

func CheckDigit(number string, base int) (digit int) {
    digits := make([]int64, (len(number)))
    var err error
    for i := 0; i < len(number); i++ {
        digits[i], err = strconv.ParseInt(string(number[i]), base, 64)
        if err != nil {
            fmt.Println(err)
        }
    }
    sum := int64(0)
    for i := len(digits) - 1; i >= 0; i-- {
        sum += 2 * digits[i] / int64(base) + 2 * digits[i] % int64(base)
        i--
        if i >= 0 {
            sum += digits[i]
        }
    }
    digit = base - int(sum%int64(base))
    if base == digit {
        digit = 0
    }
    return
}

func Checksummed(number string, base int) (check string) {
    var buffer bytes.Buffer
    buffer.WriteString(number)
    buffer.WriteString(strconv.FormatInt(int64(CheckDigit(number, base)), base))
    return buffer.String()
}

func Valid(number string, base int) (valid bool) {
    valid = false
    given_digit, err := strconv.ParseInt(number[len(number)-1:len(number)], base, 64)
    if err != nil {
        fmt.Println(err)
    }
    return int64(CheckDigit(number[:len(number)-1], base)) == given_digit
}
