<template>
    <div id="Main">
        <NavBar :page="3"></NavBar>
        <Card style="width: 47%; height: 35%; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton style="background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Eigene Stammdaten</ButtonButton>
              </template>

              <template #center></template>

              <template #end> 
                <ButtonButton icon="pi pi-pencil" style="background: #4e4e51;margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
              </template>
          </ToolBar>
          <DataField :disabled="true" title = "Leitweg-Id" :value ="this.user.leitwegid"></DataField>

        </Card>
        <Card style="width: 47%; height: 85%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton style="background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Partnerdaten</ButtonButton>
              </template>

              <template #center></template>

              <template #end> 
                <ButtonButton icon="pi pi-pencil" style="background: #4e4e51; margin-right: 1vmin;border: transparent; padding: none; height: 4vmin; border-radius: 1vmin;width: 4vmin; color: white; "/>
                <ButtonButton icon="pi pi-plus" style="background: #4e4e51; margin-right: 1vmin;border: transparent; padding: none; height: 4vmin; border-radius: 1vmin;width: 4vmin; color: white; "/>
              </template>
          </ToolBar>
          <PartnerDatenTable></PartnerDatenTable>
        </Card>
        <Card style="width: 47%; height: 25%; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton style="background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Rechnungsadressen</ButtonButton>
              </template>

              <template #center></template>

              <template #end> 
                <ButtonButton icon="pi pi-plus" style="background: #4e4e51; margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
              </template>
          </ToolBar>
          <RechnungsadressenTable></RechnungsadressenTable>
        </Card>
        <Card style="width: 47%; height: 21%; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton style="background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Rechnungskonten</ButtonButton>
              </template>

              <template #center></template>

              <template #end> 
                <ButtonButton icon="pi pi-plus" style="background: #4e4e51; margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
              </template>
          </ToolBar>
          <RechnungskontenTable></RechnungskontenTable>
        </Card>
    </div>
</template>

<script setup>
import NavBar from '../scraps/NavBar.vue';
import DataField from '../scraps/DataField.vue'
import RechnungsadressenTable from '../scraps/RechnungsadressenTable'
import RechnungskontenTable from '../scraps/RechnungskontenTable'
import PartnerDatenTable from '../scraps/PartnerDatenTable'
</script>

<script> 

  export default {
    name: 'DashBoard',
  
    created(){
        console.log(this.products);
    },
  
    data() {
      return {
        left: "left",
        right: "right",
        filtered: false,
        user: 
            {leitwegid: '1111:012345678',
             Partner: 'Hans',
             Datum: '23.01.2024',
             Summe: '1300€'  
            }
        ,
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