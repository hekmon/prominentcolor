package main

import (
	"math"
	"strings"

	"github.com/EdlinOrg/prominentcolor"
)

/*
	ColorSets (sortable)
*/

type ColorSets []ColorSet

func (cs ColorSets) Len() int {
	return len(cs)
}

func (cs ColorSets) Less(i, j int) bool {
	return cs[i].TotalDistance() < cs[j].TotalDistance()
}

func (cs ColorSets) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

/*
	ColorSet
*/

func NewColorSet(set []prominentcolor.ColorItem, params int) ColorSet {
	return ColorSet{
		params: params,
		set:    set,
	}
}

type ColorSet struct {
	params        int
	totalDistance *float64
	set           []prominentcolor.ColorItem
}

func (cs ColorSet) GetParams() int {
	return cs.params
}

func (cs ColorSet) GetParamsString() string {
	list := make([]string, 0, 4)
	// Seed random
	if prominentcolor.IsBitSet(cs.params, prominentcolor.ArgumentSeedRandom) {
		list = append(list, "Random seed")
	} else {
		list = append(list, "Kmeans++")
	}
	// Average mean
	if prominentcolor.IsBitSet(cs.params, prominentcolor.ArgumentAverageMean) {
		list = append(list, "Mean")
	} else {
		list = append(list, "Median")
	}
	// LAB or RGB
	if prominentcolor.IsBitSet(cs.params, prominentcolor.ArgumentLAB) {
		list = append(list, "LAB")
	} else {
		list = append(list, "RGB")
	}
	// Cropping or no cropping
	if prominentcolor.IsBitSet(cs.params, prominentcolor.ArgumentNoCropping) {
		list = append(list, "No cropping")
	} else {
		list = append(list, "Cropping center")
	}
	// build str
	return strings.Join(list, ", ")
}

func (cs ColorSet) GetSet() []prominentcolor.ColorItem {
	return cs.set
}

func (cs *ColorSet) TotalDistance() float64 {
	if cs.totalDistance == nil {
		cs.totalDistance = new(float64)
		var i, j int
		for i = 0; i < len(cs.set)-1; i++ {
			for j = i + 1; j < len(cs.set); j++ {
				(*cs.totalDistance) += distance(cs.set[i], cs.set[j])
			}
		}
	}
	return *cs.totalDistance
}

func distance(a, b prominentcolor.ColorItem) (distance float64) {
	return math.Sqrt(math.Pow(float64(b.Color.R)-float64(a.Color.R), 2) +
		math.Pow(float64(b.Color.G)-float64(a.Color.G), 2) +
		math.Pow(float64(b.Color.B)-float64(a.Color.B), 2))
}
