# kmp-go

* Knuth–Morris–Pratt illustrated (poorly) in Golang
* Implementing [Knuth-Morris-Pratt Illustrated](https://www.cambridge.org/core/journals/journal-of-functional-programming/article/knuthmorrispratt-illustrated/8EFA77D663D585B68630E372BCE1EBA4).

## Files and order to read 

* Naive string search works row-by-row -> `horizontal_naive.go`.
* Going column-by-column yields a new algorithm, but it is still not linear time -> `vertical_set.go` further improved by `vertical_list`.go. 
* MP takes advantage of the underlying structure of columns, representing them as a cyclic graph. This insight yields a linear-time algorithm -> `morris_pratt.go`. 
* KMP refines this algorithm further, skipping over columns that are guaranteed to fail on a mismatched character -> `knuth_morris_pratt.go`.
