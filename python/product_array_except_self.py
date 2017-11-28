class Solution(object):
def productExceptSelf(self, nums):
	"""
	:type nums: List[int]
	:rtype: List[int]
	"""
	n = len(nums)
	p = 1
	s = 1
	output = []
	for i in range(0,n):
		output.append(p)
		p = p * nums[i]
	for i in range(n-1,-1,-1):
		output[i] *= s
		s *= nums[i]
	return output
	