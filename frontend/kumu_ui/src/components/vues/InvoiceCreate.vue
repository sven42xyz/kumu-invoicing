<template>
    <div id="Main">
        <NavBar :page="1"></NavBar>
        <InvoiceTable :invoiceArray= this.products :style="left" :filtered="false" title="Eingegangene Rechnungen"></InvoiceTable>
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