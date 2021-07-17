import requests
import json
import pandas as pd
from sqlalchemy import create_engine
import logging


def update_leetcode_data():
    url = 'https://leetcode-cn.com/api/problems/all/'

    header = {
        'cookie': 'LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuZXh0X2FmdGVyX29hdXRoIjoiL3Byb2JsZW1zL21heGltdW0taWNlLWNyZWFtLWJhcnMvIiwiX2F1dGhfdXNlcl9pZCI6IjIwNzg3OTQiLCJfYXV0aF91c2VyX2JhY2tlbmQiOiJkamFuZ28uY29udHJpYi5hdXRoLmJhY2tlbmRzLk1vZGVsQmFja2VuZCIsIl9hdXRoX3VzZXJfaGFzaCI6IjEzNWJmNjk0MWIzMGYwNmMxNmI1YjkzMjVlYzU4YWE5NjA4MWY3M2NlZGY3OWE2NzJhMWVhZGUyY2NhYjU5NmQiLCJpZCI6MjA3ODc5NCwiZW1haWwiOiIiLCJ1c2VybmFtZSI6InJpLXl1ZS10b25nLWh1aS10IiwidXNlcl9zbHVnIjoicmkteXVlLXRvbmctaHVpLXQiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS1jbi5jb20vYWxpeXVuLWxjLXVwbG9hZC91c2Vycy9yaS15dWUtdG9uZy1odWktdC9hdmF0YXJfMTYwNTQ5NjE1NS5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTYyNTIyODkyMS4yMzIyNDYyLCJleHBpcmVkX3RpbWVfIjoxNjI3NzU4MDAwLCJsYXRlc3RfdGltZXN0YW1wXyI6MTYyNjE2MTU4NX0.Y0WJwr9KNeKFYEZRtNQPJUTZolh4ZbcWre96fruEXw4;'
    }

    res = requests.get(url, headers=header, verify=False)
    data = json.loads(res.text)
    print('username = ', data['user_name'])
    stats = data['stat_status_pairs']
    num_solved, num_total = data["num_solved"], data['num_total']
    ac_easy, ac_medium, ac_hard = data['ac_easy'], data['ac_medium'], data['ac_hard']

    res_dic = {
        'question_id': [],
        'sp_id': [],
        'title': [],
        'status': [],
        'difficulty': [],
        'url': []
    }

    for stat in stats:
        cur_data = stat['stat']
        # print(cur_data['question_id'], cur_data['question__title'], stat['status'], stat['difficulty']['level'])
        res_dic['question_id'].append(cur_data['question_id'])
        res_dic['sp_id'].append(cur_data['frontend_question_id'])
        res_dic['title'].append(cur_data['question__title'])
        res_dic['status'].append(stat['status'])
        res_dic['difficulty'].append(stat['difficulty']['level'])
        res_dic['url'].append(cur_data['question__title_slug'])

    # # print(res_dic)
    data_frame = pd.DataFrame(res_dic).set_index('question_id')
    print(data_frame)

    if data['user_name']:
        engine = create_engine('mysql+pymysql://root:540279653@localhost/my_data?charset=utf8mb4', encoding='utf-8')
        data_frame.to_sql('leetcode', con=engine, if_exists='replace')
        print('Complete...')
    else:
        logging.info("Error: You did not get valid data, please check your cookie...")
        return False
    return True


if __name__ == '__main__':
    update_leetcode_data()
