# データ構造によるパフォーマンスの比較

## データ構造1: Array of Anscestors
各データが親IDを持つようなデータ型はNaive Treeなどと呼ばれ効率が悪い。
各データに先祖のIDをSliceで持たせることで先祖や子孫の検索が早くなると考えた。

## データ構造2: Closure Table
Categoryとは別にCategoryPathというデータ構造を作り、親子関係の全ての組み合わせを格納する。これによってデータ量は増えてしまうが先祖や子孫の検索が早くなると考えた。

## 比較結果
https://mattn.kaoriya.net/software/lang/go/20161019124907.htm を参考に作成。
```
name             old time/op  new time/op  delta
AnscestorIs-8    5.20ns ±11%  6.49ns ±10%  +24.80%  (p=0.000 n=100+98)
DescendantsOf-8   338ns ± 8%   317ns ±17%   -6.23%  (p=0.000 n=98+100)

```

### AnscestorIs
計算量の大まかなオーダーは以下の通りである。
計算量は木の形状などにもよるが、一般的にはArray of AnscestorsにおいてはAnscestorIdsのみを走査すれば良いため、計算量は小さくなっている。
- Array of Anscestors
```Go
func (category *Category) AnscestorIs(target Category, data []Category) bool { // O(log(n))
	if category.ID == target.ID {
		return true
	}
	for _, anscestorId := range category.AnscestorIds { // O(log(n))
		if anscestorId == target.ID {
			return true
		}
	}
	return false
}
```
- Closure Table
```Go
func (category *Category) AnscestorIs(target Category, pathData []CategoryPath) bool { // O(n * log(n))
	for _, path := range pathData { // O(n * log(n))
		if path.DescendantId == category.ID && path.AnscestorId == target.ID {
			return true
		}
	}
	return false
}
```

### DescendantsOf
計算量の大まかなオーダーは以下の通りである。
今回のベンチマークではClosure Tableの方が若干速い結果になっているが、大きな差はないと考えられる。
Array of Anscestorsは子孫を探す時にはループを発生させる必要があり、
Closure Tableでは2つのテーブルのデータをJOINさせる必要がある。
- Array of Anscestors
```Go
func (category Category) Descendants(data []Category) []Category { // O(n * log(n))
	descendants := []Category{category}
	for _, datum := range data { // O(n)
		if contains(datum.AnscestorIds, category.ID) { // O(log(n))
			descendants = append(descendants, datum) // appendの計算量？
		}
	}
	return descendants
}
```
- Closure Table
```Go
func (category Category) Descendants(categoryData []Category, pathData []CategoryPath) []Category { // O(n * log(n))
	descendants := []Category{}
	for _, path := range pathData { // O(n*log(n))
		if path.AnscestorId == category.ID {
			descendants = append(descendants, CategoryOfId(path.DescendantId, categoryData)) // 追加時のみ呼ばれる, O(n), appendの計算量？
		}
	}
	return descendants
}
```
