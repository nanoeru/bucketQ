package main

import (
	"./arrayQ"
	"flag"
	"fmt"
	"log"
	"strconv"
)

/*
7リットルと3リットルの容器から5リットルの水量を作る問題に おいて,深さ8(固定)で幅優先探索を行って解を見つけよ.すでに調べた状態を省く方式を用いよ.

次の状態は
1.容器の水を捨てて空にする
2.容器Aの水を容器Bが満杯になるまで注ぐ
3.容器が満杯になるまで水を注ぐ
容器の数が2つの場合は6通り

n個
1.	n通り
2.	n*(n-1)通り
3.	n通り

Golangのinterfaceで各メソッドを定義して、上記の動作をさせて一般化する。
1~3の動作を容器Aから容器Bへ入れるという1動作だけのみで実現

最大公約数がキーワードとなる
3と7ならば、1かな
すると、1ずつ全てできるということ?!
*/

var (
	aimValue  int
	depth     int
	firstFlag bool
)

func init() {
	flag.IntVar(&aimValue, "v", 5, "aim value")
	flag.IntVar(&depth, "d", 8, "depth")
	flag.BoolVar(&firstFlag, "f", true, "first flag")
	flag.Parse()
}

func main() {
	args := flag.Args()
	bucketCaps := make([]int, len(args))
	for i, v := range args {
		tmp, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(fmt.Sprintf("%dth arg:%s", i+1, err))
		}
		bucketCaps[i] = tmp
	}

	bucketQue := arrayQ.NewQueue(100)
	c := NewBlankContainers(bucketCaps...)
	//	sumMap := make(map[int]BucketQ)
	size := len(c)

	//	重複判定マップ
	visitedMap := make(map[string]bool)

	bucketQue.Enqueue(NewBucketQ(0, 0, "", c))
	//	幅優先探索
	for !bucketQue.IsEmpty() {
		bucketQInterface, _ := bucketQue.Dequeue()
		bucketQ := bucketQInterface.(BucketQ)

		//	目標値到達判定
		if bucketQ.HasValue(aimValue) {
			fmt.Println(bucketQ)
			//	初回到達で終了
			if firstFlag {
				break
			}
			fmt.Println("--------------------------------")
		}
		//	設定深さ到達
		if bucketQ.Step > depth {
			break
		}

		//	次の状態生成
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				//	同一容器は選択されない
				if i == j {
					continue
				}

				//	値を変更する可能性があるのでコピー生成
				containers := bucketQ.Containers.Copy()
				fromContainer := containers[i]
				toContainer := containers[j]
				//	ログ出力用の情報保存
				preFromContainerString := fromContainer.String()
				preToContainerString := toContainer.String()
				//	変化あり
				if fromContainer.Pour(toContainer) {
					//	重複判定キー
					str := containers.String()
					//	訪問済み
					if _, ok := visitedMap[str]; ok {
						continue
					}
					//	マップ記録
					visitedMap[str] = true
					sum := containers.SumValue()
					//					sumMap[sum] = bucketQ

					//	キューへ
					log := fmt.Sprintf("%s\n%16s ===> %16s (%16s, %16s)\n%s", bucketQ.Log, preFromContainerString, preToContainerString, fromContainer, toContainer, containers)
					bucketQue.Enqueue(NewBucketQ(bucketQ.Step+1, sum, log, containers))
				}
			}
		}
	}

	//	演算可能水量一覧
	//	return
	//	intSlice := make([]int, len(sumMap))
	//	index := 0
	//	for k, _ := range sumMap {
	//		intSlice[index] = k
	//		index++
	//		fmt.Printf("%d ", k)
	//	}
	//	sort.Ints(intSlice)
	//	fmt.Println(intSlice)
	//	for i, v := range intSlice {
	//		fmt.Println(v)
	//	}
}
