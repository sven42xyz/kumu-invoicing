<template>
    <DataTable :value="invoiceAdress" v-model:editingRows="editingRows" size="small" stripedRows scrollable  scrollHeight="135px" style="color: black; border-radius: 0%;" editMode="row" dataKey="name" @row-edit-save="onRowEditSave">
        <ColumnColumn field="name" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="ort" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="strasse" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="plz"  style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="land"  style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" />
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
                name: 'Kumu AG Erftstadt', ort: 'Erftstadt', strasse:'abs Straße', plz: '66666', land: 'DE'
            },{
                name: 'Kumu AG Erftstadt', ort: 'Erftstadt', strasse:'abs Straße', plz: '66666', land: 'DE'
            },{
                name: 'Kumu AG Erftstadt', ort: 'Erftstadt', strasse:'abs Straße', plz: '66666', land: 'DE'
            },{
                name: 'Kumu AG Erftstadt', ort: 'Erftstadt', strasse:'abs Straße', plz: '66666', land: 'DE'
            },]);
    const onRowEditSave = (event) => {
        let { newData, index } = event;

        invoiceAdress.value[index] = newData;
    };
</script>
  
  <style scoped>
  
  </style>