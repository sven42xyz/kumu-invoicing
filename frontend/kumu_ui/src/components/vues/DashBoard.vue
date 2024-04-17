<template>
    <div id="Main">
        <TabMenu :model="items">
            <template #item="{ item, props }">
                <div v-if="item.type" v-ripple v-bind="props.action" @click="navigate(item.url)"
                    style="border-color: transparent;">
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.25vmin;" />
                    <span v-bind="props.label" style="font-size: 2.25vmin;">{{ item.label }}</span>
                </div>
                <div v-else v-ripple v-bind="props.action" @click="navigate(item.url)"
                    style="position: absolute; right: 0; border-color: transparent;">
                    <span v-bind="props.label" style="margin-right: 1vmin; font-size: 2vmin;">{{ item.label }}</span>
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.25vmin;" />
                </div>
            </template>
        </TabMenu>
        <Field-set style="width: 45%; margin-left: 2.5%; margin-top: 1%; background-color: whitesmoke;">
            <template #legend>
                <div class="flex pl-2"
                    style="background-color: rgb(124, 123, 123); border-top-left-radius: 1vmin; border-top-right-radius: 1vmin; height: 6vmin; padding: 1vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
                    Eingehend</div>
            </template>
            <p class="m-0" style="color: black;">
                <DataTable :value="products" paginator :rows="5" :rowsPerPageOptions="[5, 10, 20]"
                    style="color: black; font-size: 2vmin; border-radius: 0%;"
                    paginatorTemplate="FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink RowsPerPageDropdown"
                    currentPageReportTemplate="{first} bis {last} von {totalRecords}">
                    <ColumnColumn field="Rechnungsnummer" header="Rechnungsnummer"
                        style="background-color: whitesmoke;color: black; font-size: 2vmin;"></ColumnColumn>
                </DataTable>
            </p>
        </Field-set>
    </div>
</template>
  
  <script> 
  export default {
    name: 'DashBoard',
  
    created(){
    },
  
    data() {
      return {
        items: [
                {type: 'label', label: 'Übersicht', url: '/dashboard', icon:'pi pi-home'},
                {type: 'label', label: 'Rechnungseingabe', url: '/dashboard', icon:''},
                {type: 'label', label: 'Rechnungseingabe', url: '/dashboard', icon:''},
                {type: 'label', label: 'Stammdatenverwaltung', url: '/dashboard', icon:''},
                {type: '', label: 'Abmelden ', url: '/', icon:'pi pi-sign-out'},
            ],
        products: [
            {Rechnungsnummer: 'Rechnungsnummer'}, {Rechnungsnummer: 'Rechnungsnummer'}
        ]
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