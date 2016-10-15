#/bin/bash
./errorformat "%E[%t%.%+] %f:%l: error: %m" "%A[%t%.%+] %f:%l: %m" "%Z[%.%+] %p^" "%C[%.%+] %.%#" "%-G%.%#" < testdata/sbt.in
