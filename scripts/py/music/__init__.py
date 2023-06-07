import time

import mysql.connector
import oss2
import requests
from alibabacloud_sts20150401 import models as sts_20150401_models
from alibabacloud_sts20150401.client import Client as Sts20150401Client
from alibabacloud_tea_openapi import models as open_api_models
from alibabacloud_tea_util import models as util_models
from bs4 import BeautifulSoup

# 已完成歌手：'邓紫棋'

singers = ['蔡徐坤', '周杰伦', '张学友', '王菲', '林忆莲', '张靓颖', '范晓萱', '陈奕迅', '林俊杰', '罗志祥', '蔡依林']
search_url = "http://www.44h4.com/so/"  # 标准格式 "http://www.44h4.com/so/邓紫棋.html"
base_url = "http://www.44h4.com"
access_key = "LTAI5t5wsegQbfbzDqCiAKzC"
access_secret = "Frip75apOT2kdzhXW3SfWWUZVWOkv4"
role_arn = "acs:ram::1853932617531012:role/ramoss"
role_session_name = "pangchat-sts"
oss_endpoint = "oss-cn-beijing.aliyuncs.com"
sts_endpoint = "sts.cn-beijing.aliyuncs.com"
oss_bucket_name = "pangchat-media"
# 定义全局bucket对象
bucket = None
user = 'root'
password = 'password'
host = '43.143.137.232'
port = 3307
database = 'pangchat'
cursor = None
conn = None


# 定义一个函数，用于处理某个歌曲列表
def get_songobj_list(song_list, singer):
    # 定义一个空的列表，用于存放歌曲的url和名称
    songObjList = []
    # 遍历歌曲列表
    for song in song_list:
        # 获取歌曲的url
        song_url = song.get('href')
        # 获取歌曲的名称
        song_name = song.get_text()
        # 将歌曲的名称和url转化为对象
        if song_name == 'mv':
            continue
        # 处理歌曲名称，部分名称中带有 /,需要进行转义
        song_name = song_name.replace('/', '-')
        songObj = {'song_name': song_name, 'song_url': base_url + song_url, 'singer': singer}
        # 将歌曲对象添加到列表中
        songObjList.append(songObj)
    return songObjList


# 定义一个函数，用于获取歌曲的下载地址和lrc地址
def get_download_url(songObjList):
    for songObj in songObjList:
        songObj['is_valid'] = True
        print("正在爬取歌曲：" + songObj['song_name'] + "的下载地址和lrc地址,还剩" + str(
            len(songObjList) - songObjList.index(songObj) - 1) + "首歌曲")
        # 获取歌曲的url
        song_url = songObj['song_url']
        # 发起请求
        resp = requests.get(song_url)
        # 获取响应的文本
        resp_text = resp.text
        # 使用BeautifulSoup解析文本
        soup = BeautifulSoup(resp_text, 'html.parser')
        # 获取歌曲下载地址,div.dance_wl下的a标签
        srcs = soup.select('div.dance_wl a')
        if len(srcs) == 0:
            continue
        if len(srcs) < 2:
            songObj['is_valid'] = False
            continue
        song_download_url = srcs[1].get('href')
        # 保存歌曲下载地址
        songObj['song_download_url'] = base_url + song_download_url
        # 获取歌曲lrc地址,div.dance_wl下的第三个a标签的href属性
        song_lrc_url = soup.select('div.dance_wl a')[2].get('href')
        # 保存歌曲lrc地址
        songObj['song_lrc_url'] = base_url + song_lrc_url
        # 保存歌曲是否有效
    return songObjList


# 定义一个函数，用于去除无效的歌曲
def remove_invalid_song(songObjList):
    # 去除无效的歌曲
    songObjList = [songObj for songObj in songObjList if songObj['is_valid']]
    return songObjList


# 定义一个函数，用于下载歌曲
def download_song(songObjList):
    # 遍历歌曲对象列表
    for songObj in songObjList:
        try:
            # 打印正在下载的歌曲的名称，歌手，下载地址，lrc地址
            print("正在下载歌曲：" + songObj['song_name'] + "，歌手：" + songObj['singer'] + "，下载地址：" + songObj[
                'song_download_url'] + "，lrc地址：" + songObj['song_lrc_url'])
            # 获取歌曲的下载地址
            song_download_url = songObj['song_download_url']
            # 保存为mp3文件
            save_song(song_download_url, songObj['song_name'], songObj['singer'])
            save_song_lrc(songObj['song_lrc_url'], songObj['song_name'], songObj['singer'])
            time.sleep(0.5)
        except Exception as e:
            print("下载歌曲：" + songObj['song_name'] + ",出现异常，异常信息为：" + str(e))
            continue


