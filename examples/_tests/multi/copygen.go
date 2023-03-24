// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

// Package copygen contains the setup information for copygen generated code.
package copygen

import (
	"github.com/switchupcb/copygen/examples/_tests/multi/complex"
	"github.com/switchupcb/copygen/examples/_tests/multi/external"
)

// Placeholder represents a basic type.
type Placeholder bool

// Collection represents a type that holds collection field types.

// empty represents a struct that contains an empty struct.

// freefloat serves the purpose of checking for free-floating comments.
type freefloat struct {
	A string
}

type Collection struct {
	Arr [16]byte
	S   []string
	M   map[string]bool
	C   chan int
	I   error
	F   func() int
}

type empty struct {
	e struct{}
}

// NoMatchBasic copies a Placeholder to a Placeholder.
func NoMatchBasic(tP Placeholder, fP Placeholder) {
	// Placeholder fields
}

// NoMatchBasicExternal copies a *Placeholder to a external.Placeholder, *external.Placeholder, bool.
func NoMatchBasicExternal(tP external.Placeholder, tP1 *external.Placeholder, tb bool, fP *Placeholder) {
	// external.Placeholder fields

	// *external.Placeholder fields

	// bool fields
}

// NoMatchArraySimple copies a [16]byte to a Collection.
func NoMatchArraySimple(tC Collection, fb [16]byte) {
	// Collection fields
}

// NoMatchSliceSimple copies a []string to a Collection.
func NoMatchSliceSimple(tC Collection, fs []string) {
	// Collection fields
}

// NoMatchMap copies a map[string]bool to a Collection.
func NoMatchMap(tC Collection, fm map[string]bool) {
	// Collection fields
}

// NoMatchChan copies a chan int to a Collection.
func NoMatchChan(tC Collection, fc chan int) {
	// Collection fields
}

// NoMatchInterface copies a error to a Collection.
func NoMatchInterface(tC Collection, fe error) {
	// Collection fields
}

// NoMatchInterfaceAlias copies a interface{func() string; } to a *Collection.
func NoMatchInterfaceAlias(tC *Collection, fi interface{ func() string }) {
	// *Collection fields
}

// NoMatchFunc copies a func() int to a Collection.
func NoMatchFunc(tC Collection, ff func() int) {
	// Collection fields
}

// NoMatchExternal copies a []external.Collection to a []external.Collection.
func NoMatchExternal(te []external.Collection, fe []external.Collection) {
	// []external.Collection fields
}

// Basic copies a bool to a bool.
func Basic(tb bool, fb bool) {
	// bool fields
	tb = fb
}

// BasicPointer copies a bool to a *bool.
func BasicPointer(tb *bool, fb bool) {
	// *bool fields
	tb = &fb
}

// BasicDoublePointer copies a *bool to a **bool.
func BasicDoublePointer(tb **bool, fb *bool) {
	// **bool fields
	tb = &fb
}

// BasicSimple copies a Placeholder to a Placeholder.
func BasicSimple(tP Placeholder, fP Placeholder) {
	// Placeholder fields
	tP = fP
}

// BasicSimplePointer copies a Placeholder to a *Placeholder.
func BasicSimplePointer(tP *Placeholder, fP Placeholder) {
	// *Placeholder fields
	tP = &fP
}

// BasicPointerMulti copies a *Placeholder to a *Placeholder, *Placeholder, string.
func BasicPointerMulti(tP *Placeholder, tP1 *Placeholder, ts string, fP *Placeholder) {
	// *Placeholder fields
	tP = fP

	// *Placeholder fields

	// string fields
}

// BasicExternal copies a *external.Placeholder to a external.Placeholder.
func BasicExternal(tP external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP
}

// BasicExternalMulti copies a *external.Placeholder to a external.Placeholder, *external.Placeholder.
func BasicExternalMulti(tP external.Placeholder, tP1 *external.Placeholder, fP *external.Placeholder) {
	// external.Placeholder fields
	tP = *fP

	// *external.Placeholder fields
	tP1 = fP
}

// Array copies a [16]byte to a [16]byte.
func Array(tb [16]byte, fb [16]byte) {
	// [16]byte fields
	tb = fb
}

// ArraySimple copies a [16]byte to a *Collection.
func ArraySimple(tC *Collection, fb [16]byte) {
	// *Collection fields
	tC.Arr = fb
}

// ArrayExternal copies a [16]external.Placeholder to a [16]external.Placeholder.
func ArrayExternal(te [16]external.Placeholder, fe [16]external.Placeholder) {
	// [16]external.Placeholder fields
	te = fe
}

// ArrayComplex copies a [16]map[byte]string to a *complex.Collection.
func ArrayComplex(tC *complex.Collection, fm [16]map[byte]string) {
	// *complex.Collection fields
	tC.Arr = fm
}

// ArrayExternalComplex copies a [16]map[*external.Collection]string to a *complex.ComplexCollection.
func ArrayExternalComplex(tC *complex.ComplexCollection, fm [16]map[*external.Collection]string) {
	// *complex.ComplexCollection fields
	tC.Arr = fm
}

// Slice copies a []string to a []string.
func Slice(ts []string, fs []string) {
	// []string fields
	ts = fs
}

// SlicePointer copies a []*int to a []*int.
func SlicePointer(ti []*int, fi []*int) {
	// []*int fields
	ti = fi
}

// SliceSimple copies a []string to a *Collection.
func SliceSimple(tC *Collection, fs []string) {
	// *Collection fields
	tC.S = fs
}

