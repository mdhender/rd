// Copyright (c) 2023 Michael D Henderson
// Copyright (c) 2018 Shivam Mamgain
// SPDX-License-Identifier: MIT
//

package rd

type debugStack []*DebugTree

func (ds debugStack) isEmpty() bool {
	return len(ds) == 0
}

func (ds debugStack) peek() *DebugTree {
	return ds[len(ds)-1]
}

func (ds *debugStack) pop() *DebugTree {
	l := len(*ds)
	dt := (*ds)[l-1]
	*ds = (*ds)[:l-1]
	return dt
}

func (ds *debugStack) push(dt *DebugTree) {
	*ds = append(*ds, dt)
}
