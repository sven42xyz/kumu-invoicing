<template>
    <div id="Main">
        <NavBar :page="3"></NavBar>
    </div>
</template>

<script setup>
import NavBar from '../scraps/NavBar.vue';
</script>

<script> 

  export default {
    name: 'DashBoard',
  
    created(){
        console.log(this.products);
    },
  
    data() {
      return {
        left: "left",
        right: "right",
        filtered: false,
        products: [
            {Rechnungsnummer: '012345678',
             Partner: 'Hans',
             Datum: '23.01.2024',
             Summe: '1300€'  
            }, 
            {Rechnungsnummer: '547654576',
             Partner: 'Anna',
             Datum: '23.02.2024',
             Summe: '222€'  
            }
        ],
      };
    },

    methods: {
        navigate(i){
            this.$router.push(i);    
        }
  
    }
  }
  </script>
  
  <style scoped>
   #Main{
    width: 100%;
    height: 100%;
   }
 
   .p-paginator {
        background: whitesmoke;
        color: grey;
    }
  </style>