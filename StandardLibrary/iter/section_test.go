package __

// Testing replacing use of slices with Go 1.23 iterators

import (
	"fmt"
	"iter"
	"sync"
	"testing"
)

type (
	user struct {
		foo      int
		sections []string
	}
	config struct {
		mu    sync.RWMutex
		users map[string]user
	}
)

// AllSections returns a list of all sections in the config
func (c *config) AllSections() (r []string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, u := range c.users {
		for _, sec := range u.sections {
			r = append(r, sec)
		}
	}
	return r
}

func ensureSections(sections ...string) {
	for _, sec := range sections {
		fmt.Println("code here to ensure set up of [slice] section:", sec)
	}
}

func TestAddSections(t *testing.T) {
	data := config{
		users: map[string]user{
			"andrew": {sections: []string{"sup", "dev"}},
			"barry":  {sections: []string{"man", "dev"}},
			"cathy":  {sections: []string{"hr"}},
		},
	}
	ensureSections(data.AllSections()...)
	ensureSections2(data.AllSections2())
}

func ensureSections2(sections iter.Seq[string]) {
	for sec := range sections {
		fmt.Println("code here to ensure set up of [iter] section:", sec)
	}
}

// AllSections2 returns an iterator for all sections in the config
func (c *config) AllSections2() (r iter.Seq[string]) {
	//c.mu.RLock()
	//defer c.mu.RUnlock()

	return func(yield func(string) bool) {
		for _, u := range c.users {
			for _, sec := range u.sections {
				if !yield(sec) {
					return
				}
			}
		}
	}
}
