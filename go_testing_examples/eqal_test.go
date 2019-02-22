package go_testing_examples

import (
	"fmt"
	"testing"
)

func BenchmarkIsEqual(b *testing.B) {
	m1 := &member{}
	m2 := &member{}
	for i := 0; i < b.N; i++ { //use b.N for looping
		fmt.Println(m1.IsEqual(m2))
	}
}

type member struct {
	GroupID  int
	MemberID int
	UserID   int
	Token    string
	IP       string
}

func (m *member) IsEqual(other *member) bool {
	if m.GroupID != other.GroupID {
		return false
	}

	if m.MemberID != other.MemberID {
		return false
	}

	if m.UserID != other.UserID {
		return false
	}

	if m.Token != other.Token {
		return false
	}

	if m.IP != other.IP {
		return false
	}
	return true
}
