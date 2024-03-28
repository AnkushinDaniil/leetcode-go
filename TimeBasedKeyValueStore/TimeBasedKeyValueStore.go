package timeBasedKeyValueStore

type Value struct {
	timestamp int
	value     string
}

type TimeMap struct {
	data map[string][]Value
}

func Constructor() TimeMap {
	data := make(map[string][]Value, 0)
	return TimeMap{
		data: data,
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.data[key] = append(
		tm.data[key],
		Value{
			timestamp: timestamp,
			value:     value,
		},
	)
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.data[key]; !ok || timestamp < tm.data[key][0].timestamp {
		return ""
	}

	l, r := 0, len(tm.data[key])-1

	for l <= r {
		m := (l + r) / 2
		if tm.data[key][m].timestamp > timestamp {
			r = m - 1
		} else if tm.data[key][m].timestamp < timestamp {
			l = m + 1
		} else {
			return tm.data[key][l].value
		}
	}

	return tm.data[key][l-1].value
}
