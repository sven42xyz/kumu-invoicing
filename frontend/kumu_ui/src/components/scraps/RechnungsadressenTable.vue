<template>
    <DataTable :value="invoiceAdress" v-model:editingRows="editingRows" size="small" stripedRows scrollable  scrollHeight="135px" style="color: black; border-radius: 0%;" editMode="row" dataKey="name" @row-edit-save="onRowEditSave" :pt="{
        column: {
            bodycell: ({ state }) => ({
                style:  state['d_editing']&&'padding-top: 0.6rem; padding-bottom: 0.6rem; font-size: small; line-height: normal'
            })
        }
    }">
        <ColumnColumn field="name" style="color: black; font-size: 1.75vmin; ">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 200%" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="ort" style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 100%; margin-left: 8.75vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="strasse" style="color: black; font-size: 1.75vmin; ">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 125%; margin-left: 8vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="plz"  style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 50%; margin-left: 11.5vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn field="land"  style="color: black; font-size: 1.75vmin;">
            <template #editor="{ data, field }">
                <InputText v-model="data[field]" style="width: 60%;  margin-left: 1vmin;" />
            </template>
        </ColumnColumn>
        <ColumnColumn :rowEditor="true" style="width: 10%; min-width: 4rem" bodyStyle="text-align:center"></ColumnColumn>
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