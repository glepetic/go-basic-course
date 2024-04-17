package main

type num int

func (n num) isEven() bool {
	return n%2 == 0
}
