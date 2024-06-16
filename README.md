# bilibili收藏夹下载器
一个能快速下载bilibili原始视频的脚本。<br>
go语言编写，支持linux，方便归档在服务器中。

## 用法
- 将assets目录放在可执行文件**同级目录**，然后执行即可
- assets目录中的config.ini.temp请自行参考，使用时可将文件**重命名为config.ini**
- 首次执行会在assets目录中生成二维码，并等待扫码，手机bilibili扫一下登录就好，之后他自己会保存cookie信息，除非强制下线或者过期，通常下次就不需要再扫了

## ffmpeg
- 在linux中使用请**保证已安装ffmpeg**，我的版本是5.1.4，理论上大于这个即可
- 在windows中就去官网下载ffmpeg，然后和assets一样放在同级目录即可，注意**目录名必须是ffmpeg**
