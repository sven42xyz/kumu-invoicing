<template>
    <div id="Main">
        <NavBar :page="1"></NavBar>
        <div v-if="create==false" style="height: 100%;">
            <InvoiceTable :invoiceArray= this.products :style="left" :filtered="false" title="Erstellte Rechnungen"></InvoiceTable>
            <ToolBar style="justify-content: space-evenly; gap: 0; width: 47%; height: 5.4vmin; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin; padding: 0;">
                <template #start>
                    <InputText style="height: 5.3vmin; width: 100%; border-start-end-radius: 0; border-end-end-radius: 0;" type="text" v-model="value" placeholder="Rechnungsnummer" />
                </template>

                <template #center>
                    <DropdownDropdown style="height: 5.3vmin; width: 100%; border-radius: 0;" v-model="selectedActivity" :options="activities" optionLabel="name" placeholder="Optionen" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />
                </template>

                <template #end> <ButtonButton label="Submit" style="width: 100%; height: 5.3vmin; border-color: rgb(189, 189, 189); border-start-end-radius: 1vmin; border-end-end-radius: 1vmin;"></ButtonButton></template>
            </ToolBar>
            <Card style="width: 47%; height: 79.75%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;padding: 1vmin;">
                <InputSwitch v-model="checked" style="float: left;"/>
                <ButtonButton icon="pi pi-plus" style="float:right;background: #4e4e51; margin-right: 1vmin;border: transparent; padding: none; height: 4vmin; border-radius: 1vmin;width: 4vmin; color: white; "/>
                <div style="float: left; padding-left: 1vmin;">XML anzeigen</div>
                <ScrollPanel style="width: 100%; height: 95%; padding: 1vmin;">
                    Insert PDF here
                </ScrollPanel>
            </Card>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const checked = ref(false);
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
</style>