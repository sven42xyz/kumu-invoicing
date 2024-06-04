<template>
    <div class="row" v-if="disabled==true" style="margin-top: 1vmin;" >
        <label for="field">{{title}}</label>
    </div>
    <div v-else class="row" style="margin-top: 1vmin;">
        <label for="field">{{title}}</label>
    </div>
    <div class="row" v-if="disabled==true" >
        <InputText id="fieldText" disabled v-model="values" :placeholder="value"></InputText>
    </div>
    <div v-else class="row">
        <InputText id="fieldText" v-model="values" :placeholder="value"></InputText>
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
        font-size: 1.75vmin;
        text-align: start !important;
    }

    /deep/ .p-inputtext{
        height: 3vmin;
    }

    label{
        height: 2.5vmin;
        margin-bottom: 1%;
    }

    ::placeholder {
        color: black;
    }
  </style>