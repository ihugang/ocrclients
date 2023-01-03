import os
import sys
import glob
import requests
import shutil
import json


def postImage(serverUrl, imageFile):
    # Upload image to Imgur
    print("image file:", imageFile)
    txtFile = os.path.join(os.path.dirname(
        imageFile), os.path.basename(imageFile) + ".txt")
    if os.path.exists(txtFile):
        return

    with open(imageFile, "rb") as f:
        # up = {'file': (("file", f), "multipart/form-data")}
        r = requests.post(serverUrl, files={'file': f})
        r.encoding = 'utf-8'
        print(r.text)

    rr = json.loads(r.text)
    print(rr)

    if rr["success"] == False:
        print("Upload failed:", rr["text"])
        return

    text = rr["text"]

    print("text:", text)

    f = open(txtFile, "w", encoding="utf-8")
    if text != "":
        f.write(text)
    else:
        f.write("Non Text")
    f.close()

    if len(text) == 0:
        return

    score = 0
    if text.__contains__("自由行"):
        score += 10
    if text.__contains__("之旅"):
        score += 10
    if text.__contains__("日游"):
        score += 10
    if text.__contains__("景点"):
        score += 10
    if text.__contains__("美食"):
        score += 10
    if text.__contains__("航班"):
        score += 5
    if text.__contains__("旅游"):
        score += 5
    if text.__contains__("酒店"):
        score += 5
    if text.__contains__("麦当劳"):
        score += 5
    if text.__contains__("文三路"):
        score += 2
    if text.__contains__("浙江社群"):
        score += 2
    if text.__contains__("优惠券"):
        score += 2
    print("Score:", score)

    if score > 0:
        print("Upload success")
        newFile = os.path.join(os.path.dirname(
            imageFile), "tobedeleted", os.path.basename(imageFile))
        if not os.path.exists(os.path.dirname(newFile)):
            os.mkdir(os.path.dirname(newFile))
        shutil.move(imageFile, newFile)


def getFiles(serverUrl, folder):
    # Get files from folder
    for root, dirs, files in os.walk(folder, topdown=False):
        for name in files:
            fname = os.path.join(root, name)
            print(fname)
            file_extension = os.path.splitext(name)[1]

            if file_extension.__contains__("jpg") or file_extension.__contains__("png") or file_extension.__contains__("jpeg"):
                postImage(serverUrl, fname)

        for name in dirs:
            print(os.path.join(root, name))


def main():
    # Get parameters
    if sys.argv.__len__() < 3:
        print("Usage: python main.py <serverUrl> <folder>")
        return

    folder = sys.argv[2]
    serverUrl = sys.argv[1]

    print("Target Folder:", folder)
    print("Server URL:", serverUrl)

    getFiles(serverUrl, folder)


main()
