package scan

import (
	"InfoFinder/utils"
	"fmt"
)

func OpenFile() {

}

func Target_dirscan(target string, threads int, timeout int, output string, wordlist string) {
	fmt.Println("target_dirscan")
	targetlists := utils.InitTarget(target, wordlist)
	fmt.Println(targetlists)
}

func File_dirscan() {
	fmt.Println("file_dirscan")
}
