<template>
    <TabMenu :model="items" :active-index="page">
        <template #item="{ item, props }">
            <div v-if="item.type" :target="item.target" v-bind="props.action" @click="navigate(item.url)"
                style="border-color: transparent;">
                <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2vmin;" />
                <span v-bind="props.label" style="font-size: 2vmin;">{{ item.label }}</span>
            </div>
            <div v-else :target="item.target" v-bind="props.action" @click="navigate(item.url)"
                style="position: absolute; right: 0; border-color: transparent;">
                <span v-bind="props.label" style="margin-right: 1vmin; font-size: 2vmin;">{{ item.label }}</span>
                <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2vmin;" />
            </div>
        </template>
    </TabMenu>
</template>

<script setup>
</script>

<script> 

  export default {
    props:{    
        page:{
            type: String,
        }
    },
  
    data() {
      return {
        items: [
                {type: 'label', label: 'Übersicht', url: '/dashboard', icon:'pi pi-home'},
                {type: 'label', label: 'Rechnungseingabe', url: '/outbound', icon:'pi pi-upload'},
                {type: 'label', label: 'Rechnungseingang', url: '/inbound', icon:'pi pi-download'},
                {type: 'label', label: 'Stammdatenverwaltung', url: '/contacts', icon:'pi pi-address-book'},
                {type: '', label: 'Abmelden ', url: '/', icon:'pi pi-sign-out'},
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