# 简体中文说明
![img.txt](https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/9f/26/b7/9f26b799-ba9d-1fce-226d-9dbf1d089c6f/AppIcon-1x_U007emarketing-0-10-0-85-220.png/434x0w.webp)
## img.txt 是一个iOS应用，它可以轻松的从图片文件提取文字信息，类似于OCR。
#### 除手工从图片文件中解析文字资料外，你还可以把App当成一个服务器，它提供了一个post api，用于其它程序批量的调用去进行OCR识别操作。再也无须为OCR支付昂贵的接口调用费用。
[Available on App Store](https://apps.apple.com/us/app/img-txt/id1662261112?l=zh)
# ocrclients 
*各种语言的调用例子 Ocr client demos for img.txt.*

*You can use iPhone as OCR server, and then call its api...*
#### curl 
```
Format : curl -F 'file=<filename>' <server url of img.txt app>
Example: curl -F 'file=/user/images/photo1.png' http://192.168.0.2:8080
```
