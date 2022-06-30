package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*图片后面加个空
func fileHash(data []byte) string {
	m := md5.New()
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	imga, _ := ioutil.ReadFile("aa.png")
	imgb := append(imga, ' ')
	ioutil.WriteFile("aa.png", imgb, 0666)

	imgamd5 := fileHash(imga)

	imgb, _ = ioutil.ReadFile("aa.png")
	imgbmd5 := fileHash(imgb)

	fmt.Println("imga:", imgamd5)
	fmt.Println("imgb:", imgbmd5)
}
*/

/*遍历目录给图片后面加个空*/
func main() {
	dir, _ := os.Getwd()
	listFiles(dir)
}

//遍历目录
func listFiles(dirName string) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfos {
		fileName := dirName + "/" + file.Name()

		if !file.IsDir() && (strings.HasSuffix(fileName, "jpg") || strings.HasSuffix(fileName, "jpeg") || strings.HasSuffix(fileName, "png") || strings.HasSuffix(fileName, "gif")) {
			fmt.Println(fileName)
			imgAddSpace(fileName)
		}

		if file.IsDir() {
			listFiles(fileName)
		}
	}
}

func imgAddSpace(fileName string) {
	source := fileName
	imga, _ := ioutil.ReadFile(source)
	imgb := append(imga, ' ')
	ioutil.WriteFile(source, imgb, 0666)
}

/*遍历目录给图片左上角打白点
func main() {
	dir, _ := os.Getwd()
	listFiles(dir)
}

//遍历目录
func listFiles(dirName string) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfos {
		fileName := dirName + "/" + file.Name()

		if !file.IsDir() && (strings.HasSuffix(fileName, "jpg") || strings.HasSuffix(fileName, "jpeg") || strings.HasSuffix(fileName, "png") || strings.HasSuffix(fileName, "gif")) {
			fmt.Println(fileName)
			imgSetRGBA(fileName)
		}

		if file.IsDir() {
			listFiles(fileName)
		}
	}
}

//图片打点
func imgSetRGBA(fileName string) {
	source := fileName //输入输出图片

	ff, _ := ioutil.ReadFile(source) //读取文件
	bbb := bytes.NewBuffer(ff)

	m, _, _ := image.Decode(bbb)
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()

	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			r_uint8 := uint8(r >> 8) //转换为 255 值
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)

			if i == 0 && j == 0 {
				newRgba.SetRGBA(0, 0, color.RGBA{255, 255, 255, 255})
			} else {
				newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8})
			}
		}
	}

	f, _ := os.Create(source)
	defer f.Close()
	encode(source, f, newRgba)
}

//图片编码 写入
func encode(inputName string, file *os.File, rgba *image.RGBA) {
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		jpeg.Encode(file, rgba, nil)
	} else if strings.HasSuffix(inputName, "png") {
		png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		gif.Encode(file, rgba, nil)
	} else {
		fmt.Errorf("不支持的图片格式")
	}
}
*/
