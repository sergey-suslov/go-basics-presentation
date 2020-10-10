interface NumberHolder {
  num: number
}

const forEach = (nums: NumberHolder[]) => {
  nums.forEach(num => {
    num.num++
  })
  return nums
}

const map = (nums: NumberHolder[]) => {
  return nums.map(n => ({...n, num: n.num++}))
}

const forI = (nums: NumberHolder[]) => {
  for (let i = 0; i < nums.length; i++) {
    nums[i].num++
  }
  return nums
}

const forOf = (nums: NumberHolder[]) => {
  for (let n of nums) {
    n.num++
  }
  return nums
}

const forIn = (nums: NumberHolder[]) => {
  for (let n in nums) {
    if (n in nums){
      nums[n].num++
    }
  }
  return nums
}
