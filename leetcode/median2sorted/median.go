package median

func Median(nums1 []int, nums2 []int) float64 {

    var left, right int
    total := len(nums1) + len(nums2)
    for i := 0; i <= total/2; i++ {
        if len(nums2) == 0 {
            left = right
            right = nums1[0]
            nums1 = nums1[1:]
            continue
        }
        if len(nums1) == 0 {
            left = right
            right = nums2[0]
            nums2 = nums2[1:]
            continue
        }
        if nums1[0] < nums2[0] {
            left = right
            right = nums1[0]
            nums1 = nums1[1:]
            continue
        }
        left = right
        right = nums2[0]
        nums2 = nums2[1:]

    }
    if total%2 == 0 {
        return (float64(left) + float64(right)) / 2
    }
    return float64(right)
}

func Median2(nums1 []int, nums2 []int) float64 {

    if len(nums1) == 0 {
        return median(nums2)
    }
    if len(nums2) == 0 {
        return median(nums1)
    }
    var left, right int
    total := len(nums1) + len(nums2)
    for i := 0; i <= total/2; i++ {
        if len(nums2) == 0 {
            left = right
            right = nums1[0]
            nums1 = nums1[1:]
            continue
        }
        if len(nums1) == 0 {
            left = right
            right = nums2[0]
            nums2 = nums2[1:]
            continue
        }
        if nums1[0] < nums2[0] {
            left = right
            right = nums1[0]
            nums1 = nums1[1:]
            continue
        }
        left = right
        right = nums2[0]
        nums2 = nums2[1:]

    }
    if total%2 == 0 {
        return (float64(left) + float64(right)) / 2
    }
    return float64(right)
}

func median(sorted []int) float64 {
    length := len(sorted)
    if length%2 == 0 {
        return (float64(sorted[length/2]) + float64(sorted[length/2-1])) / 2
    }
    return float64(sorted[len(sorted)/2])
}

func Median3(nums1 []int, nums2 []int) float64 {

    if len(nums1) == 0 {
        return median(nums2)
    }
    if len(nums2) == 0 {
        return median(nums1)
    }
    merger := make([]int, len(nums1)+len(nums2))
    for i := 0; i < len(merger); i++ {
        if nums1[0] < nums2[0] {
            merger[i] = nums1[0]
            if len(nums1) == 1 {
                copy(merger[i+1:], nums2)
                break
            }
            nums1 = nums1[1:]
        } else {
            merger[i] = nums2[0]
            if len(nums2) == 1 {
                copy(merger[i+1:], nums1)
                break
            }
            nums2 = nums2[1:]
        }
    }
    return median(merger)
}
