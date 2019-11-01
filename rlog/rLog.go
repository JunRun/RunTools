/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-11-01 14:26
 */
package rlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	LOG_MAX_BUF = 1024 * 1024
)

const (
	BitDate         = iota << 1                            //日期标记位  2019/01/23
	BitTime                                                //时间标记位  01:23:12
	BitMicroSeconds                                        //微秒级标记位 01:23:12.111222
	BitLongFile                                            // 完整文件名称
	BitShortFile                                           // 最后文件名
	BitLevel                                               // 当前日志级别： 0(Debug), 1(Info), 2(Warn), 3(Error), 4(Panic), 5(Fatal)
	BitStdFlag      = BitDate | BitTime                    //标准头部日志格式
	BitDefault      = BitLevel | BitShortFile | BitStdFlag //默认日志头部格式
)

//日志级别
const (
	LogDebug = iota
	LogInfo
	LogWarn
	LogError
	LogPanic
	LogFatal
)

var levels = []string{
	"[DEBUG]",
	"[INFO]",
	"[WARN]",
	"[ERROR]",
	"[PANIC]",
	"FATAL",
}

type RLog struct {
	mu     sync.RWMutex
	prefix string
	//日志标记位
	flag int
	//日志输出文件描述符
	out io.Writer
	//输出缓冲区
	buf bytes.Buffer
	//当前日志绑定的输出文件
	file *os.File
	//是否打印调试信息
	debugBool bool
	//获取日志文件名和代码上述的runtime.Call 的函数调用层数
	callDepth int
}

func NewRLog(prefix string, out io.Writer, flag int) *RLog {
	rlog := &RLog{
		prefix:    prefix,
		flag:      flag,
		out:       out,
		file:      nil,
		debugBool: false,
		callDepth: 2,
	}
	return rlog
}

/*回收日志*/
func Clean(r *RLog) {
	r.closeFile()
}

//关闭日志绑定的文件
func (r *RLog) closeFile() {
	if r.file != nil {
		_ = r.file.Close()
		r.file = nil
		r.out = os.Stderr
	}
}

//设置头信息
func (r *RLog) formatHeader(buf *bytes.Buffer, t time.Time, file string, line int, level int) {
	if r.prefix != "" {
		buf.WriteByte('<')
		buf.WriteString(r.prefix)
		buf.WriteByte('>')
	}

	//已经设置了时间相关的标识位,那么需要加时间信息在日志头部
	if r.flag&(BitDate|BitTime|BitMicroSeconds) != 0 {
		if r.flag&BitDate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			buf.WriteByte('/') // 2019/
			itoa(buf, int(month), 2)
			buf.WriteByte('/') //2019/01
			itoa(buf, day, 2)
			buf.WriteByte(' ') //2019 /01/01
		}
	}
	//时钟位被标记
	if r.flag&(BitTime|BitMicroSeconds) != 0 {
		hour, min, sec := t.Clock()
		itoa(buf, hour, 2)
		buf.WriteByte(':') // "12:"
		itoa(buf, min, 2)
		buf.WriteByte(':') // "11:15:"
		itoa(buf, sec, 2)  // "11:15:33"
		//微秒被标记
		if r.flag&BitMicroSeconds != 0 {
			buf.WriteByte('.')
			itoa(buf, t.Nanosecond()/1e3, 6) // "11:15:33.123123
		}
		buf.WriteByte(' ')
	}
	// 日志级别位被标记
	if r.flag&BitLevel != 0 {
		buf.WriteString(levels[level])
	}
	//日志当前代码调用文件名名称位被标记
	if r.flag&(BitShortFile|BitLongFile) != 0 {
		//短文件名称
		if r.flag&BitShortFile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					//找到最后一个'/'之后的文件名称  如:/home/go/src/zinx.go 得到 "zinx.go"
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		buf.WriteString(file)
		buf.WriteByte(':')
		itoa(buf, line, -1) //行数
		buf.WriteString(": ")
	}

}

