<template>
    <div id = "Main">
        <TabMenu :model="items">
            <template #item="{ item, props }">
                <div v-if="item.type"  v-ripple v-bind="props.action" @click="navigate(item.url)" style="border-color: transparent;">
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.25vmin;" />
                    <span v-bind="props.label" style="font-size: 2.25vmin;">{{ item.label }}</span>
                </div>
                <div v-else  v-ripple v-bind="props.action" @click="navigate(item.url)" style="position: absolute; right: 0; border-color: transparent;">
                    <span v-bind="props.label" style="margin-right: 1vmin; font-size: 2vmin;">{{ item.label }}</span>
                    <span v-bind="props.icon" style="color: 'var(--primary-color)'; font-size: 2.25vmin;" />
                </div>
            </template>
        </TabMenu>
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
  </style>