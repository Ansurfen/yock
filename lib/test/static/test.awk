{ num[$2]++ } END { for (m in num) print m, num[m] }