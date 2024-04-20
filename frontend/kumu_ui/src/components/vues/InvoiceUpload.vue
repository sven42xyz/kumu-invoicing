<template>
    <div id="Main">
        <TabMenu :model="items" :active-index="2">
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
        <Card style="width: 47%; height: 5.4vmin; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;"> 
            <div style="width: 47%; float: left; padding-top: 1vmin;">Datei zum hochladen auswählen...</div>
            <FileUpload style=" margin-right: 1vmin; float: right; font-size: 2vmin" mode="basic" name="demo[]" url="/api/upload" accept="xml/*" :maxFileSize="1000000"  @upload="onUpload" :auto="true" chooseLabel="Durchsuchen" />
        </Card>
        <Card style="width: 47%; height: 5.4vmin; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            <InputGroup style="border-radius: 1vmin;">
                <InputText type="text" v-model="value" placeholder="Rechnungsnummer" />
                <DropdownDropdown v-model="selectedActivity" :options="activities" optionLabel="name" placeholder="Optionen" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />
                <ButtonButton label="Submit" style="border-start-end-radius: 1vmin; border-end-end-radius: 1vmin;"></ButtonButton>
            </InputGroup>
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
        <Card style="width: 47%; height: 77.5%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;padding: 1vmin;">
            <InputSwitch v-model="XML_PDF" style="float: left;"/>
            <div style="float: left; padding-left: 1vmin;">XML anzeigen</div>
            <ScrollPanel style="width: 100%; height: 95%; padding: 1vmin;">
                <p>
                    Hier könnte Ihre Pdf stehen ...
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Consectetur, adipisci velit, sed quia non numquam eius modi.

At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus.

Quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Consectetur, adipisci velit, sed quia non numquam eius modi.

At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus.

Quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat
                </p>
            </ScrollPanel>
        </Card>
    </div>
</template>
  
  <script> 
  import { FilterMatchMode } from 'primevue/api';
  export default {
    name: 'DashBoard',
    activeIndex : 2,
  
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
                {type: 'label', label: 'Übersicht', url: '/dashboard', icon:'pi pi-home', active: false},
                {type: 'label', label: 'Rechnungseingabe', url: '/outbound', icon:'', active: false},
                {type: 'label', label: 'Rechnungseingang', url: '/inbound', icon:'', active: true},
                {type: 'label', label: 'Stammdatenverwaltung', url: '/contacts', icon:'', active: false},
                {type: '', label: 'Abmelden ', url: '/', icon:'pi pi-sign-out', active: false},
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
        activities:[
            {name: 'Visualisierung'},{name: 'Download'},
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