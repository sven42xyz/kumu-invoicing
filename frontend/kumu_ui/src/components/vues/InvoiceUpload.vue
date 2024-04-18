<template>
    <div id="Main">
        <TabMenu :model="items">
            <template #item="{ item, props }">
                <div v-if="item.type" v-ripple v-bind="props.action" @click="navigate(item.url)"
                    style="border-color: transparent;">
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.2vmin;" />
                    <span v-bind="props.label" style="font-size: 2.2vmin;">{{ item.label }}</span>
                </div>
                <div v-else v-ripple v-bind="props.action" @click="navigate(item.url)"
                    style="position: absolute; right: 0; border-color: transparent;">
                    <span v-bind="props.label" style="margin-right: 1vmin; font-size: 2vmin;">{{ item.label }}</span>
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.2vmin;" />
                </div>
            </template>
        </TabMenu>
        <Card style="width: 47%; height: 5.5vmin; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;"> 
            <div style="width: 47%; float: left; padding-top: 1vmin;">Datei zum hochladen auswählen...</div>
            <FileUpload style=" margin-right: 1vmin; float: right; font-size: 2vmin" mode="basic" name="demo[]" url="/api/upload" accept="xml/*" :maxFileSize="1000000"  @upload="onUpload" :auto="true" chooseLabel="Durchsuchen" />
        </Card>
        <Card style="width: 47%; height: 5.5vmin; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            hihi
        </Card>
        <p style="width: 47%; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            <DataTable v-model:filters="filters" :value="products" dataKey="id" filterDisplay="row" stripedRows scrollable scrollHeight="73vmin" style="color: black; border-radius: 0%; height: 80%">
                <ColumnColumn field="Rechnungsnummer" filterField="Rechnungsnummer" header="Rechnungsnummer" style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Rechnungsnummer }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Rechnungsnummer" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Partner" filterField="Partner" header="Partner" style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Partner }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Partner" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Datum" filterField="Datum" header="Datum" style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Datum }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Datum" />
                    </template>
                </ColumnColumn>
                <ColumnColumn field="Summe" filterField="Summe" header="Summe" style="color: black; font-size: 1.75vmin; padding: 1vmin">
                    <template #body="{ data }">{{ data.Summe }}</template>
                    <template #filter="{ filterModel, filterCallback }">
                        <InputText v-model="filterModel.value" type="text" @input="filterCallback()" class="p-column-filter" placeholder="Summe" />
                    </template>
                </ColumnColumn>
            </DataTable>
        </p>
        <Card style="width: 47%; height: 77.5%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            hi
        </Card>
    </div>
</template>
  
  <script> 
  import { FilterMatchMode } from 'primevue/api';
  export default {
    name: 'DashBoard',
  
    created(){
    },
    
  
    data() {
      return {
        filters: {
                Rechnungsnummer: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Partner: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Datum: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
                Summe: { value: null, matchMode: FilterMatchMode.STARTS_WITH }
            },
            items: [
                {type: 'label', label: 'Übersicht', url: '/dashboard', icon:'pi pi-home'},
                {type: 'label', label: 'Rechnungseingabe', url: '/outbound', icon:''},
                {type: 'label', label: 'Rechnungseingang', url: '/inbound', icon:''},
                {type: 'label', label: 'Stammdatenverwaltung', url: '/contacts', icon:''},
                {type: '', label: 'Abmelden ', url: '/', icon:'pi pi-sign-out'},
            ],
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
            },
            {Rechnungsnummer: '462354645',
             Partner: 'Lina',
             Datum: '23.02.2024',
             Summe: '300€'  
            },
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