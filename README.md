![img.txt](https://is1-ssl.mzstatic.com/image/thumb/Purple113/v4/9f/26/b7/9f26b799-ba9d-1fce-226d-9dbf1d089c6f/AppIcon-1x_U007emarketing-0-10-0-85-220.png/434x0w.webp)
## img.txt is an iOS App that OCR images to texts. 
### Use img.txt, easy get text content from image, and You can use iPhone as OCR server, it provides a POST API to receive image and return OCR result directly.
[Available on App Store](https://apps.apple.com/us/app/img-txt/id1662261112?l=zh)

[简体中文说明](/Readme.zh-Hans.md)
# ocrclients
Ocr client demos for img.txt.

*You can use iPhone as OCR server, and then call its api...*
#### curl 
```
Format : curl -F 'file=<filename>' <server url of img.txt app>
Example: curl -F 'file=/user/images/photo1.png' http://192.168.0.2:8080
```