func (r *RLog) OutPut(level int, s string) error {
	now := time.Now()
	var file string
	var line int
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.flag&(BitShortFile|BitLongFile) != 0 {
		r.mu.Unlock()
		var ok bool
		//获取 ,调用文件，代码行数，
		_, file, line, ok = runtime.Caller(r.callDepth)
		if !ok {
			file = "unknown-file"
			line = 0
		}
		r.mu.Lock()
	}
	r.buf.Reset()
	r.formatHeader(&r.buf, now, file, line, level)
	r.buf.WriteString(s)
	//补充回车
	if len(s) > 0 && s[len(s)-1] != '\n' {
		r.buf.WriteByte('\n')
	}
	_, err := r.out.Write(r.buf.Bytes())
	return err

}

//判断日志文件是否存在
func (r *RLog) checkFileExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func mkdirLog(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0775); err != nil {
			if os.IsPermission(err) {
				e = err
			}
		}
	}
	return
}

// ====> Debug <====
func (r *RLog) Debugf(format string, v ...interface{}) {
	if r.debugBool == true {
		return
	}
	_ = r.OutPut(LogDebug, fmt.Sprintf(format, v...))
}

func (r *RLog) Debug(v ...interface{}) {
	if r.debugBool == true {
		return
	}
	_ = r.OutPut(LogDebug, fmt.Sprintln(v...))
}

// ====> Info <====
func (r *RLog) Infof(format string, v ...interface{}) {
	_ = r.OutPut(LogInfo, fmt.Sprintf(format, v...))
}

func (r *RLog) Info(v ...interface{}) {
	_ = r.OutPut(LogInfo, fmt.Sprintln(v...))
}

// ====> Warn <====
func (r *RLog) Warnf(format string, v ...interface{}) {
	_ = r.OutPut(LogWarn, fmt.Sprintf(format, v...))
}

func (r *RLog) Warn(v ...interface{}) {
	_ = r.OutPut(LogWarn, fmt.Sprintln(v...))
}

// ====> Error <====
func (r *RLog) Errorf(format string, v ...interface{}) {
	_ = r.OutPut(LogError, fmt.Sprintf(format, v...))
}

func (r *RLog) Error(v ...interface{}) {
	_ = r.OutPut(LogError, fmt.Sprintln(v...))
}

// ====> Fatal 需要终止程序 <====
func (r *RLog) Fatalf(format string, v ...interface{}) {
	_ = r.OutPut(LogFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (r *RLog) Fatal(v ...interface{}) {
	_ = r.OutPut(LogFatal, fmt.Sprintln(v...))
	os.Exit(1)
}

// ====> Panic  <====
func (r *RLog) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = r.OutPut(LogPanic, s)
	panic(s)
}

func (r *RLog) Panic(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_ = r.OutPut(LogPanic, s)
	panic(s)
}

// ====> Stack  <====
func (r *RLog) Stack(v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, LOG_MAX_BUF)
	n := runtime.Stack(buf, true) //得到当前堆栈信息
	s += string(buf[:n])
	s += "\n"
	_ = r.OutPut(LogError, s)
}

//获取当前日志bitmap标记
func (r *RLog) Flags() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.flag
}

//重新设置日志Flags bitMap 标记位
func (r *RLog) ResetFlags(flag int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.flag = flag
}

//添加flag标记
func (r *RLog) AddFlag(flag int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.flag |= flag
}

//设置日志的 用户自定义前缀字符串
func (r *RLog) SetPrefix(prefix string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.prefix = prefix
}

//设置日志文件输出
func (r *RLog) SetLogFile(fileDir string, fileName string) {
	var file *os.File

	//创建日志文件夹
	_ = mkdirLog(fileDir)

	fullPath := fileDir + "/" + fileName
	if r.checkFileExist(fullPath) {
		//文件存在，打开
		file, _ = os.OpenFile(fullPath, os.O_APPEND|os.O_RDWR, 0644)
	} else {
		//文件不存在，创建
		file, _ = os.OpenFile(fullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	//关闭之前绑定的文件
	r.closeFile()
	r.file = file
	r.out = file
}

//将一个整形转换成一个固定长度的字符串，字符串宽度应该是大于0的
//要确保buffer是有容量空间的
func itoa(buf *bytes.Buffer, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		buf.WriteByte('0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}

	// avoid slicing b to avoid an allocation.
	for bp < len(b) {
		buf.WriteByte(b[bp])
		bp++
	}

}
