<template>
    <div>
        <question :question_description="description" :question_choices="selection"></question>
        
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
      answer: []
    }
  },

  components: {
    question
  }
}
</script>