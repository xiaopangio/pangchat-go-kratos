import time

import mysql.connector
import requests
from bs4 import BeautifulSoup

# 热门歌手列表  已完成'邓紫棋', '周杰伦','张学友', '王菲', '林忆莲', '张靓颖', '范晓萱', '陈奕迅','林俊杰', '罗志祥', '蔡依林',
# '范玮琪', '韩红', '徐佳莹', '李宇春', '华晨宇', '刘若英', '曲婉婷', '陈慧娴', '许嵩'
singers = []
# 查找热门歌手的所有歌曲，search_url=http://www.dsod.cn/e/search/ 请求为POST，请求体为 {keyboard: '歌手名', show: 'title,newstext',
# tempid: 1, tbname: 'lrc', mid: 1, dopost: 'search'}
search_url = 'http://www.dsod.cn/e/search/'
# 初始化数据库连接
# pangchat:password@tcp(43.143.137.232:3307)/pangchat
user = 'root'
password = 'password'
host = '43.143.137.232'
port = 3307
database = 'pangchat'
conn = mysql.connector.connect(user=user, password=password, host=host, port=port, database=database)
cursor = conn.cursor()
last_unfinished_singer_name = ''
last_unfinished_song_count = 0
# 遍历热门歌手列表
for singer in singers:
    resp = requests.post(search_url,
                         data={'keyboard': singer, 'show': 'title,newstext', 'tempid': 1, 'tbname': 'lrc',
                               'mid': 1, 'dopost': 'search'})
    soup = BeautifulSoup(resp.text, 'html.parser')
    # 获取歌曲列表,为div.rec 下面的所有li标签下的a标签的href属性
    song_list = soup.select('div.rec li a')
    print("正在获取歌手：", singer, "还剩", len(singers) - singers.index(singer) - 1, "位歌手")
    if last_unfinished_singer_name == singer:
        # 截取后last_unfinished_song_count首歌曲
        song_list = song_list[len(song_list) - last_unfinished_song_count:]
    for song in song_list:
        # 获取歌曲的url
        song_url = song.get('href')
        # 获取歌曲的名称
        song_name = song.get_text()
        # 打印歌曲的名称和url
        # print(song_name, song_url)
        # 根据歌曲的url获取歌词，/lrc/94188.html 转为 /lrcdown/94188.html,再加上前缀http://www.dsod.cn
        lrc_url = 'http://www.dsod.cn' + song_url.replace('/lrc/', '/lrcdown/')
        # 获取歌词的响应，结果为附件
        lrc_resp = requests.get(lrc_url)
        # 获取附件的内容，为二进制
        lrc_content = lrc_resp.content
        # 解码为utf-8
        lrc_content = lrc_content.decode('utf-8')
        # 将\r 替换为\n,开头去掉空行
        lrc_content = lrc_content.replace('\r', '\n').lstrip('\n')
        print("正在保存歌词：", song_name, "还剩", len(song_list) - song_list.index(song) - 1, "首歌曲")
        # 构建sql语句
        sql = 'insert into song_lrcs(lrc_song_name,lrc_singer_name,lrc_song_content) values(%s,%s,%s)'
        # 执行sql语句
        execute = cursor.execute(sql, [song_name, singer, lrc_content])
        # 提交事务
        conn.commit()
        # 暂停1秒
        time.sleep(0.5)
    time.sleep(10)
# 关闭数据库连接
cursor.close()
conn.close()
