<template>
    <div id="Main">
        <NavBar :page="3"></NavBar>
        <Card style="width: 47%; height: 35%; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton disabled style="opacity: 1;background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Eigene Stammdaten</ButtonButton>
              </template>

              <template #center></template>

              <template #end> 
                <ButtonButton v-if="disabled==true" icon="pi pi-pencil" @click="enable()" style="background: #4e4e51;margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
                <div class="enabled row" v-else  style="align-content: end;">
                  <ButtonButton icon="pi pi-times" @click="disable()" style="background: #4e4e51;margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
                  <ButtonButton icon="pi pi-check" @click="save()" style="background: #4e4e51;margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
                </div>
                </template>
          </ToolBar>
          <div class="row" style="height: 45%;">
            <div class="column" style="height: 100%;width: 49.5%">
              <DataField :disabled=disabled title = "Leitweg-Id" ref="leitwegid" :value ="this.user.leitwegid"></DataField>
              <DataField :disabled=disabled title = "Name des Unternehmens" ref="name" :value ="this.user.name"></DataField>
            </div>
            <div class="column" style="height: 100%;width: 49.5%">
              <DataField :disabled=disabled title = "Umsatzsteuernummer" ref="vatid" :value ="this.user.vatid"></DataField>
              <DataField :disabled=disabled title = "Registrierter Firmenname" ref="reName" :value ="this.user.regName"></DataField>
            </div>
          </div>
          <hr style="width: 100%; margin: 0"/>
          <div class="row" style="height: 35%; margin-top: 0.75%;">
            <label style="text-align: left; font-weight: 600; padding-left: 5vmin;">Primärer Kontakt</label>
            <div class="column" style="height: 100%;width: 49.5%">
              <DataFieldSmall :disabled=disabled title = "Name" ref="pcname" :value ="this.user.pcname"></DataFieldSmall>
              <DataFieldSmall :disabled=disabled title = "E-Mail" ref="pcemail" :value ="this.user.pcemail"></DataFieldSmall>
            </div>
            <div class="column" style="height: 100%;width: 49.5%">
              <DataFieldSmall :disabled=disabled title = "Telefon" ref="pctel" :value ="this.user.pctel"></DataFieldSmall>
              <DataFieldSmall :disabled=disabled title = "Anrede" ref="anrede" :value ="this.user.anrede"></DataFieldSmall>
            </div>
          </div>

        </Card>
        <Card style="width: 47%; height: 85%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
          <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0vmin; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
              <template #start>
                  <ButtonButton disabled style="opacity: 1;background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Partnerdaten</ButtonButton>
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
                  <ButtonButton disabled style="opacity: 1;background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Rechnungsadressen</ButtonButton>
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
                  <ButtonButton disabled style="opacity: 1; background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Rechnungskonten</ButtonButton>
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
import DataField from '../scraps/DataField.vue';
import DataFieldSmall from '../scraps/DataFieldSmall.vue';
import RechnungsadressenTable from '../scraps/RechnungsadressenTable';
import RechnungskontenTable from '../scraps/RechnungskontenTable';
import PartnerDatenTable from '../scraps/PartnerDatenTable';
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
        disabled:true ,
        filtered: false,
        user: 
            {leitwegid: '1111:012345678',
             vatid: 'DE12345678',
             name: 'Kumu AG',
             regName: 'Kumu AG',
             pcname: 'Marie',
             pctel: '02318534854',
             pcemail: 'mc@kumu.de',
             anrede: 'Frau'
            }
        ,
      };
    },

    methods: {
        navigate(i){
            this.$router.push(i);    
        },

        enable(){
          this.disabled = false;
        },

        disable(){
          this.disabled = true;
        },

        save(){
          console.log(this.$refs);
          if(this.$refs['leitwegid'] !=null){
            this.user.leitwegid = this.$refs['leitwegid'].value;
          }
          if(this.$refs['vatid'] !=null){
            this.user.vatid = this.$refs['vatid'].value;
          }
          if(this.$refs['name'] !=null){
            this.user.name = this.$refs['name'].value;
          }
          if(this.$refs['regName'] !=null){
            this.user.regName = this.$refs['regName'].value;
          }
          if(this.$refs['pcname'] !=null){
            this.user.pcname = this.$refs['pcname'].value;
          }
          if(this.$refs['pctel'] !=null){
            this.user.pctel = this.$refs['pctel'].value;
          }
          if(this.$refs['pcemail'] !=null){
            this.user.pcemail = this.$refs['pcemail'].value;
          }
          if(this.$refs['anrede'] !=null){
            this.user.anrede = this.$refs['anrede'].value;
          }

          this.disabled = true;
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

    /deep/ .p-toolbar-group-end{
        margin-right: 1vmin;
        align-content: end !important;
    }

    /deep/ .p-inputtext:disabled{
      background-color: lightgray;
      color: black!important;
      font-display: unset;
    }
    .enabled.row{
        margin-right: 0.25vmin;
    }
  </style>