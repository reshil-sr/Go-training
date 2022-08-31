package polymorphism

const MTimePerPage = 5

type magazine struct {
	infrmSrc
}

func (m magazine) read() int {
	return MTimePerPage * m.pgCnt
}
