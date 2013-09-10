mt19937_64
==========

64bit Mersenne Twister (MT19937-64) written in Go. Implements the "Source interface" from "math/rand". Example use of the interface can be found in `tests/mt19937_64_concurency.go`.

### Testing

Reproduces original test vectors [http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt19937-64.out.txt](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt19937-64.out.txt). See `tests/mt19937_64_test_vectors.go` and `data/mt19937-64.out.txt` for details. Concurency safety is tested in `tests/mt19937_64_concurency.go`.

### COPYRIGHT / LICENSE

The MIT License (MIT)
