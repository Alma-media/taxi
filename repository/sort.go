package repository

import "github.com/Alma-media/taxi/model"

// ByID implements sort.Interface for []*model.Order based on the ID field
type ByID []*model.Order

func (a ByID) Len() int { return len(a) }

func (a ByID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }
