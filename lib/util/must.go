package util

func Require(e error) {
	if e != nil {
		panic(e)
	}
}

func MustRead(b []byte, err error) []byte {
	Require(err)
	return b
}
