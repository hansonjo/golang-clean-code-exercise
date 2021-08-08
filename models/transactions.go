package models

type transactions struct {
        data map[int]int
}

// TODO: implement this
// i.e. go test is green
func NewTransactions(d []int) Transactions {
        data := make(map[int]int)

        for i, v := range d {
                data[i] = v
        }

        return &transactions{data: data}
}

// TODO: implement this
// i.e. go test is green
func (t *transactions) Get(idx int) int {
        return t.data[idx]
}

// TODO: implement this method
// i.e. go test is green
func (t *transactions) GetTotal() int {
        salesTotal := 0
        for i := 0; i <= 6; i++ {
                salesTotal += t.Get(i)
        }
        return salesTotal
}

// TODO: implement this method
// i.e. go test is green
func (t *transactions) GetTotalWithinRange(i, j int) int {
        salesTotal := 0
        for i := i; i <= j; i++ {
                salesTotal += t.Get(i)
        }
        return salesTotal
}
