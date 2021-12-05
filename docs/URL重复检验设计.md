
# 1. select * from link where url=?


# 2. select * from link where link_hash = ?

hash(url)=32bit

# 3. cuckoo filter redis

劣：
1. redis 没法用云了=>redis运维成本
  
   redis bloom filter 企业版才支持
   
   云redis 不支持module load 
   
2. 网络通信



# 4. cuckoo filter memory

优：

no need redis io 


劣：

 重启之后，filter需要重建

https://github.com/irfansharif/cfilter  

https://github.com/seiflotfy/cuckoofilter

小写无法序列化保存，需要fork之后修改field为大写开头


# 综合考虑 使用方案 #4