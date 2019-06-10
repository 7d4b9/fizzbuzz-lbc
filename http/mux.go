package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"sort"
	"sync"

	log "github.com/sirupsen/logrus"
)

type call struct {
	Pattern string
	Count   int
}

type calls struct {
	all   []call
	index map[string]int
	mutex sync.Mutex
}

// By is the type of a "less" function that defines the ordering of its call arguments.
type By func(c1, c2 *call) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(calls []call) {
	ps := &callSorter{
		calls: calls,
		by:    by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// callSorter joins a By function and a slice of Planets to be sorted.
type callSorter struct {
	calls []call
	by    func(c1, c2 *call) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *callSorter) Len() int {
	return len(s.calls)
}

// Swap is part of sort.Interface.
func (s *callSorter) Swap(i, j int) {
	s.calls[i].Count, s.calls[j].Count = s.calls[j].Count, s.calls[i].Count
	s.calls[i].Pattern, s.calls[j].Pattern = s.calls[j].Pattern, s.calls[i].Pattern
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *callSorter) Less(i, j int) bool {
	return s.by(&s.calls[i], &s.calls[j])
}

// ServeMux contains an HTTP request multiplexer and a call struct for stats
type ServeMux struct {
	*http.ServeMux
	calls calls
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux.ServeMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		requestFingerPrint := r.URL.RequestURI()
		if r.Body != nil {
			rawBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.WithError(err).Error("missing request body")
			}
			if err := r.Body.Close(); err != nil {
				log.WithError(err).Error("failure closing request body")
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(rawBody))
			requestFingerPrint = path.Join(r.URL.RequestURI(), string(rawBody))
		}
		var wg sync.WaitGroup
		defer wg.Wait()
		wg.Add(1)
		go func() {
			defer wg.Done()
			mux.calls.mutex.Lock()
			defer mux.calls.mutex.Unlock()
			if _, ok := mux.calls.index[requestFingerPrint]; !ok {
				mux.calls.all = append(mux.calls.all, call{Pattern: requestFingerPrint})
				mux.calls.index[requestFingerPrint] = len(mux.calls.all) - 1
			}
			mux.calls.all[mux.calls.index[requestFingerPrint]].Count++
		}()
		handler(w, r)
	})
}

// NewServeMux allocates and returns a new ServeMux.
func NewServeMux() *ServeMux {
	mux := &ServeMux{
		http.NewServeMux(),
		calls{
			index: make(map[string]int),
		},
	}
	mux.HandleFunc("/statistics", mux.Statistics)
	return mux
}

// Criteria returns true if the  parameters are giveen in ascening order
func callCountsGreaterFirstOrderCriteria(c1, c2 *call) bool {
	return c1.Count > c2.Count
}

// Statistics can be used to register an endpoint which deliver the mux stats
func (mux *ServeMux) Statistics(w http.ResponseWriter, r *http.Request) {
	switch meth := r.Method; meth {
	case "GET":
		mux.calls.mutex.Lock()
		defer mux.calls.mutex.Unlock()
		// Order the calls counter in an decreasing count order
		By(callCountsGreaterFirstOrderCriteria).Sort(mux.calls.all)
		// update the indexe
		for i := range mux.calls.all {
			mux.calls.index[mux.calls.all[i].Pattern] = i
		}
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(mux.calls.all[0]); err != nil {
			log.WithError(err).Error("no json encoded payload")
		}
	default:
		MethodNotAllowed(w, meth)
	}
}
