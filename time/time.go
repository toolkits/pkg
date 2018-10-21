package time

import (
	"fmt"
	"time"
)

func Duration(now, before int64) string {
	d := now - before
	if d <= 60 {
		return "just now"
	}

	if d <= 120 {
		return "1 minute ago"
	}

	if d <= 3600 {
		return fmt.Sprintf("%d minutes ago", d/60)
	}

	if d <= 7200 {
		return "1 hour ago"
	}

	if d <= 3600*24 {
		return fmt.Sprintf("%d hours ago", d/3600)
	}

	if d <= 3600*24*2 {
		return "1 day ago"
	}

	return fmt.Sprintf("%d days ago", d/3600/24)
}

func Format(ts int64, pattern ...string) string {
	defp := "2006-01-02 15:04:05"
	if pattern != nil && len(pattern) > 0 {
		defp = pattern[0]
	}
	return time.Unix(ts, 0).Format(defp)
}
