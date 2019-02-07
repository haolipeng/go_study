package persist

import "log"

//获取itemSaver的input 通道
func ItemSaver() chan interface{} {
	out := make(chan interface{})

	//print item
	go func() {
		itemCount := 0
		for ; ; {
			item := <-out
			log.Printf("Item Saver: Got item "+
				"#%d: %v", itemCount, item)

			itemCount++
		}
	}()

	return out
}
