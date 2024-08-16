package board

type Board2 [9]*uint8

type Move string

// func (m Move) IsValid() {
// 	p := strings.ToLower(string(m[0]))

// 	if strings.Compare(p, "x") != 0 && strings.Compare(p, "o") != 0 {
// 		return fmt.Errorf("error happend")
// 	}

// }

// func (b Board2) HasBeenSet(p int8) (int, bool) {

// 	v := b[p]

// 	if v == nil {
// 		return 0, false
// 	}

// 	return *v, true

// }
