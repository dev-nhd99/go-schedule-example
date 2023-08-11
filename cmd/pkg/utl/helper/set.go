package helper


type SetS struct {
	MapS  map[string]bool
}


func NewSetS() SetS {
	mapS := make(map[string]bool)
	return SetS{
		MapS: mapS,
	}
}

func (S SetS) Add(ss string) {
	S.MapS[ss] = true
}

func (S SetS) ToSlice() (slice []string) {
	for k, v := range S.MapS{
		if v {
			slice = append(slice, k)
		}
	}
	return slice
}

type SetI struct {
	MapI  map[int]bool
}

func NewSetI() SetI {
	mapI := make(map[int]bool)
	return SetI{
		MapI: mapI,
	}
}

func (S SetI) Add(i int) {
	S.MapI[i] = true
}

func (S SetI) ToSlice() (slice []int) {
	for k, v := range S.MapI{
		if v {
			slice = append(slice, k)
		}
	}
	return slice
}