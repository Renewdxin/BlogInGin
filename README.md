`runtime.Callers` 和 `runtime.CallersFrames` 是 Go 编程语言标准库中 `runtime` 包提供的两个用于获取调用栈信息的函数。

1. **`runtime.Callers` 函数：**
    - `runtime.Callers` 函数用于获取当前 goroutine 的调用栈信息，即获取一系列程序计数器（PC）值。这些 PC 值可以通过 `runtime.FuncForPC` 函数转换为函数信息，包括函数名、文件名和行号等。具体的函数签名如下：
      ```go
      func Callers(skip int, pc []uintptr) int
      ```
        - `skip` 参数表示要跳过的调用栈帧数，例如，如果希望获取当前函数的调用栈信息，可以将 `skip` 设置为 `1`。
        - `pc` 参数是一个用于接收 PC 值的切片。

2. **`runtime.CallersFrames` 函数：**
    - `runtime.CallersFrames` 函数用于将程序计数器（PC）值转换为调用栈帧信息，提供更丰富的调用者信息。具体的函数签名如下：
      ```go
      func CallersFrames(callers []uintptr) *Frames
      ```
        - `callers` 参数是一个包含 PC 值的切片，通常是通过 `runtime.Callers` 获取的。
        - `Frames` 类型是一个迭代器，提供了每个调用栈帧的详细信息，包括文件名、行号、函数名等。

这两个函数通常用于在调试或记录日志时获取调用栈信息，以便更好地理解程序的执行流程，或者在日志中附加额外的调用者信息。在上述的日志库中，`WithCallersFrames` 方法使用了这两个函数来获取多个调用者的详细信息，并将其保存在 `Logger` 实例的 `callers` 字段中。