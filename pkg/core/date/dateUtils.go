package date

import "time"

var (
	defaultFormat string = time.RFC3339
)

func Now() time.Time {
	return time.Now()
}

func IsAfterNow(date time.Time) bool {
	return date.After(time.Now())
}

func IsBeforeNow(date time.Time) bool {
	return date.Before(Now())
}

func StringToTime(timeString string) (time.Time, error) {
	return StringToTimeWithFormat(timeString, defaultFormat)
}

func StringToTimeWithFormat(timeString string, format string) (time.Time, error) {
	return time.Parse(format, timeString)
}

func TimeToString(time time.Time) string {
	return time.Format(defaultFormat)
}

func TimeToStringWithFormat(time time.Time, format string) string {
	return time.Format(format)
}

func AddMinuts(date time.Time, minuts int) time.Time{
	return date.Add(time.Duration(minuts) * time.Minute)
}
