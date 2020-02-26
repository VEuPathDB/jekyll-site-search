package util

func TrimLeft(tgt string, remove uint8) string {
	ln := len(tgt)

	for i := 0; i < ln; i++ {
		if tgt[i] != remove {
			return tgt[i:]
		}
	}

	return ""
}

func TrimRight(tgt string, remove uint8) string {
	for i := len(tgt) - 1; i > -1; i-- {
		if tgt[i] != remove {
			return tgt[:i+1]
		}
	}

	return ""
}

func Trim(tgt string, remove uint8) string {
	return TrimLeft(TrimRight(tgt, remove), remove)
}
