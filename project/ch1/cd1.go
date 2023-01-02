package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	//ExecCommand("go", "run", "hello.go")
	ExecCommand("touch", "../readme11.md")
	//ExecCommand("cd", "/home/")
}
func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	return out.String(), err
}
func ExecCommand(name string, args ...string) {
	cmd := exec.Command(name, args...) // 拼接参数与命令
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	var err error
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		log.Println(err)
	}
	fmt.Print(stdout.String())
	fmt.Print(stderr.String())
}
