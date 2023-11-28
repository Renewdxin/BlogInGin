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


关于报错

```shell
gorm.io/driver/postgres 
C:\Program Files\JetBrains\GoLand 2023.2.2\pkg\mod\gorm.io\driver\postgres@v1.2.3\migrator.go:341:38: cannot use column (variable of type Column) as gorm.ColumnType value in argument to append: Column does not implement gorm.ColumnType (missing method AutoIncrement)
```
此时可能是版本运行过低的原因
输入命令
```shell
PS D:\GoWokrs\src\BloginGin> go get -u ./...                                             
...
go: upgraded gorm.io/driver/postgres v1.2.3 => v1.5.4
```
之后就发现好了



页码在分页显示大量数据时非常有用。当数据集很大时，一次性加载整个数据集可能会导致性能问题和用户体验下降。通过分页，你可以将数据分割成小块，每次只加载一页数据，以减轻服务器和客户端的负担。

以下是一些使用页码的主要原因：

1. **减轻服务器负担：** 如果你的应用程序需要处理大量数据，一次性加载所有数据可能会对服务器造成很大的负担。通过分页，服务器可以根据用户请求只返回一页数据，降低了数据传输的负担。

2. **提高用户体验：** 用户不需要等待加载整个数据集，他们可以逐页浏览数据。这提高了页面加载速度，使用户体验更加流畅。

3. **节省带宽：** 通过分页，只有用户请求的数据才会被传输，而不是整个数据集。这可以节省带宽，特别是对于移动设备用户或网络条件较差的用户。

4. **便于导航：** 用户可以通过页码轻松地导航到所需的数据页。这种方式比在长列表中滚动更为直观。

5. **避免数据浪费：** 如果用户仅关注数据集中的一小部分，分页可以避免不必要的数据传输和处理。

在 Web 应用程序中，分页通常是通过 URL 参数来实现的，如 `page` 和 `page_size`。这些参数用于指定用户请求的页码和每页的数据量。在服务端，你可以使用这些参数来计算偏移量，并仅返回所需的数据。




