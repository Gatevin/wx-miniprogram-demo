<template>
    <div>
        <question :question_description="description" :question_choices="selection"></question>
        <button @click="requestData"> 点我请求服务器 </button> <div> {{ message_from_server }} </div>
    </div>
</template>

<script>
import question from '@/components/question'

export default {
  data () {
    return {
      question_id: 0,
      description: '选出你喜欢的搜索引擎吧～',
      selection: {
        A: '谷歌',
        B: '百度',
        C: '必应',
        D: '雅虎'
      },
      answer: [],
      message_from_server: 'server-message'
    }
  },

  components: {
    question
  },

  methods: {
    // 这里之后从服务端请求得到的问题json文件拿到问题数据
    requestData () {
      // 这个地方的that有必要...暂时还不知道为什么
      var that = this
      var api = 'http://localhost:23333/'
      wx.request({
        url: api,
        data: {
          question_id: 0
        },
        header: {
          'content-type': 'text'
        },
        success: function (res) {
          console.log(res.data)
          that.message_from_server = res.data
        }
      })
    },
    // 初始的时候调用
    created () {
      this.requestData()
    }
  }

}
</script>