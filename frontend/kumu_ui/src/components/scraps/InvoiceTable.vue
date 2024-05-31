<template>
    <Field-set v-if="filtered == false" :style="cssProps">
        <template #legend>
            <div class="flex pl-2"
                style="background-color: rgb(124, 123, 123); border-top-left-radius: 1vmin; border-top-right-radius: 1vmin; height: 6vmin; padding-top: 1.2vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin; font-weight: bold">
                {{this.title}}</div>
        </template>
        <p class="m-0" style="color: black;">
            <DataTable v-model:selection="selectedProduct" :value="invoiceArr" selectionMode="single" dataKey="Rechnungsnummer" :metaKeySelection="false"
        @rowSelect="onRowSelect" @rowUnselect="onRowUnselect" stripedRows paginator removableSort  :rows="15" style="color: black; border-radius: 0%;"
                paginatorTemplate="FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink RowsPerPageDropdown"
                currentPageReportTemplate="{first} bis {last} von {totalRecords}">
                <ColumnColumn field="Rechnungsnummer" header="Rechnungsnummer" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Partner" header="Partner" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Datum" header="Datum" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Summe" header="Summe" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
            </DataTable>
        </p>
    </Field-set>
    <p v-else :style="cssProps">
            <DataTable v-model:selection="selectedProduct" v-model:filters="filters" :value="invoiceArr" selectionMode="single" :metaKeySelection="false"
        @rowSelect="onRowSelect" @rowUnselect="onRowUnselect" removableSort  dataKey="Rechnungsnummer" filterDisplay="row" stripedRows scrollable scrollHeight="73vmin" style="color: black; border-radius: 0%;">
                <ColumnColumn field="Rechnungsnummer" sortable filterField="Rechnungsnummer" header="Rechnungsnummer"
                    style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Rechnungsnummer }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()"
                            class="p-column-filter" placeholder="Rechnungsnummer" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Partner" sortable filterField="Partner" header="Partner"
                    style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Partner }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()"
                            class="p-column-filter" placeholder="Partner" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Datum" sortable filterField="Datum" header="Datum"
                    style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Datum }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()"
                            class="p-column-filter" placeholder="Datum" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Summe" sortable filterField="Summe" header="Summe"
                    style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Summe }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()"
                            class="p-column-filter" placeholder="Summe" />
                    </template>
                </ColumnColumn>
            </DataTable>
        </p>
</template>

<script setup>
    import { ref } from 'vue';
    import { toRef } from "vue";
</script>
  
  <script>
  import { FilterMatchMode } from 'primevue/api';

  export default {
    props:{    
        invoiceArr: {
            type: Array,
            required: true
        },
        style: {
            type: String,
        },
        filtered:{
            type: Boolean,
        },
        title:{
            type: String,
        }
    },


    created(){
        console.log(this.style);
        
        const invoiceArray = toRef(this.invoiceArr, 'invoiceArr');
        console.log(invoiceArray);
        console.log(this.invoiceArr);
    },

    data(){
        return{
            filters: {
                Rechnungsnummer: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Partner: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Datum: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Summe: { value: null, matchMode: FilterMatchMode.STARTS_WITH }
            }, 
        }       
    },

    computed: {
        cssProps() {
            if(this.style == "left" && this.filtered == false){
                return {
                    'width': '47%',
                    'float': this.style,
                    'margin-left' : '2%',
                    'margin-top': '1%',
                    'background-color': 'whitesmoke',
                    'height': '85%'
                }
            }else if (this.style == "right" && this.filtered == false){
                return {
                    'width': '47%',
                    'float':this.style,
                    'margin-right' : '2%',
                    'margin-top': '1%',
                    'background-color': 'whitesmoke',
                    'height': '85%'
                }
            }else{
                return {
                    'width': '47%',
                    'float': this.style,
                    'margin-left' : '2%',
                    'margin-top': '1%',
                    'background-color': 'whitesmoke',
                    'height': '77.5%'
                }
            }
        }
    }
  }
    const selectedProduct = ref();
    const onRowSelect = (event) => {
        console.log("selected "+  event.data.Rechnungsnummer)
    };
    const onRowUnselect = (event) => {
        console.log("unselected "+  event.data.Rechnungsnummer)
    }
  </script>
  
  <style scoped>
  </style>