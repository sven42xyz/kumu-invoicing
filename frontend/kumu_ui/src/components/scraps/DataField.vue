<template>
      <div class="flex justify-content-center">
        <div v-if="disabled==true" class="flex">
            <label for="field">{{title}}</label>
            <InputText id="fieldText" disabled v-model="values" :placeholder="value" />
        </div>
        <div v-else class="flex">
            <label for="field">{{title}}</label>
            <InputText id="fieldText" v-model="values">{{ value }}</InputText>
        </div>
      </div>
</template>
  
  <script>

  export default {
    props:{    
        disabled:{
            type: Boolean,
        },
        title:{
            type: String,
        },
        value:{
            type: String,
        }
    },

    created(){
    },

    data(){
   
    },

    computed: {

    }
  }
  </script>
  
  <style scoped>
  </style>