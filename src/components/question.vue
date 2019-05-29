<template>
  <div>
    <div class="question">
      <p>{{question_description}}</p>
      
        <div v-for="(content, choice_id) in question_choices" v-bind:key="choice_id">
          <input type="radio" :v-model="unique_select" :value=content> {{content}}
        </div>
        
    </div>
  </div>
</template>

<script>
export default {
  props: {
    'question_description': {
      type: String,
      default: 'question_description'
    },
    'question_choices': {
      type: Object,
      default: () => {}
    }
  }
}
</script>

<style>
.question {
  padding: 10px;
}
</style>
<!--
  <input type="checkbox" id="A" value="A" v-model="checkedNames">
        <label for="A">{{question_selection}}{{A}}</label>
        <input type="checkbox" id="B" value="B" v-model="checkedNames">
        <label for="B">{{question_selection}}{{B}}</label>
        <input type="checkbox" id="C" value="C" v-model="checkedNames">
        <label for="C">{{question_selection}}{{C}}</label>
        <input type="checkbox" id="D" value="D" v-model="checkedNames">
        <label for="D">{{question_selection}}{{D}}</label>
        <br>
        <span>选择的值为: {{ checkedNames }}</span>
  -->