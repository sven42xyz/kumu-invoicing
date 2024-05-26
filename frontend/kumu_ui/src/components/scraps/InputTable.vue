<template>
    <Field-set :style="cssProps">
        <template #legend>
            <div class="flex pl-2"
                style="background-color: rgb(124, 123, 123);  border-top-left-radius: 1vmin; border-top-right-radius: 1vmin; height: 6vmin; padding: 1vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
                {{this.title}}</div>
        </template>
        <p class="m-0" style="color: black;">
            <DataTable :value="invoiceArray" stripedRows paginator removableSort  :rows="15" style="color: black; border-radius: 0%;"
                paginatorTemplate="FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink RowsPerPageDropdown"
                currentPageReportTemplate="{first} bis {last} von {totalRecords}">
                <ColumnColumn field="Rechnungsnummer" header="Rechnungsnummer" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Partner" header="Partner" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Datum" header="Datum" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
                <ColumnColumn field="Summe" header="Summe" sortable style="color: black; font-size: 1.75vmin; padding: 1vmin"></ColumnColumn>
            </DataTable>
        </p>
    </Field-set>
</template>
  
  <script>
  import { FilterMatchMode } from 'primevue/api';

  export default {
    props:{    
        invoiceArray: {
            type: Array,
            required: true
        },
        style: {
            type: String,
        },
        title:{
            type: String,
        }
    },

    created(){

        console.log(this.style);

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
        console.log("selected "+  event.data.name)
    };
    const onRowUnselect = (event) => {
        console.log("unselected "+  event.data.name)
    }
  </script>
  
  <style scoped>
  </style>