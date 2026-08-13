type timeMapEntry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	mapa map[string][]timeMapEntry
}

func Constructor() TimeMap {
	return TimeMap{
		mapa: make(map[string][]timeMapEntry, 0),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.mapa[key] = append(this.mapa[key], timeMapEntry{
		value:     value,
		timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	var (
		arr = this.mapa[key]
		l   = 0
		r   = len(arr) - 1
		temp string
	)

	for l <= r {
		curr := (l + r) / 2

		if arr[curr].timestamp <= timestamp {
			l = curr + 1
			temp = arr[curr].value

		} else {
			r = curr - 1
		}	
	}

	return temp 
}