package timeBasedKeyValueStore

type Node struct {
	timestamp int
	value     string
}

type Data []*Node

func NewData(capacity int) *Data {
	data := make(Data, 0, capacity)
	return &data
}

func (d *Data) Set(n *Node) {
	*d = append(*d, n)
}

func (d Data) Get(t int) string {
	if t < d[0].timestamp {
		return ""
	}
	l, r := 0, len(d)-1
	for l <= r {
		m := (l + r) / 2
		if t == d[m].timestamp {
			return d[m].value
		}
		if t > d[m].timestamp {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return d[l-1].value
}

type TimeMap map[string]*Data

func Constructor() TimeMap {
	return make(TimeMap, 4)
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := (*t)[key]; !ok {
		(*t)[key] = NewData(1)
	}
	(*t)[key].Set(&Node{
		timestamp: timestamp,
		value:     value,
	})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if data, ok := (*t)[key]; ok {
		return data.Get(timestamp)
	}
	return ""
}
