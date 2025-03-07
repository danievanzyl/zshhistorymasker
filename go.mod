module github.com/danievanzyl/zshhistorymasker

go 1.23.6

require github.com/urfave/cli/v3 v3.0.0-beta1

replace github.com/danievanzyl/zshhistorymasker/pkg/actions => ./pkg/actions

replace github.com/danievanzyl/zshhistorymasker/pkg/sensitive_patterns => ./pkg/sensitive_patterns
