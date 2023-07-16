# go-kenall-example

郵便番号データ、通称 KEN_ALL でファイル操作を行っているサンプルです。

(新しく追加された utf_all.csv を使っています）

![Go Version](https://img.shields.io/badge/go-1.20-blue.svg)

## Environments

```sh
$ lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 22.04.2 LTS
Release:        22.04
Codename:       jammy


$ go version
go version go1.20.5 linux/amd64
```

## Requirements

### [Task](https://taskfile.dev/)

```sh
$ go install github.com/go-task/task/v3/cmd/task@latest
```

## How to

### Run

```sh
$ task download
task: [download] wget https://www.post.japanpost.jp/zipcode/utf_all.csv
--2023-07-16 17:46:56--  https://www.post.japanpost.jp/zipcode/utf_all.csv
Resolving www.post.japanpost.jp (www.post.japanpost.jp)... 43.253.212.17
Connecting to www.post.japanpost.jp (www.post.japanpost.jp)|43.253.212.17|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 18305986 (17M) [text/csv]
Saving to: ‘utf_all.csv’

utf_all.csv                      100%[========================================================>]  17.46M  1.89MB/s    in 10s     

2023-07-16 17:47:07 (1.70 MB/s) - ‘utf_all.csv’ saved [18305986/18305986]

task: [download] file -i utf_all.csv
utf_all.csv: text/csv; charset=utf-8
task: [download] wc -l utf_all.csv
124270 utf_all.csv
task: [download] ls -lh utf_all.csv | awk '{print $5}'
18M
```

```sh
$ task run
task: [run] go build
task: [run] ./kenall "^東京.*銀座"
東京都中央区銀座
======================================
task: [run] ./kenall "^[^(東京)].*銀座"
北海道夕張郡長沼町銀座
栃木県鹿沼市銀座
埼玉県熊谷市銀座
埼玉県本庄市銀座
富山県富山市新庄銀座
長野県岡谷市銀座
長野県岡谷市東銀座
長野県飯田市銀座
静岡県静岡市清水区銀座
静岡県熱海市銀座町
静岡県伊東市銀座元町
愛知県半田市銀座本町
愛知県刈谷市銀座
滋賀県彦根市銀座町
山口県周南市銀座
山口県周南市みなみ銀座
徳島県徳島市銀座
福岡県北九州市戸畑区銀座
```

## 参考情報

- https://www.post.japanpost.jp/zipcode/download.html
- https://devlights.hatenablog.com/entry/2023/07/14/073000
