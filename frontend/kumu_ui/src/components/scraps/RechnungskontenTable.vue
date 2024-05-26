<template>
    <DataTable :value="invoiceAdress" v-model:editingRows="editingRows" size="small" stripedRows scrollable  scrollHeight="110px" style="color: black; border-radius: 0%;" editMode="row" dataKey="name" @row-edit-save="onRowEditSave">
        <ColumnColumn field="name" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 125%; margin-left: 0vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="iban" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 120%; margin-left: 4vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="bank" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 80%; margin-left: 8vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn :rowEditor="true" style="width: 10%; min-width: 8rem" bodyStyle="text-align:center"></ColumnColumn>
    </DataTable>
</template>

<script setup>
    import { ref } from 'vue';

    const editingRows = ref([]);
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
        
    }
  }
    const invoiceAdress = ref([{
                name: 'Kumu AG', iban: 'DE12345678935489', bank:'Sparkasse'
            }]);
    const onRowEditSave = (event) => {
        let { newData, index } = event;

        invoiceAdress.value[index] = newData;
    };
</script>
  
  <style scoped>
  
  </style>