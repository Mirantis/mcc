package api

import (
	"fmt"
	"strings"
	"sync"
)

// Hosts is a collection of Hosts
type Hosts []*Host

// First returns the first host
func (hosts *Hosts) First() *Host {
	if len(*hosts) == 0 {
		return nil
	}
	return (*hosts)[0]
}

// Last returns the last host
func (hosts *Hosts) Last() *Host {
	c := len(*hosts) - 1

	if c < 0 {
		return nil
	}

	return (*hosts)[c]
}

// Filter returns a filtered list of Hosts. The filter function should return true for hosts matching the criteria.
func (hosts *Hosts) Filter(filter func(h *Host) bool) Hosts {
	result := make(Hosts, 0, len(*hosts))

	for _, h := range *hosts {
		if filter(h) {
			result = append(result, h)
		}
	}

	return result
}

// Find returns the first matching Host. The finder function should return true for a Host matching the criteria.
func (hosts *Hosts) Find(filter func(h *Host) bool) *Host {
	for _, h := range *hosts {
		if filter(h) {
			return (h)
		}
	}
	return nil
}

// Index returns the index of the first matching Host. The finder function should return true for a Host matching the criteria.
func (hosts *Hosts) Index(filter func(h *Host) bool) int {
	for i, h := range *hosts {
		if filter(h) {
			return (i)
		}
	}
	return -1
}

// IndexAll returns the indexes of the matching Hosts. The finder function should return true for a Host matching the criteria.
func (hosts *Hosts) IndexAll(filter func(h *Host) bool) []int {
	result := make([]int, 0, len(*hosts))
	for i, h := range *hosts {
		if filter(h) {
			result = append(result, i)
		}
	}
	return result
}

// Each runs a function on every Host. The function should return nil or an error. The first encountered error
// will be returned and the process will be halted.
func (hosts *Hosts) Each(filter func(h *Host) error) error {
	for _, h := range *hosts {
		if err := filter(h); err != nil {
			return fmt.Errorf("%s: %s", h, err.Error())
		}
	}
	return nil
}

// ParallelEach runs a function on every Host parallelly. The function should return nil or an error.
// Any errors will be concatenated and returned.
func (hosts *Hosts) ParallelEach(filter func(h *Host) error) error {
	var wg sync.WaitGroup
	var errors []string
	type erritem struct {
		address string
		err     error
	}
	ec := make(chan erritem, 1)

	wg.Add(len(*hosts))

	for _, h := range *hosts {
		go func(h *Host) {
			ec <- erritem{h.String(), filter(h)}
		}(h)
	}

	go func() {
		for e := range ec {
			if e.err != nil {
				errors = append(errors, fmt.Sprintf("%s: %s", e.address, e.err.Error()))
			}
			wg.Done()
		}
	}()

	wg.Wait()

	if len(errors) > 0 {
		return fmt.Errorf("failed on %d hosts:\n - %s", len(errors), strings.Join(errors, "\n - "))
	}

	return nil
}

// Map returns a new slice which is the result of running the map function on each host.
func (hosts *Hosts) Map(filter func(h *Host) interface{}) []interface{} {
	result := make([]interface{}, len(*hosts))
	for i, h := range *hosts {
		result[i] = filter(h)
	}
	return result
}

// MapString returns a new slice which is the result of running the map function on each host
func (hosts *Hosts) MapString(filter func(h *Host) string) []string {
	result := make([]string, len(*hosts))
	for i, h := range *hosts {
		result[i] = filter(h)
	}
	return result
}

// Include returns true if any of the hosts match the filter function criteria.
func (hosts *Hosts) Include(filter func(h *Host) bool) bool {
	for _, h := range *hosts {
		if filter(h) {
			return true
		}
	}
	return false
}

// Count returns the count of hosts matching the filter function criteria.
func (hosts *Hosts) Count(filter func(h *Host) bool) int {
	return len(hosts.IndexAll(filter))
}

func (h *Hosts) Validate() error {
	if h.Count(func(h *Host) bool { return h.Role == "manager" }) < 1 {
		return fmt.Errorf("spec.hosts at least one host with role=manager required")
	}
	return nil
}
