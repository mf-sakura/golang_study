# map

## mapに対するfor
mapに対して、forを回すとkey, valueが取得できます。
ただし、順序は保証されません。
同じバイナリに対して実行毎に順序が変わります。
## mapのコピー
mapはSlice同様、参照型です。
`=`でコピーすると、同じポインタを参照します。


## 課題
学生名の表示を学籍番号昇順で並び変わる様にしてください。
ヒント: SliceのSortを使うと楽です。
https://golang.org/pkg/sort/#Slice
