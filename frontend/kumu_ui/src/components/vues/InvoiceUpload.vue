<template>
    <div id="Main">
        <NavBar :page="2"></NavBar>
        <Card style="width: 47%; height: 5.4vmin; margin-left: 2%; float:left; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;"> 
            <div style="width: 47%; float: left; padding-top: 1vmin;">Datei zum hochladen auswählen...</div>
            <FileUpload style=" margin-right: 1vmin; float: right; font-size: 2vmin" mode="basic" name="demo[]" :withCredentials="true"
                url="http://172.17.224.127:3030/api/file" accept="application/xml,text/xml" :maxFileSize="1000000" @upload="onUpload" :auto="true"
                chooseLabel="Durchsuchen" />
        </Card>
        <Card style="width: 47%; height: 5.4vmin; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;">
            <!-- evlt durch Toolbar ersetzen -->
            <InputGroup style="border-radius: 1vmin;">
                <InputText type="text" v-model="value" placeholder="Rechnungsnummer" />
                <DropdownDropdown v-model="selectedActivity" :options="activities" optionLabel="name" placeholder="Optionen" checkmark :highlightOnSelect="false" class="w-full md:w-14rem" />
                <ButtonButton label="Submit" style="border-start-end-radius: 1vmin; border-end-end-radius: 1vmin;"></ButtonButton>
            </InputGroup>
        </Card>
        <InvoiceTable  :invoiceArray= this.products :style="left" :filtered="filtered"></InvoiceTable>
        <Card style="width: 47%; height: 77.5%; margin-right: 2%; float:right; margin-top: 1%; background-color: whitesmoke; border-radius: 1vmin;padding: 1vmin;">
            <InputSwitch v-model="checked" style="float: left;"/>
            <div style="float: left; padding-left: 1vmin;">XML anzeigen</div>
            <ScrollPanel style="width: 100%; height: 95%; padding: 1vmin;">
            </ScrollPanel>
        </Card>
    </div>
</template>

<script setup>
import { ref } from 'vue';

const checked = ref(false);
const selectedActivity = ref();
const activities = ref([
                { name: 'Visualisierung' }, { name: 'Download' },
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
</style>