# 定义一个函数，用于保存歌词到数据库
def save_song_lrc(song_lrc_url, song_name, song_singer):
    # song_name 分离出歌手和歌曲名
    parts = song_name.split('-')
    # 歌曲名,去除空格
    song_name = parts[0].strip()
    # 歌手名，去除空格
    singer = parts[1].strip()
    if singer != song_singer:
        return
    resp = requests.get(song_lrc_url)
    # 获取响应的文本
    resp_text = resp.text
    # 构建sql语句
    sql = 'insert into song_lrcs(lrc_song_name,lrc_singer_name,lrc_song_content) values(%s,%s,%s)'
    # 执行sql语句
    cursor.execute(sql, (song_name, singer, resp_text))
    conn.commit()


# 定义一个函数,保存歌曲到oss
def save_song(song_download_url, song_name, song_singer):
    # song_name 分离出歌手和歌曲名
    parts = song_name.split('-')
    # 歌手名，去除空格
    singer = parts[1].strip()
    if singer != song_singer:
        return
    # 构建歌曲在oss中的路径
    path = "music/" + song_name + ".mp3"
    # 如果oss中已经存在该歌曲，则跳过
    if bucket.object_exists(path):
        print("歌曲：" + song_name + "已经存在，跳过")
        return
    # 发起请求
    resp = requests.get(song_download_url)
    bucket.put_object(path, resp)
    # 构建sql语句
    sql = 'insert into songs(song_name) values(%s)'
    # 执行sql语句
    cursor.execute(sql, (song_name,))
    conn.commit()


# 定义一个函数，用户处理某个歌手的歌曲
def handle_singer(singer):
    print("正在爬取歌手：" + singer + "的歌曲,还剩" + str(len(singers) - singers.index(singer) - 1) + "位歌手")
    # 发起requests请求
    resp = requests.get(search_url + singer + ".html")
    # 获取响应的文本
    resp_text = resp.text
    # 使用BeautifulSoup解析文本
    soup = BeautifulSoup(resp_text, 'html.parser')
    # 获取div.play_list下的所有li标签的a标签的href属性和text
    song_list = soup.select('div.play_list li a')
    # 定义一个空的列表，用于存放歌曲的url和名称
    songObjList = get_songobj_list(song_list, singer)
    # 获取歌曲的下载地址和lrc地址
    songObjList = get_download_url(songObjList)
    # 去除无效的歌曲
    songObjList = remove_invalid_song(songObjList)
    # 下载歌曲
    download_song(songObjList)


# 定义一个函数，用于处理所有歌手的歌曲
def handle_singers(s):
    for singer in s:
        handle_singer(singer)


# 初始化AcsClient，用于发起请求
def init_sts(access_key_id, access_key_secret):
    config = open_api_models.Config(
        # 必填，您的 AccessKey ID,
        access_key_id=access_key_id,
        # 必填，您的 AccessKey Secret,
        access_key_secret=access_key_secret,
        endpoint=sts_endpoint,
    )
    return Sts20150401Client(config)


def serve():
    client = init_sts(access_key, access_secret)
    assume_role_request = sts_20150401_models.AssumeRoleRequest(
        # 指定角色Arn
        role_arn=role_arn,
        role_session_name=role_session_name,
    )
    runtime = util_models.RuntimeOptions()
    runtime.timeout = 10000
    try:
        res = client.assume_role_with_options(assume_role_request, runtime)
        # 从res中分分离出，AccessKeyId,AccessKeySecret,SecurityToken的值
        res = res.to_map()
        access_key_id = res.get('body').get('Credentials').get('AccessKeyId')
        access_key_secret = res.get('body').get('Credentials').get('AccessKeySecret')
        security_token = res.get('body').get('Credentials').get('SecurityToken')
        return {
            'access_key_id': access_key_id,
            'access_key_secret': access_key_secret,
            'security_token': security_token
        }
    except Exception as e:
        print(e)


# 初始化oss
def init_oss_bucket():
    credential = serve()
    sts_auth = oss2.StsAuth(credential['access_key_id'], credential['access_key_secret'],
                            credential['security_token'])
    # 将sts_auth转为auth
    oss_bucket = oss2.Bucket(auth=sts_auth, endpoint=oss_endpoint, bucket_name=oss_bucket_name, region="oss-cn-beijing")
    return oss_bucket


# 初始化数据库
def init_db():
    return mysql.connector.connect(user=user, password=password, host=host, port=port, database=database)


# 初始化cursor
def init_cursor():
    return conn.cursor()


# main
if __name__ == '__main__':
    bucket = init_oss_bucket()
    conn = init_db()
    cursor = init_cursor()
    handle_singers(singers)
