/*
* @Author: scottxiong
* @Date:   2019-11-19 15:41:27
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-12-30 21:20:23
 */
package parse

import (
	"flag"
	"github.com/scott-x/TextToJson/defs"
	"github.com/scott-x/gutils/cl"
	"github.com/scott-x/gutils/fs"
	"os"
	"strconv"
	"strings"
)

func Parse() {
	isExam := false
	customed := flag.Bool("e", false, "costomed:default false")
	flag.Parse()
	if *customed {
		// 用户自定义
		isExam = true
	}
	flag, c := getContent("a.txt")
	if !flag {
		cl.BoldGreen.Println("Please create a file named a.txt in your current folder, then input as following")
		cl.BoldCyan.Println("How to use iphone?")
		cl.BoldCyan.Println("1.Getting Familiar with the Buttons")
		cl.BoldCyan.Println("2.Turn on your iPhone if it isn't already on. ")
		cl.BoldCyan.Println("3.Charge your iPhone if necessary.")
		cl.BoldCyan.Println("4.Get to know your iPhone's buttons. ")
		cl.BoldCyan.Println("5.Press the Lock button. ")
		cl.BoldCyan.Println("6.Press the Home button once the Lock screen displays.")
		cl.BoldCyan.Println("7.Type in your passcode using the buttons on the screen.")
		return
	}
	o := getObject(c)
	writeToJson(o, isExam)
}

func getContent(file string) (bool, string) {
	str, _ := os.Getwd()
	flag := fs.IsExist(str + "/" + file)
	if !flag {
		return false, ""
	}
	c, err := fs.ReadFile1(file)
	if err != nil {
		panic(err)
	}
	return true, c
}

func getObject(c string) *defs.Matters {
	matters := &defs.Matters{}
	ms := strings.Split(c, "\n\n")
	for k, m := range ms {
		matter := &defs.Matter{}
		mt := strings.Split(m, "\n")
		(*matter).Question = mt[0]
		(*matter).Id = strconv.Itoa(k + 1)
		(*matter).Answers = append((*matter).Answers, mt[1:]...)
		*matters = append(*matters, *matter)
	}
	return matters
}

func writeToJson(matters *defs.Matters, isExam bool) {
	var data = ""
	if isExam {

		data += "[\n"

		for x, matter := range *matters {
			data += tab(2) + "{\n"
			data += tab(4) + "\"myAnswer\":\"\",\n"
			data += tab(4) + "\"id\":\"" + matter.Id + "\",\n"
			data += tab(4) + "\"title\":\"" + matter.Question + "\",\n"
			data += tab(4) + "\"answer\":[\n"
			for y, v := range matter.Answers {
				if y == len(matter.Answers)-1 {
					data += tab(6) + "\"" + v + "\"" + "\n"
				} else {
					data += tab(6) + "\"" + v + "\"," + "\n"
				}
			}
			data += tab(4) + "]\n"
			if x == len(*matters)-1 {
				data += tab(2) + "}\n"
			} else {
				data += tab(2) + "},\n"
			}
		}

		data += "]\n"

	} else {
		data += "{\n"
		data += tab(2) + "\"subject\": \"\",\n"
		data += tab(2) + "\"answers\":[\n"

		for x, matter := range *matters {
			data += tab(4) + "{\n"
			data += tab(6) + "\"question\":\"" + matter.Question + "\",\n"
			data += tab(6) + "\"answers\":[\n"
			for y, v := range matter.Answers {
				if y == len(matter.Answers)-1 {
					data += tab(8) + "\"" + v + "\"" + "\n"
				} else {
					data += tab(8) + "\"" + v + "\"," + "\n"
				}
			}
			data += tab(6) + "]\n"
			if x == len(*matters)-1 {
				data += tab(4) + "}\n"
			} else {
				data += tab(4) + "},\n"
			}
		}

		data += tab(2) + "]\n"
		data += "}"
	}

	fs.WriteString("./data.json", data)
}

func tab(n int) string {
	return strings.Repeat(" ", n)
}
