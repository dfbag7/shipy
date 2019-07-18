curl http://192.168.243.154:8080/rpc ^
 -XPOST ^
 --verbose ^
 -H "Content-Type: application/json" ^
 --data-binary @auth-data
 

exit /b

curl "http://209.222.20.203/application/dashboard/revisionsCreatedByUser" ^
-X POST ^
-H "Cookie: remember_b8eb1e1be5db847046a2188a2075cc52=100^%^7CAuKaPC8OcueZRIPdJLYGX2dOxg48AM29q2GCNjZWiYsI3ZfoG3sttkXWsXwL; XSRF-TOKEN=l4peNAsEqCFVYa2TOurRJ9s5pBMQppFOJj7qEftt; laravel_session=e46a04cfb3fb1d91c16889738890acc009b1735d" ^
-H "Origin: http://209.222.20.203" ^
-H "X-XSRF-TOKEN: l4peNAsEqCFVYa2TOurRJ9s5pBMQppFOJj7qEftt" ^
-H "Accept-Language: en-US,en;q=0.8,ru;q=0.6" ^
-H "User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36" ^
-H "Content-Type: application/json;charset=UTF-8" ^
-H "Accept: application/json, text/plain, */*" ^
-H "Referer: http://209.222.20.203/" ^
-H "Accept-Encoding: gzip, deflate" ^
-H "Connection: keep-alive" ^
--data-binary @data --compressed
