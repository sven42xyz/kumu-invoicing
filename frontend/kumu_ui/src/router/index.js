import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../components/vues/LandingPage.vue'
import DashBoard from '../components/vues/DashBoard.vue'
import InvoiceUpload from '../components/vues/InvoiceUpload.vue'
import InvoiceCreate from '../components/vues/InvoiceCreate.vue'
import CustomerData from '../components/vues/CustomerData.vue'

const routes = [
    { path: '/', component: LandingPage },
    { path: '/dashboard', name: 'dashboard', component: DashBoard },
    { path: '/inbound', name: 'inbound', component: InvoiceUpload },
    { path: '/outbound', name: 'outbound', component: InvoiceCreate },
    { path: '/contacts', name: 'contacts', component: CustomerData },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;