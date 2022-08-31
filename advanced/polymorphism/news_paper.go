package polymorphism

const NTimePerPage = 30

type newsPaper struct {
	infrmSrc
}

func (n newsPaper) read() int {
	return NTimePerPage * n.pgCnt
}
