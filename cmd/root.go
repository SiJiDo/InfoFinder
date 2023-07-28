package cmd

import "InfoFinder/scan"

func Execute(MODE int, target string, threads int, timeout int, output string, jsurlflag bool, jsinfoflag bool, dirbflag bool, wordlist string) {
	if MODE == 0 {
		if dirbflag == true {
			scan.Target_dirscan(target, threads, timeout, output, wordlist)
		}
		if jsurlflag == true {
			scan.Target_jsurlscan()
		}
		if jsinfoflag == true {
			scan.Target_jsinfoflag()
		}
	}
	if MODE == 1 {
		if dirbflag == true {
			scan.File_dirscan()
		}
		if jsurlflag == true {
			scan.File_jsurlscan()
		}
		if jsinfoflag == true {
			scan.File_jsinfoflag()
		}
	}
}
