# go-createfile

このスクリプトは、更新日時を過去にさかのぼって、ファイルを生成するCLIツールです。

## 使い方

```
  go run main.go
```

本日の日付から、カレントパスに、過去10日分のファイルを生成します。

## コマンドのオプション

### count

指定した数分ファイルを生成します。

デフォルト 10

例）過去100日分のファイルを生成する。

```
  go run main.go　-count 100
  go run main.go　-count=100
  go run main.go　-c 100
```

ショートコマンド c


### path

指定したパスの配下にファイルを生成する。

フォルダ生成はしない。生成済みフォルダを指定すること。

デフォルト カレントパス

例）過去100日分のファイルを生成する。

```
  go run main.go　-path ./test
  go run main.go　-path=./test
  go run main.go　-p ./test
```

ショートコマンド p


### term

指定した期間に応じてファイルを生成する

デフォルト 

|期間|引数|説明|
|---|:---:|---|
|毎日|d|daily(一日ずつ過去に遡って作成する)|
|毎月|m|monthly(1ヶ月ずつ過去に遡って作成する)|
|毎年|y|yearly(1年ずつ過去に遡って作成する)|

例）毎月過去に遡ってファイルを生成する。

```
  go run main.go　-term m
  go run main.go　-term=m
  go run main.go　-t m
```

ショートコマンド t

これら、上記コマンドはすべて組み合わせて使用できます。
