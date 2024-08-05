<template>
    <div id="Main">
        <NavBar :page="1"></NavBar>
        <ToolBar v-if="create == false" style="justify-content: space-evenly; gap: 0; float:left; width: 47%; height: 5.4vmin; margin-left: 2%; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin; padding: 0;">
            <template #start>
                <InputText style="height: 5.3vmin; width: 100%; border-start-end-radius: 0; border-end-end-radius: 0;" type="text" v-model="value" placeholder="Rechnungsnummer" />
            </template>

            <template #center>
                <DropdownDropdown style="height: 5.3vmin; width: 100%; border-radius: 0;" v-model="selectedActivity" :options="activities" optionLabel="name" placeholder="Optionen" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />
            </template>

            <template #end> <ButtonButton label="Submit" style="width: 100%; height: 5.3vmin; border-color: rgb(189, 189, 189); border-start-end-radius: 1vmin; border-end-end-radius: 1vmin;"></ButtonButton></template>
        </ToolBar>  
        <Card :style="cssProps()" style="width: 47%; height: 85%; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            <ToolBar style="background-color: rgb(124, 123, 123); border-color: rgb(124, 123, 123); border-bottom-left-radius: 0; border-bottom-right-radius: 0; border-top-left-radius: 0.9vmin; border-top-right-radius: 0.9vmin; height: 6vmin; padding: 0; text-align: left; text-indent: 2.5%; font-size: 2.5vmin;">
                <template #start>
                    <ButtonButton disabled style="opacity: 1;background: transparent; border: transparent; padding: none;border-radius: 1vmin; color: white; font-size: 2.5vmin; font-weight: bold;">Neue Rechnung erstellen</ButtonButton>
                </template>
                <template #end> 
                    <ButtonButton icon="pi pi-trash" @click="disable()" style="background: #4e4e51; margin-right: 1vmin; border: transparent; padding: none; height: 4vmin; width: 4vmin; border-radius: 1vmin; color: white; "/>
                </template>
            </ToolBar>
            <StepPer linear>
                <StepperPanel>
                    <template #content="{ nextCallback }">
                        <div class="flex flex-column h-12rem">
                            <div class="border-2 border-dashed surface-border border-round surface-ground flex-auto flex justify-content-center align-items-center font-medium">Content I</div>
                        </div>
                        <div class="flex pt-4 justify-content-end">
                            <ButtonButton label="Next"  class="panelbutton" icon="pi pi-arrow-right" iconPos="right" @click="nextCallback" @click.stop="createWindow()"/>
                        </div>
                    </template>
                </StepperPanel>
                <StepperPanel>
                    <template #content="{ prevCallback, nextCallback }">
                        <div class="flex flex-column h-12rem">
                            <div class="border-2 border-dashed surface-border border-round surface-ground flex-auto flex justify-content-center align-items-center font-medium">Content II</div>
                        </div>
                        <div class="flex pt-4 justify-content-between">
                            <ButtonButton label="Back" class="panelbutton" severity="secondary" icon="pi pi-arrow-left" @click="prevCallback" @click.stop="createWindow()" />
                            <ButtonButton label="Next"  class="panelbutton" icon="pi pi-arrow-right" iconPos="right" @click="nextCallback" />
                        </div>
                    </template>
                </StepperPanel>
                <StepperPanel>
                    <template #content="{ prevCallback }">
                        <div class="flex flex-column h-12rem">
                            <div class="border-2 border-dashed surface-border border-round surface-ground flex-auto flex justify-content-center align-items-center font-medium">Content III</div>
                        </div>
                        <div class="flex pt-4 justify-content-start">
                            <ButtonButton label="Back"  class="panelbutton" severity="secondary" icon="pi pi-arrow-left" @click="prevCallback" />
                        </div>
                    </template>
                </StepperPanel>
            </StepPer>           
        </Card>
        <InvoiceTable  v-if="create == false" :invoiceArr= this.products :style="left" :filtered="filtered"></InvoiceTable>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const selectedActivity = ref();
const activities = ref([
                { name: 'Visualisierung' }, { name: 'Download' },{ name: 'Erstellung' }
            ])
</script>

<script>
import InvoiceTable from '../scraps/InvoiceTable.vue';
import NavBar from '../scraps/NavBar.vue';


export default {
    name: 'DashBoard',
    activeIndex: 2,
    created() {
    },
    data() {
        return {
            left: "left",
            right: "right",
            filtered: true,
            create: false,
            products: [
                { Rechnungsnummer: '012345678',
                    Partner: 'Hans',
                    Datum: '23.01.2024',
                    Summe: '1300€'
                },
                { Rechnungsnummer: '547654576',
                    Partner: 'Anna',
                    Datum: '23.02.2024',
                    Summe: '222€'
                },
                { Rechnungsnummer: '462354645',
                    Partner: 'Lina',
                    Datum: '23.02.2024',
                    Summe: '300€'
                },
            ]
        };
    },
    methods: {
        navigate(i) {
            this.$router.push(i);
        },
        onUpload() {
            console.log("hello");
        },
        createWindow() {
            if(this.create == false){
                this.create = true;
            }else{
                this.create = false;
            }
            
        },
        cssProps() {
            if(this.create == false){
                return {
                    'justify-content': 'space-evenly',
                    'gap': '0',
                    'width': '47%',
                    'height': '5.4vmin',
                    'margin-right': '2%',
                    'float': 'right',
                    'margin-top': '1%',
                    'background-color': 'whitesmoke',
                    'border-radius': '1vmin',
                    'padding': '0'
                }
            }else{
                return {
                    'justify-content': 'space-evenly',
                    'gap': '0',
                    'width': '47%',
                    'height': '5.4vmin',
                    'margin-left': '2%',
                    'float': 'left',
                    'margin-top': '1%',
                    'background-color': 'whitesmoke',
                    'border-radius': '1vmin',
                    'padding': '0'
                }
            }
        }
    },
    components: { InvoiceTable }
}
</script>

<style scoped>
#Main {
    width: 100%;
    height: 100%;
}

.p-paginator {
    background: whitesmoke;
    color: grey;
}


:deep() .p-toolbar-group-center{
    width: 23.33%;
}

:deep()  .p-toolbar-group-start{
    width: 63.33%;
}

:deep() .p-toolbar-group-end{
    width: 13.33%;
}

:deep() .p-stepper-panels {
    height: 70vmin;
    margin-top: 2.5vmin;
    background-color: whitesmoke;
}
:deep() .p-stepper-action {
    background-color: whitesmoke;
}
:deep() .p-stepper-number {
    background-color: whitesmoke;
    color: #020617;
}

.panelbutton{
    background-color: gainsboro;
    color: #020617;
}

:deep() card .p-toolbar-group-center{
    width: 0;
}
:deep() card .p-toolbar-group-end{
    width: unset;
}
</style>