
### 用户数据表——`dh_api_data` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|YES|唯一id|
|user_id|varchar(32)|NO|用户id|
|type|varchar(16)|NO|类型|
|content|text|NO|内容|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 评论表——`dh_comment` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|relate_type|varchar(16)|NO|关联类型|
|relate_id|varchar(32)|NO|关联id|
|user_id|varchar(32)|NO|用户id|
|type|varchar(16)|NO|类型|
|content|varchar(1000)|NO|内容|
|replay_id|varchar(32)|NO|回复id|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据库连接表——`dh_connect` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|crop_id|varchar(32)|NO|团队id|
|user_id|varchar(32)|NO|用户id|
|name|varchar(64)|NO|名称|
|type|varchar(16)|NO|类型|
|config|text|NO|配置项(json格式)|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 团队表——`dh_corp` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|name|varchar(128)|NO|名称|
|email|varchar(128)|NO|邮箱|
|mobile|varchar(32)|NO|手机号|
|vcode|varchar(16)|NO|验证码|
|status|tinyint(4)|NO|状态:0——新建，1——审核通过，-1——删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 故事版表——`dh_dashboard` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|group_id|varchar(32)|NO|分组id|
|name|varchar(64)|NO|名称|
|thumbnail|text|NO|缩略图(base64字符串)|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 看版表——`dh_dashboard_group` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|name|varchar(64)|NO|名称|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 看板预警表——`dh_dashboard_warning` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|widget_id|varchar(32)|NO|组件id|
|key|varchar(32)|NO|预警字段|
|compare|varchar(8)|NO|预警条件|
|value|bigint(20)|NO|预警值|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 故事版页面组件表——`dh_dashboard_widget` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|dashboard_id|varchar(32)|NO|看板id|
|grid|varchar(100)|NO|定位(json格式)|
|config|text|NO|配置|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据源表——`dh_datasource` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|group_id|varchar(32)|NO|分组id|
|name|varchar(64)|NO|名称|
|table|varchar(64)|NO|数据表|
|type|varchar(8)|NO|类型|
|format|varchar(8)|NO|格式|
|url|varchar(11)|NO|来源url|
|connect_id|varchar(32)|NO|链接id|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据源组表——`dh_datasource_group` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|name|varchar(64)|NO|名称|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据库字段描述表——`dh_datasource_meta` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|datasource_id|varchar(32)|NO|数据源id|
|column|varchar(64)|NO|列|
|name|varchar(64)|NO|名称|
|type|varchar(16)|NO|类型|
|extra|varchar(64)|NO|额外描述|
|status|tinyint(4)|NO|状态:1—正常，0—隐藏|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据库字段更改记录表——`dh_datasource_modify` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|user_id|varchar(32)|NO|用户id(操作人)|
|datasource_id|varchar(32)|NO|数据源id|
|table|varchar(64)|NO|表名称|
|row_id|varchar(128)|NO|行号|
|column|varchar(64)|NO|列|
|type|varchar(16)|NO|类型|
|old_value|varchar(1000)|NO|旧值|
|new_value|varchar(1000)|NO|新值|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 数据源关联表——`dh_datasource_relation` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|datasource1|varchar(32)|NO|数据源1|
|datasource2|varchar(32)|NO|数据源2|
|join|varchar(8)|NO|关联方式|
|on|varchar(500)|NO|关联条件|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 邀请码表——`dh_icode` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|code|varchar(16)|NO|编码|
|status|tinyint(4)|NO|状态:0—新建，1—已经使用，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 链接分享表——`dh_linkshare` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|relate_type|varchar(16)|YES|关联类型|
|relate_id|varchar(32)|NO|关联id|
|type|varchar(16)|NO|类型|
|password|varchar(16)|NO|密码|
|url|varchar(256)|NO|分析链接|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 通知表——`dh_notify` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|from_crop_id|varchar(32)|NO|来源团队id|
|from_user_id|varchar(32)|NO|来源用户id|
|user_id|varchar(32)|NO|用户id|
|type|varchar(16)|NO|类型|
|config|varchar(1000)|NO|通知配置(json格式)|
|status|tinyint(4)|NO|状态:0—新建，1—已读，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 关联关系表——`dh_relation` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|crop_id|varchar(32)|NO|团队id|
|user_id|varchar(32)|NO|用户id|
|relate_type|varchar(16)|NO|关联类型|
|relate_id|varchar(32)|NO|关联id|
|auth|varchar(16)|NO|权限|
|name|varchar(64)|NO|自定义名称|
|sort|tinyint(4) unsigned|NO|自定义排序|
|create_time|timestamp|NO|新增时间|
|uodate_time|timestamp|YES|修改时间|

### 故事版表——`dh_storyboard` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|group_id|varchar(32)|NO|分组id|
|name|varchar(64)|NO|名称|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 故事版组表——`dh_storyboard_group` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|name|varchar(64)|NO|名称|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 故事版页面表——`dh_storyboard_page` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|board_id|varchar(32)|NO|故事版id|
|thumbnail|text|NO|缩略图(base64字符串)|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|sort|tinyint(4) unsigned|NO|排序|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 故事版页面组件表——`dh_storyboard_widget` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|page_id|varchar(32)|NO|故事版页面id|
|type|varchar(16)|NO|类型|
|position|varchar(100)|NO|定位(json格式)|
|content|text|NO|内容|
|status|tinyint(4)|NO|状态:0—正常，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 用户表——`dh_user` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|name|varchar(128)|NO|名称|
|corp|varchar(128)|YES|公司|
|email|varchar(128)|NO|邮箱|
|mobile|varchar(32)|NO|手机号|
|password|varchar(128)|NO|密码|
|auth|varchar(128)|NO|授权号|
|avatar|varchar(256)|NO|头像|
|icode|varchar(16)|NO|邀请码|
|status|tinyint(4)|NO|状态:0——新建，1——审核通过，-1——删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 用户团队表——`dh_user_corp` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|user_id|varchar(32)|NO|用户id|
|crop_id|varchar(32)|NO|团队id|
|role|varchar(16)|NO|角色|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|

### 手机号验证码表——`dh_vcode` ###
===============
|字段名称|字段类型|是否可空|描述|
|---|---|---|---|
|id|int(11) unsigned|NO|自增id|
|object_id|varchar(32)|NO|唯一id|
|mobile|varchar(16)|NO|手机号|
|code|varchar(16)|NO|编码|
|status|tinyint(4)|NO|状态:0—新建，1—已经使用，-1—删除|
|create_time|timestamp|NO|新增时间|
|update_time|timestamp|YES|修改时间|
