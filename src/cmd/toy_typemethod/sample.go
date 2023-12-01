package main

type Sample struct {
	Field1 int
}

func (s *Sample) Method1() {
	// do nothing
}

func (s *Sample) Sample() {
	// do nothing
}

// 구조체 필드와 메서드 이름이 겹치면 컴파일 에러가 발생한다.
//func (s *Sample) Field1() {
//	// do nothing
//}
