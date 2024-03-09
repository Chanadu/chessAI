package extras

import "strconv"

func FormatInt32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}
