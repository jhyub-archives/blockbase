module blockbase

go 1.14

replace (
    github.com/blbase/serve v0.0.0 => ./serve
    github.com/blbase/base v0.0.0 => ./base
)

require (
    github.com/jessevdk/go-flags v1.4.0
    github.com/blbase/serve v0.0.0
    github.com/blbase/base v0.0.0
)