// SliceExternal copies a []external.Placeholder to a []external.Placeholder.
func SliceExternal(te []external.Placeholder, fe []external.Placeholder) {
	// []external.Placeholder fields
	te = fe
}

// SliceComplex copies a []map[string][16]int to a *complex.Collection.
func SliceComplex(tC *complex.Collection, fm []map[string][16]int) {
	// *complex.Collection fields
	tC.S = fm
}

// SliceExternalComplex copies a []map[string]func(*external.Collection) string to a *complex.ComplexCollection.
func SliceExternalComplex(tC *complex.ComplexCollection, fm []map[string]func(*external.Collection) string) {
	// *complex.ComplexCollection fields
	tC.S = fm
}

// Map copies a map[string]bool to a map[string]bool.
func Map(tm map[string]bool, fm map[string]bool) {
	// map[string]bool fields
	tm = fm
}

// MapSimple copies a map[string]bool to a *Collection.
func MapSimple(tC *Collection, fm map[string]bool) {
	// *Collection fields
	tC.M = fm
}

// MapExternal copies a map[string]external.Placeholder to a map[string]external.Placeholder.
func MapExternal(tm map[string]external.Placeholder, fm map[string]external.Placeholder) {
	// map[string]external.Placeholder fields
	tm = fm
}

// MapComplex copies a map[string]interface{func() string; } to a *complex.Collection.
func MapComplex(tC *complex.Collection, fm map[string]interface{ func() string }) {
	// *complex.Collection fields
	tC.M = fm
}

// MapExternalComplex copies a map[*external.Collection]external.Placeholder to a *complex.ComplexCollection.
func MapExternalComplex(tC *complex.ComplexCollection, fm map[*external.Collection]external.Placeholder) {
	// *complex.ComplexCollection fields
	tC.M = fm
}

// Chan copies a chan int to a chan int.
func Chan(tc chan int, fc chan int) {
	// chan int fields
	tc = fc
}

// ChanSimple copies a chan int to a *Collection.
func ChanSimple(tC *Collection, fc chan int) {
	// *Collection fields
	tC.C = fc
}

// ChanExternal copies a chan external.Placeholder to a chan external.Placeholder.
func ChanExternal(tc chan external.Placeholder, fc chan external.Placeholder) {
	// chan external.Placeholder fields
	tc = fc
}

// ChanComplex copies a chan *[]int to a *complex.Collection.
func ChanComplex(tC *complex.Collection, fc chan *[]int) {
	// *complex.Collection fields
	tC.C = fc
}

// ChanExternalComplex copies a chan *[]external.Collection to a complex.ComplexCollection.
func ChanExternalComplex(tC complex.ComplexCollection, fc chan *[]external.Collection) {
	// complex.ComplexCollection fields
	tC.C = fc
}

// Interface copies a interface{} to a interface{}.
func Interface(ti interface{}, fi interface{}) {
	// interface{} fields
	ti = fi
}

// InterfaceSimple copies a error to a *Collection.
func InterfaceSimple(tC *Collection, fe error) {
	// *Collection fields
	tC.I = fe
}

// InterfaceExternal copies a error to a *external.Collection.
func InterfaceExternal(tC *external.Collection, fe error) {
	// *external.Collection fields
	tC.I = fe
}

// InterfaceComplex copies a interface{func(rune) *int; } to a *complex.Collection.
func InterfaceComplex(tC *complex.Collection, fi interface{ func(rune) *int }) {
	// *complex.Collection fields
	tC.I = fi
}

// InterfaceExternalComplex copies a interface{func(string) map[*external.Collection]bool; func() (int, byte); } to a complex.ComplexCollection.
func InterfaceExternalComplex(tC complex.ComplexCollection, fi interface {
	func(string) map[*external.Collection]bool
	func() (int, byte)
}) {
	// complex.ComplexCollection fields
	tC.I = fi
}

// Func copies a func() int to a func() int.
func Func(tf func() int, ff func() int) {
	// func() int fields
	tf = ff
}

// FuncSimple copies a func() int to a *Collection.
func FuncSimple(tC *Collection, ff func() int) {
	// *Collection fields
	tC.F = ff
}

// FuncExternal copies a func(external.Placeholder) int to a func(external.Placeholder) int.
func FuncExternal(tf func(external.Placeholder) int, ff func(external.Placeholder) int) {
	// func(external.Placeholder) int fields
	tf = ff
}

// FuncComplex copies a func([]string, uint64) *byte to a *complex.Collection.
func FuncComplex(tC *complex.Collection, ff func([]string, uint64) *byte) {
	// *complex.Collection fields
	tC.F = ff
}

// FuncExternalComplex copies a func(external.Collection) []string to a *complex.ComplexCollection.
func FuncExternalComplex(tC *complex.ComplexCollection, ff func(external.Collection) []string) {
	// *complex.ComplexCollection fields
	tC.F = ff
}

// EmptyStruct copies a struct{} to a empty.
func EmptyStruct(te empty, fs struct{}) {
	// empty fields
	te.e = fs
}

// Struct copies a Collection to a Collection.
func Struct(tC Collection, fC Collection) {
	// Collection fields
	tC = fC
}

// StructExternal copies a external.Collection to a *Collection.
func StructExternal(tC *Collection, fC external.Collection) {
	// *Collection fields
	tC.Arr = fC.Arr
	tC.S = fC.S
	tC.M = fC.M
	tC.C = fC.C
	tC.I = fC.I
	tC.F = fC.F
}
