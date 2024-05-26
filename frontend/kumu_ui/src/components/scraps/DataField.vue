<template>
    <div class="row" v-if="disabled==true" >
        <label for="field">{{title}}</label>
    </div>
    <div v-else class="row">
        <label for="field">{{title}}</label>
    </div>
    <div class="row" v-if="disabled==true" >
        <InputText id="fieldText" disabled v-model="values" :placeholder="value" />
    </div>
    <div v-else class="row">
        <InputText id="fieldText" v-model="values">{{ value }}</InputText>
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
    .row{
        width: 40%;
        padding-left: 5vmin;
        text-align: start !important;
    }
  </style>