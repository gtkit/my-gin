// @Author xiaozhaofu 2023/7/27 19:27:00
package model

type Meta struct {
	Page int `json:"page,omitempty"`
	// From        int   `json:"from,omitempty"`
	// PerPage     int   `json:"per_page,omitempty"`
	Total int64 `json:"total,omitempty"`
}
