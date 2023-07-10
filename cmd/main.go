package main

import (
	"context"
	"fmt"
	"github.com/abdullayev13/gorm_vs_bun/bun_db"
	"github.com/abdullayev13/gorm_vs_bun/config"
	"github.com/abdullayev13/gorm_vs_bun/gorm_db"
	"github.com/abdullayev13/gorm_vs_bun/model"
	"time"
)

var (
	item = new_item()
	ctx  = context.Background()
)

func main() {
	db := &config.DB{
		Username:  "postgres",
		Password:  "1",
		Name:      "postgres",
		DebugMode: false,
	}
	gorm := gorm_db.New(db)
	bun := bun_db.New(db)
	_, _ = bun, gorm

	//fng := func() { gorm.Create(new_item()) }
	//fnb := func() { bun.NewInsert().Model(new_item()).Exec(ctx) }
	fng := func() { gorm.Find(&model.Item{}) }
	fnb := func() { bun.NewSelect().Model(&model.Item{}).Scan(ctx) }
	for i := 0; i < 160; i++ {
		go fmt.Println("gorm : ", timer(fng))
		go fmt.Println("bun  : ", timer(fnb))
	}

	fmt.Println("gorm : ", timer(fng))
	fmt.Println("bun  : ", timer(fnb))
}

func timer(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

func new_item() *model.Item {
	return &model.Item{Name: "laptop", Price: "650$"}
}
