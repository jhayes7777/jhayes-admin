<!--
    @auther: bypanghu<bypanghu@163.com>
    @date: 2024/5/8
!-->

<template>
  <div>
    <el-table :data="tableData" stripe style="width: 100%">
      <el-table-column prop="ranking" label="排名" width="80" align="center" />
      <el-table-column prop="title" label="咨询标题" show-overflow-tooltip>
        <template #default="{ row }">
          <a class="text-active" :href="row.link" target="_blank">{{
            row.title
          }}</a>
        </template>
      </el-table-column>
      <el-table-column prop="click_num" label="关注度" width="100" />
      <el-table-column prop="hot" label="热度值" width="100" />
    </el-table>
  </div>
</template>

<script setup>
  const tableData = [
    {
      ranking: 1,
      title: '字节跳动旗下零代码 AI-Coze - 快速搭建属于你的专属AI助手。',
      click_num: 523,
      hot: 263,
      link: 'https://space.coze.cn/'
    },
    {
      ranking: 2,
      title:
        '【豆包】全能写作bot - 智能写作，豆包AI助你高效',
      click_num: 416,
      hot: 223,
      link: 'https://www.doubao.com/chat/'
    },
    {
      ranking: 3,
      title:
        '腾讯元宝满血版DeepSeek,问得快答得准!',
      click_num: 337,
      hot: 176,
      link: 'https://yuanbao.tencent.com/chat/naQivTmsDa'
    },
    {
      ranking: 4,
      title:
        '文心一言既是你的智能伙伴，可以陪你聊天、回答问题、画图识图；也是你的AI助手，可以提供灵感、撰写文案、阅读文档、智能翻译，帮你高效完成工作和学习任务。',
      click_num: 292,
      hot: 145,
      link: 'https://yiyan.baidu.com/'
    },
    {
      ranking: 5,
      title: 'AI人工智能学习网站大全 | AI工具集',
      click_num: 173,
      hot: 110,
      link: 'https://ai-bot.cn/favorites/websites-to-learn-ai/'
    }
  ]
</script>

<style scoped lang="scss"></style>
