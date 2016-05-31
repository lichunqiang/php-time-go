package time

import (
	"time"
)

//This is equal to PHP `time()` function
//return current timestamp
func Time() int64 {
	return time.Now().Unix()
}

//equal to PHP `date()`
//May be to learn about slice and variable paramter func
func Date(format string, timestamp ...int64) string {
	println(format)
	mTimestamp := Time()

	if len(timestamp) > 1 {
		mTimestamp = timestamp[0]
	}

	return time.Unix(mTimestamp, 0).Format("2006-01-02 15:04:05")
}

//equal to PHP `strtotime(string $time, [int $now = time()])`
func Strtotime(t string) {

}
