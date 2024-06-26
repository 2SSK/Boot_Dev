# SLICES REVIEW

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimensions such as trnasformation matrices, most array programming in Go is done with slices rather than simple arrays.

Slices hole regerernces to an underlying array, and if you assign one slice to another, both reger to the **same** array. If a function takes a slice argument, any changes it makes to the elements of the slice _will be visible to the caller_, analogous to passing a pointer to the underlying array. A Read function can therefor accept a slice argument rather than a pinter and a count; the length within the slice sets an upper limit of how much data to read. Here is the signature of the Read() method of the `File` type in package `os`:

    func (f *File) Read(buf []byte) (n int, err error)
