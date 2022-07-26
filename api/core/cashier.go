package core

var cinst *Cashier
type Cashier struct {
	store map[float32]Notedata
}

type Notedata struct {
	Num int
	Maxnum int
}

func Initstore(n map[float32]Notedata) Cashier {
	if cinst != nil {
		return *cinst
	}

	cinst = &Cashier{
		store: n,
	}
	return *cinst
}

func (c *Cashier) Add(notetype float32, n int) int {
	exceed := 0
	s := c.store[notetype]
	s.Num += n
	
	if (s.Num > c.store[notetype].Maxnum) {	
		exceed = s.Num - c.store[notetype].Maxnum
		s.Num = c.store[notetype].Maxnum
	}

	c.store[notetype] = s
	return exceed
}

func (c *Cashier) Remove(notetype float32, n int) int {
	take := 0
	s := c.store[notetype]
	
	if (s.Num < n) {
		take = s.Num
		s.Num = 0
	}else {
		take = n
		s.Num -= n
	}

	c.store[notetype] = s
	return take
}

func (c *Cashier) Getnum(notetype float32) int {
	return c.store[notetype].Num
}

func (c *Cashier) Getmax(notetype float32) int {
	return c.store[notetype].Maxnum
}