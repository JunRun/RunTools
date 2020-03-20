/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020-03-20 14:46
 */
package test

import (
	"context"
	"github.com/JunRun/RunTools/rent/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestEnt(t *testing.T) {
	client, err := ent.Open("mysql", "root:123456@tcp(localhost:3306)/entTest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal("failed crating schema ", err)
	}
}
