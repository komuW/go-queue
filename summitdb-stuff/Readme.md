# Install and run:
1. redis-cli
```bash
wget http://download.redis.io/releases/redis-4.0.6.tar.gz
tar xzf redis-4.0.6.tar.gz
cd redis-4.0.6
make                
redis-4.0.6/src/redis-cli -p 7481
```
2. summitDB
```
docker-compose up --build
```

# RPUSH:
`script load "sdb.call('set', KEYS[0], '[]', 'nx');sdb.call('jset', KEYS[0], '-1', ARGV[0]);return parseInt(sdb.call('jget', KEYS[0], '#'));"`                
  "95e4005b8d5dff3778154ae39dd819a21c363d13"                     

Then you can use the script like:

`evalsha 95e4005b8d5dff3778154ae39dd819a21c363d13 1 myqueue foo`                 
1                    
`evalsha 95e4005b8d5dff3778154ae39dd819a21c363d13 1 myqueue item2`           
2                
`evalsha 95e4005b8d5dff3778154ae39dd819a21c363d13 1 myqueue item3`             
3                 
`evalsha 95e4005b8d5dff3778154ae39dd819a21c363d13 1 myqueue 22`             
4                   

`GET myqueue`              
"[\"foo\", \"item1\", \"item2\", 22]"          


# LPOP:
`script load "var res = sdb.call('jget', KEYS[0], '0');sdb.call('jdel', KEYS[0], '0');return res;"`                  
  "e3f15d9503ca033132d0e83b35ec3afcf28ef922"            

then use it like:        

`evalsha e3f15d9503ca033132d0e83b35ec3afcf28ef922 1 myqueue`            
"foo"            
`evalsha e3f15d9503ca033132d0e83b35ec3afcf28ef922 1 myqueue`                
"item1"                

`GET myqueue`                  
"[\"item2\", 22]"         


# documentation      
1. https://redis.io/commands/eval                

`eval "return redis.call('set','foo','bar')" 0`         
OK            
The above script sets the key foo to the string bar. However it violates the EVAL(https://redis.io/commands/eval) command semantics as             
all the keys that the script uses should be passed using the KEYS array. ie EVAL signature is;
`EVAL script numkeys key [key ...] arg [arg ...]`             
so we fix it as;            
`eval "return redis.call('set',KEYS[1],'bar')" 1 foo`           
OK              

Since we are using summitDB and not redis, the equivalent is;             
`eval "return sdb.call('set',KEYS[0],'value');" 1 mykey`               
`GET mykey`                 
"value"               

2. https://github.com/tidwall/summitdb/issues/22


# MISC
To create multiple queues and RPUSH to them;             
`EVALSHA sha1 1 queue1 val1`           
`EVALSHA sha1 1 queue1 val2`         
`EVALSHA sha1 1 queue1 val3`         
           
`EVALSHA sha1 1 myqueue2 value1`           
`EVALSHA sha1 1 myqueue2 value2`         
`EVALSHA sha1 1 myqueue2 value3`         
          
SummitDB appends all writeable commands to a persistent file, to see it:               
`tail -fn100 data/summitdb/raft.db`           
            
To see state of leaders/followers:               
`redis-4.0.6/src/redis-cli -p 7481`             
`> RAFTSTATE`           
"Leader"               

`redis-4.0.6/src/redis-cli -p 7482`           
`> RAFTSTATE`           
"Follower"             

`redis-4.0.6/src/redis-cli -p 7483`              
`> RAFTSTATE`                
"Follower"            


