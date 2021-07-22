import Leetcode
import time
import requests

if __name__ == '__main__':
    while True:
        if not Leetcode.update_leetcode_data():
            break
        requests.get("http://123.57.66.63:8080/api/v1/leetcode/update-statistics")
        time.sleep(1800)
