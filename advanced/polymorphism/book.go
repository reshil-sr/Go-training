package polymorphism

const BTimePerPage = 10

type book struct {
	bk infrmSrc
}

func (b book) read() int {
	return BTimePerPage * b.bk.pgCnt
}
