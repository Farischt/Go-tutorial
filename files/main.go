package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)

	if !strings.Contains(line, old) && !strings.Contains(line, oldLower) {
		return false, line, 0
	}

	occ += strings.Count(line, old) + strings.Count(line, oldLower)
	res = strings.Replace(line, old, new, -1)
	res = strings.Replace(res, oldLower, new, -1)

	return true, res, occ
}

func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	var lineIdx = 1
	var scanner = bufio.NewScanner(srcFile)
	var writer = bufio.NewWriter(dstFile)
	defer writer.Flush()

	for scanner.Scan() {
		found, res, o := ProcessLine(scanner.Text(), old, new)

		if found {
			occ += o
			lines = append(lines, lineIdx)
		}
		fmt.Fprint(writer, res+"\n")
		lineIdx++
	}

	return occ, lines, nil
}

func Main() {
	var old = "Go"
	var new = "Python"

	occ, lines, erro := FindReplaceFile("files/old.txt", "files/new.txt", old, new)

	if erro != nil {
		fmt.Printf("Error while finding and replacing file : %e", erro)
	}

	fmt.Println("======= SUMMARY =======")
	fmt.Printf("Number of occurences of %s : %d\n", old, occ)
	fmt.Printf("Number of lines: %d\n", len(lines))
	fmt.Printf("Lines: %v\n", lines)
	fmt.Println("========= END =========")
	fmt.Printf("You can view the replacement in file %s\n", "files/new.txt")
}
