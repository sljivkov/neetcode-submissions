type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder
	for _, st := range strs{
		wordLen := len(st)
		res.WriteString(strconv.Itoa(wordLen));res.WriteString("#");res.WriteString(st)

	}

	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0

	lengy:=""
	for i < len(encoded) {
		if encoded[i] == '#' {

			leng, _ := strconv.Atoi(lengy)

			res = append(res, encoded[i+1: i+1+leng])

			lengy = ""

			i = i+1+leng

			continue
		}

		lengy += string(encoded[i])
		i += 1
	}

	return res
}

