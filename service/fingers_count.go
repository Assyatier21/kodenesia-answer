package service

func FingersCount(start int, end int) string {

	var (
		finger = map[int]string{
			1: "Jempol",
			2: "Telunjuk",
			3: "Tengah",
			4: "Manis",
			5: "Kelingking",
		}
		num int
	)

	num = end - start + 1
	if num > 8 {
		num -= 8
	}

	if num == 1 {
		return finger[1]
	}
	return finger[num%5]
}
