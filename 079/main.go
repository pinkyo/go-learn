package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// fmt.Println(generateSubset([]int{1, 2, 3}, 2))
	// arr := []int{4, 2, 4, 5, 6, 7, 8, 1, 3, 9, 0}
	// bubbleSort(arr)
	// fmt.Println(arr)
	// readFile("/Users/yinkainian/Codes/codes/go-learn/079/main.go")
	readFile("/Users/yinkainian/Codes/codes/go-learn/079/main.go")
	crawler("http://www.baidu.com")
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range list {
		fmt.Println(v)
	}
}

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	powerSetSize := 1 << uint(len(nums))
	return generateSubsets(nums, powerSetSize)
}

func generateSubsets(nums []int, powerSetSize int) [][]int {
	var result [][]int
	for i := 0; i < powerSetSize; i++ {
		cur := generateSubset(nums, i)
		result = append(result, cur)
	}
	return result
}

func generateSubset(nums []int, i int) []int {
	var cur []int
	for j := 0; j < len(nums); j++ {
		if i&(1<<uint(j)) != 0 {
			cur = append(cur, nums[j])
		}
	}
	return cur
}

func max(values []int) int {
	if len(values) == 0 {
		return -1
	}
	// get max of values
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

// min function
func min(values []int) int {
	if len(values) == 0 {
		return -1
	}
	// get min of values
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// bubble sort
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// read a file and print to console
func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 写一个爬虫，爬取www.baidu.com里面的所有的链接
func crawler(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting url:", err)
		return
	}
	defer resp.Body.Close()
	// bodyBs, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyBs))
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error parsing url:", err)
		return
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "http") {
			fmt.Println(href)
		}
	})
}

// // 使用aws s3上传一个文件
// func s3Upload(fileName string, bucketName string, objectName string) {
// 	sess := session.Must(session.NewSession())
// 	svc := s3.New(sess)
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	fileInfo, _ := file.Stat()
// 	var size int64 = fileInfo.Size()
// 	fileBytes := make([]byte, size)
// 	file.Read(fileBytes)
// 	_, err = svc.PutObject(&s3.PutObjectInput{
// 		Bucket:        aws.String(bucketName),
// 		Key:           aws.String(objectName),
// 		Body:          bytes.NewReader(fileBytes),
// 		ContentLength: aws.Int64(size),
// 	})
// 	if err != nil {
// 		fmt.Println("Error uploading file:", err)
// 		return
// 	}
// 	fmt.Println("File uploaded successfully")
// }
