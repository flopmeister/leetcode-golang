""" 
Ideas are from: https://discuss.leetcode.com/topic/30421/share-my-thinking-process/119
Assumption: every day, we have three possible states, rest, buy, sell.
             rest must happen before buy, and buy must happen before sell
State representation: each state[i] represents the maximum profit it can gain
    rest[i] = max(buy[i-1], sell[i-1], rest[i-1])
    buy[i] = max(rest[i-1]-prices[i], buy[i-1])
    sell[i] = max(buy[i-1]+prices[i], sell[i-1])

Note: we have buy[i] <=rest[i] <=sell[i]
      and it's obvious, rest[i] == rest[i-1]
Therefore, above equation can be changed to
    rest[i] = sell[i-1]
    buy[i] = max(sell[i-2]-prices[i], buy[i-1])
    sell[i] = max(buy[i-1]+prices[i], sell[i-1])
    
    sell[i-2] can be represented as the day before yesterday's sell choice
    sell[i-1] can be represented as yesterday's sell choice
    therefore: four states required only
    last_buy = buy
    buy = max(last_sell - price, last_buy) // here last_sell is sell[i-2]
    last_sell = sell                       // last_sell now is sell[i-1]
    sell = max(last_buy + price, last_sell) 

"""
class Solution:
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :type: int
        """
        if len(prices) < 2:
            return 0
        buy, sell, last_buy, last_sell = -prices[0], 0, 0 ,0
        for price in prices:
            last_buy=buy
            buy = max(last_buy, last_sell-price)
            last_sell=sell
            sell=max(last_buy+price, sell)
        return sell