package main

import "fmt"

type Slack struct{}

func (s *Slack) Add (a,b int) error {
	return nil
}

func (s Slack) Add2 (a,b int) error {
	return nil
}

func main() {
	s := Slack{}
	s2 := &Slack{}
	s2.Add2(1,2)
	s2.Add(1,2)//可以变更S2的实际内容
	s.Add2(1,2)//不会变
	s.Add(1,2)
	fmt.Println(s,s2)
}