package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("请输入章节名称cname，章数l，节数m和题目数n")

	}
	cname, l, m, n := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	topDir := fmt.Sprintf("%s.%s.%s", l, m, cname)
	err := os.Mkdir(topDir, 0775)
	if err != nil {
		fmt.Println("无法创建顶级目录")
		os.Exit(1)
	}
	intn, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println("无法转换成题目数", err)
		os.Exit(1)
	}
	gofile := `package main

import (
	"fmt"
)

func main() {
	fmt.Println("程序开始")


	
}
	`
	dirs := []string{}
	for i := 0; i < intn; i++ {
		dir := fmt.Sprintf("%s.%s.%d", l, m, i+1)
		dirs = append(dirs, dir)
		dname := fmt.Sprintf("%s/%s", topDir, dir)
		err := os.Mkdir(dname, 0775)
		if err != nil {
			fmt.Println("无法创建目录", dname)
			os.Exit(1)
		}
		fname := fmt.Sprintf("%s/main.go", dname)
		err = ioutil.WriteFile(fname, []byte(gofile), 0755)
		if err != nil {
			fmt.Println("无法创建", fname)
			os.Exit(1)
		}
	}
	md := "#目录\n\n##练习解答\n"
	for _, v := range dirs {
		md += fmt.Sprintf("* [%s](./%s/main.go)\n", v, v)
	}
	fname := fmt.Sprintf("%s/README.md", topDir)
	err = ioutil.WriteFile(fname, []byte(md), 0755)
	if err != nil {
		fmt.Println("无法创建", fname)
		os.Exit(1)
	}
}
