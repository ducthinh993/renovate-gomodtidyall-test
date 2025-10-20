module github.com/ducthinh993/renovate-gomodtidyall-test/web

go 1.21

require (
	github.com/ducthinh993/renovate-gomodtidyall-test/api v0.0.0
	github.com/ducthinh993/renovate-gomodtidyall-test/sdk v0.0.0
)

replace github.com/ducthinh993/renovate-gomodtidyall-test/api => ../api
replace github.com/ducthinh993/renovate-gomodtidyall-test/sdk => ../sdk