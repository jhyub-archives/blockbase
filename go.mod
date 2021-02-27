module blockbase

go 1.14

replace (
    github.com/blbase/serve v0.0.0 => ./serve
)

require (
    github.com/jessevdk/go-flags v1.4.0
    github.com/blbase/serve v0.0.0
)
