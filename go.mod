module blockbase

go 1.14

require (
    github.com/jessevdk/go-flags v1.4.0
    github.com/blbase/serve v0.0.0
    github.com/blbase/base v0.0.0
)

replace (
    github.com/blbase/serve v0.0.0 => ./serve
    github.com/blbase/base v0.0.0 => ./base
)
