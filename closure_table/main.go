package main

import "fmt"

type Category struct {
	ID   int
	name string
}

type CategoryPath struct {
	AnscestorId  int
	DescendantId int
	PathLength   int
}

// Name カテゴリ名を取得するメソッド
func (c *Category) Name() string {
	return c.name
}

func CategoryOfId(id int, data []Category) Category { // O(n)
	var target Category
	for _, category := range data { // O(n)
		if category.ID == id {
			target = category
		}
	}
	return target
}

func (category Category) Ancestors(categoryData []Category, pathData []CategoryPath) []Category {
	anscestors := []Category{}
	for _, path := range pathData {
		if path.DescendantId == category.ID {
			anscestors = append(anscestors, CategoryOfId(path.AnscestorId, categoryData))
		}
	}
	return anscestors
}

// AnscestorIs 指定されたカテゴリが先祖カテゴリかどうかを判定するメソッド
func (category *Category) AnscestorIs(target Category, pathData []CategoryPath) bool { // O(n*log(n))
	for _, path := range pathData { // O(n*log(n))
		if path.DescendantId == category.ID && path.AnscestorId == target.ID {
			return true
		}
	}
	return false
}

// 一般的なcontainsメソッドが存在しないため実装
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (category Category) Descendants(categoryData []Category, pathData []CategoryPath) []Category { // O(n^2 * log(n))
	descendants := []Category{}
	for _, path := range pathData { // O(n*log(n))
		if path.AnscestorId == category.ID {
			descendants = append(descendants, CategoryOfId(path.DescendantId, categoryData)) // O(n)
		}
	}
	return descendants
}

// DescendantIs 指定されたカテゴリが子孫カテゴリかどうかを判定するメソッド
func (category Category) DescendantIs(target Category, pathData []CategoryPath) bool {
	for _, path := range pathData {
		if path.AnscestorId == category.ID && path.DescendantId == target.ID {
			return true
		}
	}
	return false
}

// AllCategories すべてのカテゴリを取得する
func AllCategories() []Category {
	return []Category{
		Category{1, "ファッション"},
		Category{2, "家具・インテリア"},
		Category{29, "レディース"},
		Category{30, "メンズ"},
		Category{38, "収納家具"},
		Category{200, "靴"},
	}
}

// RootOf 指定されたカテゴリのルートカテゴリを取得する
func RootOf(category Category, categoryData []Category, pathData []CategoryPath) Category {
	var root Category
	for _, datum := range pathData {
		anscestor := CategoryOfId(datum.AnscestorId, categoryData)
		if anscestor.DescendantIs(category, pathData) && len(anscestor.Ancestors(categoryData, pathData)) == 1 {
			root = anscestor
		}
	}
	return root
}

// AnscestorsOf 指定されたカテゴリの先祖カテゴリをすべて取得する
func AnscestorsOf(category Category, categoryData []Category, pathData []CategoryPath) []Category {
	return category.Ancestors(categoryData, pathData)
}

// DescendantsOf 指定されたカテゴリの子孫カテゴリをすべて取得する
func DescendantsOf(category *Category, categoryData []Category, pathData []CategoryPath) []Category {
	return category.Descendants(categoryData, pathData)
}

var categories = AllCategories()

var paths = []CategoryPath{
	CategoryPath{1, 1, 0},
	CategoryPath{1, 29, 1},
	CategoryPath{1, 30, 1},
	CategoryPath{1, 200, 2},
	CategoryPath{2, 2, 0},
	CategoryPath{2, 38, 1},
	CategoryPath{29, 29, 0},
	CategoryPath{29, 200, 1},
	CategoryPath{30, 30, 0},
	CategoryPath{38, 38, 0},
	CategoryPath{200, 200, 0},
}

// デバッグ用
func main() {
	fmt.Println(categories[0])
	fmt.Println(categories[0].AnscestorIs(categories[0], paths))
	fmt.Println(categories[5].Ancestors(categories, paths))
	fmt.Println(categories[0].Descendants(categories, paths))
	fmt.Println(RootOf(categories[4], categories, paths))
}
