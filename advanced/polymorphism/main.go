package polymorphism

import "fmt"

type journal interface {
	read() int
}

type infrmSrc struct {
	name  string
	pgCnt int
}

func ReadTime() {
	fmt.Println("Total read time for all Journals")
	journals := mkEnv()
	calReadTime(journals)
}

func mkEnv() []journal {

	infrmSrc1 := infrmSrc{"The Go programming", 100}
	b := book{infrmSrc1}
	infrmSrc2 := infrmSrc{"Times of India", 10}
	n := newsPaper{infrmSrc2}
	infrmSrc3 := infrmSrc{"CODE", 25}
	m := magazine{infrmSrc3}

	return []journal{b, n, m}
}

func calReadTime(j []journal) {
	time := 0
	for _, jrnl := range j {
		time += jrnl.read()
	}
	fmt.Printf("Total Time is %d Minutes \n", time)
}
