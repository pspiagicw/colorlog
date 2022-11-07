# `colorlog`

Simple fancy logging package for Go. 
Built for usage in simple personal projects

## This
- Is a simple logging built upon basic Golang packages.
- Prints logging statements with a bit of color and zing.

## This is not
- Performant logging library for performant applications.
- To be used for critical applications where logging is a debugging/auditing requirement.

## Usage
This library simply has 8 functions.
Which follow the familiar `Printf` format.

- LogSuccess: Prints green good stuff
- LogError: Prints red bad stuff
- LogInfo: Prints neutral info stuff
- LogFatal: Prints the given message in red and exits! Use only for very fatal errors.

`LogSuccess`, `LogError` and `LogInfo` has their own `Fancy` counterpart, which prints the same with basic ASCII symbols.

```go
colorlog.LogSuccess("Everything is %s", "right")
colorlog.LogError("Everything is %s", "bad")
colorlog.LogInfo("Everything is %s", "okay")
colorlog.LogFatal("Everything is %s", "fucked")

colorlog.LogSuccessFancy("Everything is %s", "right")
colorlog.LogErrorFancy("Everything is %s", "bad")
colorlog.LogInfoFancy("Everything is %s", "okay")
```




