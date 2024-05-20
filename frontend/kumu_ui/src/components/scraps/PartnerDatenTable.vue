<template>
    <DataTable v-model:selection="selectedProduct" :value="invoiceAdress" selectionMode="single" dataKey="name"  :metaKeySelection="false"
        @rowSelect="onRowSelect" @rowUnselect="onRowUnselect"  size="small" stripedRows scrollable scrollHeight="135px" style="color: black; border-radius: 0%;">
        <ColumnColumn field="name" style="color: black; font-size: 1.75vmin;"></ColumnColumn>
        <ColumnColumn field="ort" style="color: black; font-size: 1.75vmin;"></ColumnColumn>
        <ColumnColumn field="strasse" style="color: black; font-size: 1.75vmin;"></ColumnColumn>
        <ColumnColumn field="plz"  style="color: black; font-size: 1.75vmin;"></ColumnColumn>
        <ColumnColumn field="land"  style="color: black; font-size: 1.75vmin;"></ColumnColumn>
    </DataTable>
</template>

<script setup>
    import { ref } from 'vue';
</script>

  <script>

  export default {
    props:{    
    },

    created(){
    },

    data(){
        return{
        }       
    },

    computed: {
        
    },
    methods:{
        popUp(data){
            console.log(data);
        },
    }
  }
    const invoiceAdress = ref([{
                name: 'Kumu AG Erftstadt', ort: 'Erftstadt', strasse:'abs Straße', plz: '66666', land: 'DE'
            },{
                name: 'Kumu AG Bonn', ort: 'Bonn', strasse:'abs Straße', plz: '66666', land: 'DE'
            },{
                name: 'Kumu AG Köln', ort: 'Köln', strasse:'abs Straße', plz: '66666', land: 'DE'
            },]);
    const selectedProduct = ref();
    const onRowSelect = (event) => {
        console.log("selected "+  event.data.name)
    };
    const onRowUnselect = (event) => {
        console.log("unselected "+  event.data.name)
    }
</script>
  
  <style scoped>
  
  </style>