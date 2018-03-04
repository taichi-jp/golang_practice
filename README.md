# データ構造によるパフォーマンスの比較

## データ構造1: Array of Anscestors
各データが親IDを持つようなデータ型はNaive Treeなどと呼ばれ効率が悪い。
各データに先祖のIDをSliceで持たせることで先祖や子孫の検索が早くなると考えた。

## データ構造2: Closure Table
Categoryとは別にCategoryPathというデータ構造を作り、親子関係の全ての組み合わせを格納する。データ量は増えてしまうが、データ構造1と比べて先祖や子孫の検索が早くなると考えた。

## 比較結果
https://mattn.kaoriya.net/software/lang/go/20161019124907.htm を参考に作成。
```
name             old time/op  new time/op  delta
AnscestorIs-8     811ns ± 1%    49ns ± 1%  -93.96%  (p=0.000 n=10+10)
DescendantsOf-8   848ns ± 3%   767ns ± 2%   -9.56%  (p=0.000 n=10+10)

```

### AnscestorIs
計算量の大まかなオーダーは以下の通りである。
Closure Tableの方の計算量は木の形状などにもよるが、一般的にはArray of Anscestorsと比較して大きく計算量が小さくなっている。
- Array of Anscestors
```Go
func (category *Category) AnscestorIs(target Category, data []Category) bool { // O(n^3)
	parents := category.Ancestors(data) // O(n^2)
	for _, ancestor := range parents {  // O(n)
		if ancestor.ID == target.ID {
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
appendの操作の計算量は調査できていないが、DescendantsOfにおいてはあまり計算量を減らすことができなかった。
- Array of Anscestors
```Go
func (category Category) Descendants(data []Category) []Category { // O(n^2) + append
	descendants := []Category{category}
	for _, datum := range data { // O(n)
		if contains(datum.AnscestorIds, category.ID) { // O(n)
			descendants = append(descendants, datum) // appendの計算量？
		}
	}
	return descendants
}
```
- Closure Table
```Go
func (category Category) Descendants(categoryData []Category, pathData []CategoryPath) []Category { // O(n^2 * log(n))
	descendants := []Category{}
	for _, path := range pathData { // O(n*log(n))
		if path.AnscestorId == category.ID {
			descendants = append(descendants, CategoryOfId(path.DescendantId, categoryData)) // O(n)
		}
	}
	return descendants
}
```
