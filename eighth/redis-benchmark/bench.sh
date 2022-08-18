#!/bin/sh


redis-benchmark -t get,set -d 10 >> result_10
redis-benchmark -t get,set -d 20 >> result_20
redis-benchmark -t get,set -d 50 >> result_50
redis-benchmark -t get,set -d 100 >> result_100
redis-benchmark -t get,set -d 200 >> result_200
redis-benchmark -t get,set -d 1000 >> result_1000
redis-benchmark -t get,set -d 5000 >> result_5000


