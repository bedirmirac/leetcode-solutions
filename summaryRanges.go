import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	ranges := []string{}
	first := ""
	second := ""
	nums = append(nums, nums[0])
	for i := 0; i < len(nums); i++ {
		if i+1 >= len(nums) {
			break
		}
		if nums[i+1] == nums[i]+1 && first == "" {
			first = strconv.Itoa(nums[i])
		} else if nums[i+1] != nums[i]+1 && first == "" {
			ranges = append(ranges, strconv.Itoa(nums[i]))
		}
		if first != "" && nums[i+1] != nums[i]+1 {
			second = strconv.Itoa(nums[i])
			if first == second {
				ranges = append(ranges, first)
			} else {
				text := fmt.Sprintf("%v->%v", first, second)
				ranges = append(ranges, text)
			}
			first = ""
		}

	}
	return ranges
}

