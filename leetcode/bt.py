class Solution(object):
    def permute(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        length = len(nums)
        res = []
        def permute_dfs(numset , checked, deepth):
            if deepth == length:
                res.append(checked[:])
                return

            for i in range(length):
                if numset[i] in checked:
                    continue
                checked.append(numset[i])       # 先向下递归
                print('-->', checked)
                permute_dfs(numset, checked, deepth+1)
                checked.pop()                   # 回溯到上一个节点
                print('<--', checked)

        permute_dfs(nums , [] , 0)
        return res



def main():
    nums = [1, 2, 3]
    s = Solution().permute(nums)
    print(s)
    pass

if __name__ == '__main__':
    main()
    