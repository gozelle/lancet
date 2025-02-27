# System

Package system contains some functions about os, runtime, shell command.

<div STYLE="page-break-after: always;"></div>

## Source:

-   [https://github.com/duke-git/lancet/blob/main/system/os.go](https://github.com/duke-git/lancet/blob/main/system/os.go)

<div STYLE="page-break-after: always;"></div>

## Usage:

```go
import (
    "github.com/gozelle/lancet/system"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

-   [IsWindows](#IsWindows)
-   [IsLinux](#IsLinux)
-   [IsMac](#IsMac)
-   [GetOsEnv](#GetOsEnv)
-   [SetOsEnv](#SetOsEnv)
-   [RemoveOsEnv](#RemoveOsEnv)
-   [CompareOsEnv](#CompareOsEnv)
-   [ExecCommand](#ExecCommand)
-   [GetOsBits](#GetOsBits)

<div STYLE="page-break-after: always;"></div>

## Documentation

### <span id="IsWindows">IsWindows</span>

<p>Check if current os is windows.</p>

<b>Signature:</b>

```go
func IsWindows() bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    isOsWindows := system.IsWindows()
    fmt.Println(isOsWindows)
}
```

### <span id="IsLinux">IsLinux</span>

<p>Check if current os is linux.</p>

<b>Signature:</b>

```go
func IsLinux() bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    isOsLinux := system.IsLinux()
    fmt.Println(isOsLinux)
}
```

### <span id="IsMac">IsMac</span>

<p>Check if current os is macos.</p>

<b>Signature:</b>

```go
func IsMac() bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    isOsMac := system.IsMac()
    fmt.Println(isOsMac)
}
```

### <span id="GetOsEnv">GetOsEnv</span>

<p>Gets the value of the environment variable named by the key.</p>

<b>Signature:</b>

```go
func GetOsEnv(key string) string
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    result := system.GetOsEnv("foo")

    fmt.Println(err)
    fmt.Println(result)
    // Output:
    // <nil>
    // abc
}
```

### <span id="SetOsEnv">SetOsEnv</span>

<p>Sets the value of the environment variable named by the key.</p>

<b>Signature:</b>

```go
func SetOsEnv(key, value string) error
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    result := system.GetOsEnv("foo")

    fmt.Println(err)
    fmt.Println(result)
    // Output:
    // <nil>
    // abc
}
```

### <span id="RemoveOsEnv">RemoveOsEnv</span>

<p>Remove a single environment variable.</p>

<b>Signature:</b>

```go
func RemoveOsEnv(key string) error
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    err1 := system.SetOsEnv("foo", "abc")
    result1 := GetOsEnv("foo")

    err2 := system.RemoveOsEnv("foo")
    result2 := GetOsEnv("foo")

    fmt.Println(err1)
    fmt.Println(err2)
    fmt.Println(result1)
    fmt.Println(result2)

    // Output:
    // <nil>
    // <nil>
    // abc
    //
}
```

### <span id="CompareOsEnv">CompareOsEnv</span>

<p>Get env named by the key and compare it with comparedEnv.</p>

<b>Signature:</b>

```go
func CompareOsEnv(key, comparedEnv string) bool
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    err := system.SetOsEnv("foo", "abc")
    if err != nil {
        return
    }

    result := system.CompareOsEnv("foo", "abc")

    fmt.Println(result)

    // Output:
    // true
}
```

### <span id="ExecCommand">ExecCommand</span>

<p>Execute shell command, return the stdout and stderr string of command, and error if error occur. param `command` is a complete command string, like, ls -a (linux), dir(windows), ping 127.0.0.1. In linux, use /bin/bash -c to execute command, In windows, use powershell.exe to execute command.</p>

<b>Signature:</b>

```go
type (
	Option func(*exec.Cmd)
)
func ExecCommand(command string, opts ...Option) (stdout, stderr string, err error)
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    // linux or mac
    stdout, stderr, err := system.ExecCommand("ls")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)
    assert.Equal("", stderr)

    // windows
    stdout, stderr, err = system.ExecCommand("dir")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)

    // error command
    stdout, stderr, err = system.ExecCommand("abc")
    fmt.Println("std out: ", stdout)
    fmt.Println("std err: ", stderr)
    if err != nil {
        fmt.Println(err.Error())
    }
}
```

### <span id="GetOsBits">GetOsBits</span>

<p>Get current os bits, 32bit or 64bit. return 32 or 64</p>

<b>Signature:</b>

```go
func GetOsBits() int
```

<b>Example:</b>

```go
import (
    "fmt"
    "github.com/gozelle/lancet/system"
)

func main() {
    osBit := system.GetOsBits()
    fmt.Println(osBit) // 32 or 64
}
```
