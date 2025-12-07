package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"pz13/internal/work"
)

func main() {
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("Fib(38)")()

		n := 38
		res := work.FibFast(n)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(fmtInt(res)))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
