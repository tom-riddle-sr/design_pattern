package main

import (
	"fmt"
	"runtime"
)

// ===== Flyweight (Intrinsic State) =====

// TreeType 代表樹的類型（內部狀態 - 可共享）
type TreeType struct {
	name    string
	color   string
	texture string
}

func (t *TreeType) Draw(x, y int) {
	fmt.Printf("🌳 Drawing %s tree (color: %s) at position (%d, %d)\n", t.name, t.color, x, y)
}

// ===== Flyweight Factory =====

// TreeFactory 管理並共享 TreeType 物件
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]*TreeType),
	}
}

// GetTreeType 獲取或創建 TreeType（確保相同類型的樹只創建一次）
func (tf *TreeFactory) GetTreeType(name, color, texture string) *TreeType {
	key := name + "_" + color + "_" + texture

	if treeType, exists := tf.treeTypes[key]; exists {
		fmt.Printf("♻️  Reusing existing TreeType: %s\n", name)
		return treeType
	}

	fmt.Printf("✨ Creating new TreeType: %s\n", name)
	treeType := &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
	tf.treeTypes[key] = treeType
	return treeType
}

func (tf *TreeFactory) GetTotalTreeTypes() int {
	return len(tf.treeTypes)
}

// ===== Context Object (Extrinsic State) =====

// Tree 代表森林中的一棵樹（外部狀態 - 不可共享）
type Tree struct {
	x        int
	y        int
	treeType *TreeType // 參考共享的 TreeType
}

func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}

// ===== Forest (Client) =====

// Forest 管理所有樹木
type Forest struct {
	trees       []*Tree
	treeFactory *TreeFactory
}

func NewForest() *Forest {
	return &Forest{
		trees:       make([]*Tree, 0),
		treeFactory: NewTreeFactory(),
	}
}

func (f *Forest) PlantTree(x, y int, name, color, texture string) {
	treeType := f.treeFactory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)
	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	fmt.Println("\n🌲 Drawing Forest...")
	for _, tree := range f.trees {
		tree.Draw()
	}
}

func (f *Forest) GetStats() {
	fmt.Printf("\n📊 Forest Statistics:\n")
	fmt.Printf("   Total Trees: %d\n", len(f.trees))
	fmt.Printf("   Unique TreeTypes: %d\n", f.treeFactory.GetTotalTreeTypes())
	fmt.Printf("   Memory Saved: ~%.2f%%\n", (1-float64(f.treeFactory.GetTotalTreeTypes())/float64(len(f.trees)))*100)
}

// ===== Memory Comparison =====

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("💾 Memory Allocated: %.2f MB\n", float64(m.Alloc)/1024/1024)
}

func main() {
	fmt.Println("===== Flyweight Pattern Example: Forest Rendering =====\n")

	forest := NewForest()

	fmt.Println("🌱 Planting trees in the forest...\n")

	// 種植大量的樹（只有 3 種樹型，但有很多樹）
	forest.PlantTree(1, 1, "Oak", "Green", "oak_texture.png")
	forest.PlantTree(2, 3, "Pine", "Dark Green", "pine_texture.png")
	forest.PlantTree(5, 5, "Birch", "White", "birch_texture.png")
	forest.PlantTree(10, 10, "Oak", "Green", "oak_texture.png")        // 重用 Oak
	forest.PlantTree(15, 20, "Pine", "Dark Green", "pine_texture.png") // 重用 Pine
	forest.PlantTree(25, 30, "Oak", "Green", "oak_texture.png")        // 重用 Oak
	forest.PlantTree(35, 40, "Birch", "White", "birch_texture.png")    // 重用 Birch
	forest.PlantTree(45, 50, "Pine", "Dark Green", "pine_texture.png") // 重用 Pine
	forest.PlantTree(55, 60, "Oak", "Green", "oak_texture.png")        // 重用 Oak
	forest.PlantTree(65, 70, "Oak", "Green", "oak_texture.png")        // 重用 Oak

	// 繪製森林
	forest.Draw()

	// 顯示統計資料
	forest.GetStats()

	// 顯示記憶體使用
	fmt.Println()
	printMemoryUsage()

	fmt.Println("\n===== Flyweight Pattern Benefits =====")
	fmt.Println("✓ 大幅減少記憶體使用（共享相同的 TreeType）")
	fmt.Println("✓ 只創建 3 個 TreeType 物件，卻能渲染 10 棵樹")
	fmt.Println("✓ 內部狀態（name, color, texture）被共享")
	fmt.Println("✓ 外部狀態（x, y 座標）對每棵樹都是唯一的")
	fmt.Println("✓ 適合需要大量相似物件的場景（遊戲、圖形系統）")

	fmt.Println("\n===== Scaling Example =====")
	fmt.Println("Without Flyweight:")
	fmt.Println("  1,000,000 trees × 1KB per tree = ~1GB memory")
	fmt.Println("\nWith Flyweight:")
	fmt.Println("  3 TreeTypes × 1KB + 1,000,000 positions × 8 bytes = ~8MB memory")
	fmt.Println("  Memory Saved: ~99.2%")
}
