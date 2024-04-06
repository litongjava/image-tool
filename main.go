package main

import (
  "flag"
  "fmt"
  "github.com/chai2010/webp"
  "github.com/cloudwego/hertz/pkg/common/hlog"
  "image"
  "image/png"
  "os"
  "strings"
)

func main() {
  // 打开图像文件
  var inputFilePath = flag.String("i", "", "input image")
  var outputFilePath = flag.String("o", "", "output image")
  flag.Parse()
  if inputFilePath == nil || *inputFilePath == "" {
    hlog.Error("please input file")
    return
  }
  hlog.Info("input:", *inputFilePath)
  var file *os.File
  var err error
  file, err = os.Open(*inputFilePath)
  if err != nil {
    hlog.Error(err.Error())
    return
  }
  defer file.Close()
  var img image.Image
  if strings.HasSuffix(*inputFilePath, ".webapp") || strings.HasSuffix(*inputFilePath, ".webp") {
    // 解码 WebP 数据
    img, err = webp.Decode(file)
    if err != nil {
      hlog.Error(err.Error())
      return
    }
  } else {
    // 解码图像
    img, _, err = image.Decode(file)
    if err != nil {
      hlog.Error(err.Error())
      return
    }
  }

  // 创建输出文件
  if outputFilePath == nil || *outputFilePath == "" {
    *outputFilePath = "output.png"
  }
  hlog.Info("output:", *outputFilePath)
  outFile, err := os.Create(*outputFilePath)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer outFile.Close()

  // 将图像编码为PNG格式并写入输出文件
  err = png.Encode(outFile, img)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println("Image converted to PNG format successfully.")
}
