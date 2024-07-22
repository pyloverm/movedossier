package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

type fileAndMd5 struct {
  name string
  md5Hash  string
}

// what i want to do is check in a specific folder  and compare the files in that folder with the files in an other folder so that i can copy new files to the 2nd folder
// i will use the os package to get the files in the folder and then compare them with the files in the 2nd folder
// if the file is not in the 2nd folder i will copy it to the 2nd folder
// i will use the ioutil package to copy the files
var folder2 = "C:/Users/vctfe/OneDrive/Documents/stad ex"
var folder1 = "C:/Users/vctfe/OneDrive/Documents/projet stad"


func ConvertFileTomd5(filePath string) string {
	file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        log.Fatal(err)
    }

		// convert the hash to a string
		var md5HashString = fmt.Sprintf("%x", hash.Sum(nil))

		return  md5HashString
}

func checkFiles(folderPath  string) []fileAndMd5 {
	
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)

	}
	// create array of names of files in the folder

	var filesInFolder []fileAndMd5

	for _, file := range files {
		fileMd5 :=  ConvertFileTomd5(folderPath + "/" + file.Name())

		filesInFolder = append(filesInFolder, fileAndMd5{name: file.Name(), md5Hash: fileMd5})
	}
	return filesInFolder
}

func compareFiles(filesInFolder1 []fileAndMd5, filesInFolder2 []fileAndMd5) []fileAndMd5 {

	var missingFiles []fileAndMd5
	var hit bool = false

	for _, file1 := range filesInFolder1 {
		hit = false
		for _, file2 := range filesInFolder2 {
			if file1.name == file2.name {
				hit = true
				if file1.md5Hash != file2.md5Hash {
					missingFiles = append(missingFiles, file1)
					break
				}
			}
		}
		if !hit {
			missingFiles = append(missingFiles, file1)
		}
	}
	return missingFiles
}



func copyMissingFiles(filesToCopy []fileAndMd5) {
	// copy the files to the 2nd folder
	for _, file := range filesToCopy {
		// copy the file to the 2nd folder
		data, err := os.ReadFile(folder1 + "/" + file.name )

		if err != nil {
			log.Fatal(err)
			return
		}

		// copy data to the 2nd folder

		err = os.WriteFile(folder2 + "/" + file.name, data, 0644)

		if err != nil {
      fmt.Println("Error writing file:", err)
      return
   	}

	fmt.Println("File copied successfully")  //success message will be printed when     
	}
}

func main() {
		// Call the function
		filesInFolder1 := checkFiles(folder1)
		filesInFolder2 := checkFiles(folder2)
		fmt.Println(filesInFolder1)
		filestomove := compareFiles(filesInFolder1, filesInFolder2)
		fmt.Println(filestomove)
		// compare the files in the folder with the files in the 2nd folder
		copyMissingFiles(filestomove)
		// make it wait for the user to press enter
		fmt.Scanln()
		
}
