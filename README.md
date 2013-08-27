# luhn mod n

An implementation of the [Luhn mod N](http://en.wikipedia.org/wiki/Luhn_mod_N_algorithm) algorithm for Go

## About

A learning project for Go by implementing Luhn mod N check digig algorithm. Should allow check digit generation & validation of numbers in various bases (10,16,36)

## Use

### Install

    go get https://github.com/glennsb/luhnmodn

### Quick Use

```Go
luhnmodn.Checksummed("7992739871",10) //Get back a string 79927398713
luhnmodn.Valid("d4fkd51a39",36) //false
luhnmodn.Valid("d4kfd51a39",36) //true
```

## License

(c) 2013 Stuart Glenn under MIT license