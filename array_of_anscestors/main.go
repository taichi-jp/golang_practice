package main

import "fmt"

type Category struct {
	ID           int
	ParentID     int
	AnscestorIds []int
	name         string
}

// Name カテゴリ名を取得するメソッド
func (c *Category) Name() string {
	return c.name
}

// func (category *Category) Parent(data []Category) Category {
// 	var parent Category
// 	for _, datum := range data {
// 		if datum.ID == category.ParentID {
// 			parent = datum
// 		}
// 	}
// 	return parent
// }

func CategoryOfId(id int, data []Category) Category { // O(n)
	var target Category
	for _, category := range data { // O(n)
		if category.ID == id {
			target = category
		}
	}
	return target
}

// func (category Category) IsRoot() bool {
// 	return category.ID == category.ParentID
// }

func (category Category) Ancestors(data []Category) []Category { // O(n^2)
	parents := []Category{category}
	for _, anscestorId := range category.AnscestorIds { // O(n)
		anscestor := CategoryOfId(anscestorId, data) // O(n)
		parents = append(parents, anscestor)
	}
	return parents
}

// AnscestorIs 指定されたカテゴリが先祖カテゴリかどうかを判定するメソッド
func (category *Category) AnscestorIs(target Category, data []Category) bool { // O(n^3)
	parents := category.Ancestors(data) // O(n^2)
	for _, ancestor := range parents {  // O(n)
		if ancestor.ID == target.ID {
			return true
		}
	}
	return false
}

// func (category Category) Children(data []Category) []Category {
// 	children := []Category{}
// 	for _, datum := range data {
// 		if datum.ParentID == category.ID {
// 			if datum.ID == category.ID {
// 				continue
// 			}
// 			children = append(children, datum)
// 		}
// 	}
// 	return children
// }

// 一般的なcontainsメソッドが存在しないため実装
func contains(s []int, e int) bool { // O(n)
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (category Category) Descendants(data []Category) []Category { // O(n^2) + append
	descendants := []Category{category}
	// Naive Tree の場合
	// offSpring := category.Children(data)
	// for len(offSpring) != 0 {
	// 	descendants = append(descendants, offSpring...)
	// 	nextOffSpring := []Category{}
	// 	for _, child := range offSpring {
	// 		nextOffSpring = append(nextOffSpring, child.Children(data)...)
	// 	}
	// 	offSpring = nextOffSpring
	// }
	for _, datum := range data { // O(n)
		if contains(datum.AnscestorIds, category.ID) { // O(n)
			descendants = append(descendants, datum) // appendの計算量？
		}
	}
	return descendants
}

// DescendantIs 指定されたカテゴリが子孫カテゴリかどうかを判定するメソッド
func (category Category) DescendantIs(target Category, data []Category) bool {
	descendants := category.Descendants(data)
	for _, descendant := range descendants {
		if descendant.ID == target.ID {
			return true
		}
	}
	return false
}

// AllCategories すべてのカテゴリを取得する
func AllCategories() []Category {
	return []Category{
		Category{1, 1, []int{}, "ファッション"},
		Category{2, 2, []int{}, "家具・インテリア"},
		Category{29, 1, []int{1}, "レディース"},
		Category{30, 1, []int{1}, "メンズ"},
		Category{38, 2, []int{2}, "収納家具"},
		Category{200, 29, []int{1, 29}, "靴"},
	}
}

// RootOf 指定されたカテゴリのルートカテゴリを取得する
func RootOf(category Category, data []Category) Category {
	var root Category
	for _, datum := range data {
		if len(datum.AnscestorIds) == 0 && contains(category.AnscestorIds, datum.ID) {
			root = datum
		}
	}
	return root
}

// AnscestorsOf 指定されたカテゴリの先祖カテゴリをすべて取得する
func AnscestorsOf(category Category, data []Category) []Category {
	return category.Ancestors(data)
}

// DescendantsOf 指定されたカテゴリの子孫カテゴリをすべて取得する
func DescendantsOf(category *Category, data []Category) []Category {
	return category.Descendants(data)
}

var categories = AllCategories()

// デバッグ用
func main() {
	fmt.Println(categories[0])
	fmt.Println(categories[0].AnscestorIs(categories[0], categories))
	fmt.Println(categories[0].Descendants(categories))
	fmt.Println(RootOf(categories[5], categories))
}
