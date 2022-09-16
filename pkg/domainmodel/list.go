package domainmodel

import "log"

//TODO library用のリポジトリを作る
func PartitionedList[T any](list []*T, chunkSize int) [][]*T {
	if chunkSize <= 0 {
		log.Fatalln("chunkSize must be positive")
	}
	n := len(list)
	var partitionedList [][]*T

	for i := 0; i < n; i+=chunkSize {
		j := i + chunkSize
		if j > n {
			j = n
		}

		partitionedList = append(partitionedList, list[i:j])
	}

	return partitionedList
}
