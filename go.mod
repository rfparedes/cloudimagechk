module cloudimagechk

go 1.14

replace github.com/rfparedes/cloudimagechk/utils => ./utils

replace github.com/rfparedes/cloudimagechk/providers/ => ./providers/

replace github.com/rfparedes/cloudimagechk/providers/aws => ./providers/

require (
	github.com/rfparedes/cloudimagechk/providers/aws v0.0.0-00010101000000-000000000000
	github.com/rfparedes/cloudimagechk/utils v0.0.0-00010101000000-000000000000
)
