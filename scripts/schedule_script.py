import Leetcode
import time

if __name__ == '__main__':
    while True:
        if not Leetcode.update_leetcode_data():
            break
        time.sleep(1800)
