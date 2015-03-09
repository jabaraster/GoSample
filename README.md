Go言語開発環境作成メモ

# IntelliJ IDE
* http://qiita.com/kaiinui/items/433eb86c022ffcad0bea#3-6
* http://dev.classmethod.jp/server-side/language/golang-2/

# goファイル以外をバイナリ化してデプロイを簡単にする
* http://blog.satotaichi.info/one-binary-using-go-bindata/

GoSampleフォルダで次のコマンドを打つ.

```
go-bindata -pkg=static -o=static/assets.go html/...
```
