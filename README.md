Taken from https://github.com/karalabe/cookiejar/tree/master/collections/queue 
with the only exception been that `karalabe/cookiejar` is a generic queue(it uses an interace{} as main type), this one is a queue of strings.                 
And I also added race protection.                


// before race protection;                    
BenchmarkPush-8   	10000000	       157 ns/op         
BenchmarkPop-8    	20000000	       106 ns/op

// after race protection;                
BenchmarkPush-8   	 3000000	       394 ns/op             
BenchmarkPop-8    	 5000000	       403 ns/op

