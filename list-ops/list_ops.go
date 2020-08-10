package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

func (l IntList) Filter(fn predFunc) IntList {
	var ret = make([]int, 0, len(l))
	for _, num := range l {
		if fn(num) {
			ret = append(ret, num)
		}
	}

	return IntList(ret)
}

func (l IntList) Map(fn unaryFunc) IntList {
	var ret = make([]int, 0, len(l))
	for _, num := range l {
		ret = append(ret, fn(num))
	}
	return IntList(ret)
}

func (l *IntList) Append(m IntList) IntList {
	*l = append([]int(*l), m...)
	return *l
}

func (l *IntList) Concat(m []IntList) IntList {
	for _, n := range m {
		*l = append([]int(*l), n...)
	}
	return *l
}

func (l IntList) Reverse() IntList {
	var i, j int = 0, len(l) - 1
	for i < j {
		l[i], l[j] = l[j], l[i]
		i++
		j--
	}
	return l
}

func (l IntList) Foldr(fn binFunc, a int) int {
	for _, v := range []int(l.Reverse()) {
		a = fn(v, a)
	}
	return a
}

func (l IntList) Foldl(fn binFunc, a int) int {
	for _, v := range []int(l) {
		a = fn(a, v)
	}
	return a
}

func (l IntList) Length() int {
	return len(l)
}
