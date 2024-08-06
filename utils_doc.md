# utils_doc

## file

- Home() (string, error)

获取当前用户目录，例如：在linux中获取 `~` 目录，实际返回的值可能是 `/home/tom`

- DecodeEncoding(r *bufio.Reader) (*transform.Reader, error)

将其他编码格式的文件转换为UTF-8

- DetermineEncoding(r *bufio.Reader) (encoding.Encoding, error)

判断文件的编码格式

## math

- Uint8ToLower(n uint8) uint8

获取大写字母的uint8值对应的小写字母uint8值，例如："D" == 68 and "d" == 100, Uint8ToLower(68) == 100

- Int2Bytes(i int) ([]byte, error)

int转byte切片

- Bytes2Int64(b []byte) (int64, error)

byte切片转int64

## str

- IPStrToHost(host string) ([4]byte, error)

字符串ip地址转为[4]byte数组，例如："127.0.0.1" => [4]byte{127,0,0,1}

- CamelCaseToUnderscore(str string) string

驼峰表示法单词转下划线表示法单词，例如："HelloWorld" => "hello_world"

- String2Int64(s string, intVal *int64) (bool, error)

字符串转int64

- String2Float64(s string, intVal *float64) (bool, error)

字符串转float64

## time

- GetMsTime() int64

获取当前时间的毫秒时间戳

- GetMsTimeByTime(t *time.Time) int64

通过 `*time.Time` 获取对应时间的毫秒时间戳

## ulog

- Error(v ...any) 、 ErrorF(format string, v ...any)

打印错误日志，同时退出程序

- ErrorP(v ...any) 、 ErrorPf(format string, v ...any)

打印错误日志，程序继续执行

- Info(v ...any) 、 InfoF(format string, v ...any)

打印日志

- SetLevel(level int)

设置日志可见等级: 

1. `InfoLevel`同时打印error和info日志
2. `ErrorLevel`只打印error日志
3. `Disabled`不打印日志