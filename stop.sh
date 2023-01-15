ps -ef | grep -i `pwd` | grep -v grep | awk '{print $2}' | xargs kill
