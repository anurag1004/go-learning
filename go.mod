module mymodule

go 1.20

require github.com/anurag0608/myutil2 v0.0.0-20230608145336-98ff9569733e

require (
	example.com/mypackage v0.0.0-00010101000000-000000000000 // indirect
	github.com/anurag0608/myutil v0.0.0-20230608143836-ff52187e33a6 // indirect
)

replace example.com/mypackage => ./mypackage
