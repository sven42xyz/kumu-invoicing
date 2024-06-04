<template>
    <div class="row" v-if="disabled==true" style="margin-top: 0.75vmin;" >
        <label for="field">{{title}}</label>
        <InputText id="fieldText" disabled v-model="values" :placeholder="value" />
    </div>
    <div v-else class="row" style="margin-top: 0.75vmin;">
        <label for="field">{{title}}</label>
        <InputText id="fieldText" v-model="values" :placeholder="value" />
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
        width: 100%;
        height: max-content;
        padding-left: 4vmin;
        font-size: 1.5vmin;
        text-align: start !important;
    }

    /deep/ .p-inputtext{
        height: 2.5vmin;
        width: 75%;
    }

    label{
        height: 2vmin;
        margin-bottom: 1%;
        width: 25%;
    }
    ::placeholder {
        color: black;
    }
  </style>